import socket
import threading
import queue
import json
from urllib.parse import urlparse

def is_live(target, port=80, timeout=2):
    """Check if target (IP or domain) is live on the given port."""
    try:
        ip = socket.gethostbyname(target) if not target.replace('.', '').isdigit() else target
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(timeout)
        result = sock.connect_ex((ip, port))
        sock.close()
        return result == 0
    except Exception:
        return False

def scan_ports(target, start_port=1, end_port=1024, timeout=2):
    """Scan a range of ports on the target."""
    open_ports = []
    for port in range(start_port, end_port + 1):
        if is_live(target, port, timeout):
            open_ports.append(port)
    return open_ports

def batch_check(targets, port=80, timeout=2, threads=10):
    """Check multiple targets in parallel."""
    results = {}
    q = queue.Queue()
    for target in targets:
        q.put(target)

    def worker():
        while not q.empty():
            target = q.get()
            results[target] = is_live(target, port, timeout)
            q.task_done()

    thread_list = []
    for _ in range(threads):
        t = threading.Thread(target=worker)
        t.start()
        thread_list.append(t)

    q.join()
    for t in thread_list:
        t.join()

    return results

def parse_target(target):
    """Parse URL to extract domain."""
    if target.startswith('http'):
        parsed = urlparse(target)
        return parsed.hostname
    return target
