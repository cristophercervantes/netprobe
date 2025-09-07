# NetProbe

NetProbe is a high-performance command-line tool for verifying the liveness of IPs and domains, checking open ports, and performing large-scale network probes. It is designed to handle 80,000+ targets efficiently and is inspired by tools like httpx, focusing on scalability and speed. Author: Cristopher. This tool is intended for educational and legitimate cybersecurity purposes, such as penetration testing on systems you own or have explicit permission to scan. Always comply with laws and obtain permission before scanning any network. NetProbe allows you to verify if an IP or domain is live by checking connectivity on a specific port (default: 80), scan a single IP or domain for open ports within a range, check live status of domains/URLs with HTTP/HTTPS support, process large lists (80k+ IPs/domains) with asynchronous and multi-threaded execution, track progress with a progress bar, output results in JSON or plain text, and customize timeout, threads, and batch sizes for performance tuning.

To install NetProbe, clone the repository and install dependencies using:

```
git clone https://github.com/yourusername/netprobe.git
cd netprobe
pip install -r requirements.txt
pip install .
```

To check if an IP is live on port 80, run `netprobe ip 192.168.1.1`. To check an IP on a specific port, run `netprobe ip 192.168.1.1 --port 443`. To scan ports on an IP, run `netprobe scan-ip 192.168.1.1 --start-port 1 --end-port 1024`. To check if a domain is live, run `netprobe domain example.com --port 80`. To scan multiple targets from a file (supports 80k+ targets), run `netprobe batch --file targets.txt --port 80` (where `targets.txt` contains IPs or domains, one per line).

Available options include `--port PORT` to specify the port (default 80), `--timeout TIMEOUT` to set the connection timeout in seconds (default 2), `--threads THREADS` to set the number of threads for IP/port scans (default 50), `--batch-size SIZE` for the number of targets per batch (default 1000), `--output FILE` to save results to a file, and `--json` to output results in JSON format.

NetProbe requires Python 3.8+ and the dependencies `aiohttp` and `tqdm`. See `requirements.txt` for full details. The project is licensed under the MIT License. Use responsibly: scanning without permission may be illegal. The author is not responsible for misuse.
