import argparse
import sys
import json
import asyncio
from tqdm import tqdm
from netprobe.core import is_ip_live, scan_ports, batch_check, parse_target, is_domain_live

def print_banner():
    """Print ASCII art banner with tool name and author."""
    banner = """
    ╔════════════════════════════════════════════╗
    ║          NetProbe - Network Scanner        ║
    ║          Author: Cristopher                ║
    ║          Version: 0.2.0                    ║
    ╚════════════════════════════════════════════╝
    """
    print(banner)

def main():
    print_banner()
    
    parser = argparse.ArgumentParser(description="NetProbe: High-performance network probing tool.")
    subparsers = parser.add_subparsers(dest='command', required=True)

    # IP command
    ip_parser = subparsers.add_parser('ip', help='Check single IP')
    ip_parser.add_argument('target', help='IP address')
    ip_parser.add_argument('--port', type=int, default=80, help='Port to check')

    # Scan IP command
    scan_ip_parser = subparsers.add_parser('scan-ip', help='Scan ports on IP')
    scan_ip_parser.add_argument('target', help='IP address')
    scan_ip_parser.add_argument('--start-port', type=int, default=1, help='Start port')
    scan_ip_parser.add_argument('--end-port', type=int, default=1024, help='End port')

    # Domain command
    domain_parser = subparsers.add_parser('domain', help='Check single domain or URL')
    domain_parser.add_argument('target', help='Domain or URL')
    domain_parser.add_argument('--port', type=int, default=80, help='Port to check')

    # Batch command
    batch_parser = subparsers.add_parser('batch', help='Check multiple from file')
    batch_parser.add_argument('--file', required=True, help='File with targets (one per line)')
    batch_parser.add_argument('--port', type=int, default=80, help='Port to check')
    batch_parser.add_argument('--threads', type=int, default=50, help='Number of threads')
    batch_parser.add_argument('--batch-size', type=int, default=1000, help='Batch size for processing')

    # Common options
    for p in [ip_parser, scan_ip_parser, domain_parser, batch_parser]:
        p.add_argument('--timeout', type=int, default=2, help='Timeout in seconds')
        p.add_argument('--output', help='Output file')
        p.add_argument('--json', action='store_true', help='Output in JSON')

    args = parser.parse_args()

    loop = asyncio.get_event_loop()

    if args.command == 'ip':
        live = is_ip_live(args.target, args.port, args.timeout)
        result = {'target': args.target, 'port': args.port, 'live': live}

    elif args.command == 'scan-ip':
        print(f"Scanning {args.target} from port {args.start_port} to {args.end_port}...")
        open_ports = loop.run_until_complete(
            scan_ports(args.target, args.start_port, args.end_port, args.timeout, args.threads)
        )
        result = {'target': args.target, 'open_ports': open_ports}

    elif args.command == 'domain':
        async def check_domain():
            async with aiohttp.ClientSession() as session:
                return await is_domain_live(session, args.target, args.port, args.timeout)
        result = loop.run_until_complete(check_domain())

    elif args.command == 'batch':
        try:
            with open(args.file, 'r') as f:
                targets = [line.strip() for line in f if line.strip()]
        except FileNotFoundError:
            print(f"Error: File {args.file} not found")
            sys.exit(1)
        
        print(f"Scanning {len(targets)} targets on port {args.port}...")
        results = loop.run_until_complete(
            batch_check(targets, args.port, args.timeout, args.threads, args.batch_size)
        )
        result = {r['target']: r for r in results}

    if args.json:
        output = json.dumps(result, indent=4)
    else:
        output = str(result)

    print(output)
    if args.output:
        with open(args.output, 'w') as f:
            f.write(output)

if __name__ == "__main__":
    import aiohttp
    main()
