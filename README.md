NetProbe

NetProbe is a high-performance command-line tool for verifying the liveness of IPs and domains, checking open ports, and performing large-scale network probes. It is designed to handle 80,000+ targets efficiently and is inspired by tools like httpx, with a focus on scalability and speed.

Author: Cristopher

This tool is intended for educational and legitimate cybersecurity purposes, such as penetration testing on systems you own or have explicit permission to scan. Always comply with laws and obtain permission before scanning any network.

Features





Verify if an IP or domain is live by checking connectivity on a specific port (default: 80).



Scan a single IP or domain for open ports within a range.



Check live status of domains/URLs with HTTP/HTTPS support.



Process large lists (80k+ IPs/domains) with asynchronous and multi-threaded execution.



Progress tracking with a progress bar.



Output results in JSON or plain text.



Customizable timeout, threads, and batch sizes for performance tuning.

Installation

Clone the repository and install dependencies:

git clone https://github.com/yourusername/netprobe.git
cd netprobe
pip install -r requirements.txt
pip install .

Usage

Run netprobe --help for options.

Examples





Check if an IP is live on port 80:

netprobe ip 192.168.1.1



Check IP on specific port:

netprobe ip 192.168.1.1 --port 443



Scan ports on IP:

netprobe scan-ip 192.168.1.1 --start-port 1 --end-port 1024



Check domain live:

netprobe domain example.com --port 80



Scan multiple from file (supports 80k+ targets):

netprobe batch --file targets.txt --port 80

Where targets.txt contains IPs or domains, one per line.

Options





--port PORT: Specify the port to check (default: 80).



--timeout TIMEOUT: Connection timeout in seconds (default: 2).



--threads THREADS: Number of threads for IP/port scans (default: 50).



--batch-size SIZE: Number of targets per batch (default: 1000).



--output FILE: Save results to a file.



--json: Output in JSON format.

Requirements





Python 3.8+



Dependencies: aiohttp, tqdm

See requirements.txt for details.

License

MIT License. See LICENSE file.

Disclaimer

Use responsibly. Scanning without permission may be illegal. The author is not responsible for misuse.
