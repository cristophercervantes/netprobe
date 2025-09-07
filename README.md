# NetProbe

NetProbe is a simple command-line tool for verifying the liveness of IPs and domains, checking open ports, and performing basic network probes. It is inspired by tools like **httpx** but focuses on general TCP port checking rather than HTTP-specific features.

> **Note:** This tool is intended for educational and legitimate cybersecurity purposes, such as penetration testing on systems you own or have explicit permission to scan. Always comply with laws and obtain permission before scanning any network.

---

## Features

- Verify if an IP is live by checking connectivity on a specific port (default: 80).
- Scan a single IP for open ports within a range.
- Check if a domain or URL is live by resolving it and checking a port.
- Scan multiple domains or IPs from a file or list.
- Multi-threaded for faster scans.
- Output results in **JSON** or **plain text**.

---

## Installation

You can install NetProbe via pip once it's published, but for now, clone the repository and install locally:

```bash
git clone https://github.com/cristophercervantes/netprobe.git
cd netprobe
pip install .
