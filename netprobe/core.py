import socket
import threading
import queue
import asyncio
import aiohttp
import json
from urllib.parse import urlparse
from tqdm import tqdm

async def is_domain_live(session, target, port=80, timeout=2):
    """Check if a domain or URL is live on the given port using HTTP/HTTPS."""
    target = parse_target(target)
    try:
        ip = socket.gethostbyname(target)
    except socket.gaierror:
        return {'target': target, 'port': port, 'live': False, 'ip': None}
    
    scheme = 'https' if port == 443 else 'http'
    url = f"{scheme}://{target}:{port}"
    
    try:
        async with session.get(url, timeout=timeout) as response:
            return {
                'target': target,
                'port': port,
                'live': response.status in range(200, 400),
                'ip': ip,
                'status': response.status
            }
    except (aiohttp.ClientError, asyncio.TimeoutError):
        # Fallback to TCP check
        return {
            'target': target,
            'port': port,
            'live': is_ip_live(ip, port, timeout),
            'ip': ip,
            'status': None
        }

def is_ip_live(ip, port=80, timeout=2):
    """Check if an IP is live on the given port."""
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(timeout)
        result = sock.connect_ex((ip, port))
        sock.close()
        return result == 0
    except socket.error:
        return False

async def scan_ports(target, start_port=1, end_port=1024, timeout=2, threads=50):
    """Scan a range of ports on the target using threading."""
    open_ports = []
    q = queue.Queue()
    for port in range(start_port, end_port + 1):
        q.put(port)

    def worker():
        while not q.empty():
            port = q.get()
            if is_ip_live(target, port, timeout):
                open_ports.append(port)
            q.task_done()

    thread_list = []
    for _ in range(min(threads, end_port - start_port + 1)):
        t = threading.Thread(target=worker)
        t.start()
        thread_list.append(t)

    q.join()
    for t in thread_list:
        t.join()

    return open_ports

async def batch_check(targets, port=80, timeout=2, threads=50, batch_size=1000):
    """Check multiple targets in parallel with batch processing."""
    results = []
    
    async with aiohttp.ClientSession() as session:
        for i in range(0, len(targets), batch_size):
            batch = targets[i:i + batch_size]
            tasks = [is_domain_live(session, target, port, timeout) for target in batch]
            batch_results = await asyncio.gather(*tasks, return_exceptions=True)
            results.extend(batch_results)
    
    return results

def parse_target(target):
    """Parse URL to extract domain."""
    if target.startswith(('http://', 'https://')):
        parsed = urlparse(target)
        return parsed.hostname
    return target
