import argparse
import sys
import json
from netprobe.core import is_live, scan_ports, batch_check, parse_target

def main():
    parser = argparse.ArgumentParser(description="NetProbe: Network probing tool.")
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
    batch_parser.add_argument('--threads', type=int, default=10, help='Number of threads')

    # Common options
    for p in [ip_parser, scan_ip_parser, domain_parser, batch_parser]:
        p.add_argument('--timeout', type=int, default=2, help='Timeout in seconds')
        p.add_argument('--output', help='Output file')
        p.add_argument('--json', action='store_true', help='Output in JSON')

    args = parser.parse_args()

    if args.command == 'ip':
        target = args.target
        live = is_live(target, args.port, args.timeout)
        result = {'target': target, 'port': args.port, 'live': live}

    elif args.command == 'scan-ip':
        target = args.target
        open_ports = scan_ports(target, args.start_port, args.end_port, args.timeout)
        result = {'target': target, 'open_ports': open_ports}

    elif args.command == 'domain':
        target = parse_target(args.target)
        live = is_live(target, args.port, args.timeout)
        result = {'target': target, 'port': args.port, 'live': live}

    elif args.command == 'batch':
        with open(args.file, 'r') as f:
            targets = [line.strip() for line in f if line.strip()]
        results = batch_check(targets, args.port, args.timeout, args.threads)
        result = results

    if args.json:
        output = json.dumps(result, indent=4)
    else:
        output = str(result)

    print(output)
    if args.output:
        with open(args.output, 'w') as f:
            f.write(output)

if __name__ == "__main__":
    main()
