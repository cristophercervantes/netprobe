# xprobe - Advanced Network Scanning Tool

![xprobe](https://img.shields.io/badge/xprobe-Network%2520Scanner-blue)
![Go](https://img.shields.io/badge/Go-1.21%252B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Platform](https://img.shields.io/badge/Platform-Linux-lightgrey)

**xprobe** is a powerful, high-performance network scanning tool written in Go.  
It provides comprehensive network reconnaissance capabilities with a focus on **speed, accuracy, and ease of use.**

---

## What's New in Version 1.2

* **Go Install Support**: Now installable directly via `go install github.com/cristophercervantes/xprobe@latest`.
* **Official Release**: This marks the first stable release with version tagging.
* **Updated Repository**: The official home for the project is now at `github.com/cristophercervantes/xprobe`.
* **Enhanced Module Support**: Proper Go module configuration has been implemented for easy installation.
* **Multiple Target Support**: You can now scan multiple IPs/domains by providing a text file.
* **Improved Output Formatting**: Results are better organized and clearer when scanning multiple targets.
* **File Input Handling**: The tool now supports comments in target files; lines starting with `#` will be ignored.


---

## Features

* 🚀 **Fast Concurrent Scanning**: Utilizes Go's goroutines for high-speed parallel scanning.
* 🌐 **Multiple Protocol Support**: Conducts TCP port scanning and HTTP/HTTPS service detection.
* 📊 **Comprehensive Results**: Reports on port status, detected services, HTTP status codes, and response times.
* 🔍 **Host Discovery**: Employs multiple methods (ICMP, TCP) to determine if a host is available.
* ⚡ **Performance Metrics**: Measures and displays connection response times.
* 🎯 **Flexible Targeting**: Accepts IP addresses, domain names, and custom port ranges.
* 📝 **Detailed Reporting**: Provides clean, formatted output with helpful scan summaries.
* 🔧 **Configurable**: Allows for adjustable timeouts, concurrency levels, and verbosity.
* 📁 **Batch Scanning**: Capable of scanning multiple targets listed in a file.

---

## ✅ Prerequisites

Before you begin, ensure you have the following installed on your system:

* **Go**: Version 1.21 or later is required if you plan to build from source.
* **Operating System**: Linux (Ubuntu/Debian is recommended).

---

## 🛠️ Installation

You can install `xprobe` using any of the methods below.

### Option 1: Go Install (Recommended)

If you have Go installed and configured, you can install `xprobe` with a single command:

```bash
# Install the latest version
go install [github.com/cristophercervantes/xprobe@latest](https://github.com/cristophercervantes/xprobe@latest)
```

## Option 2: Download using curl
```
# Create project directory
mkdir xprobe && cd xprobe

# Download all files directly
curl -O https://gist.githubusercontent.com/assistant/raw/xprobe/main.go
curl -O https://gist.githubusercontent.com/assistant/raw/xprobe/go.mod
curl -O https://gist.githubusercontent.com/assistant/raw/xprobe/Makefile

# Build and install
make install
```
## Option 3: Download Pre-compiled Binary
Visit the [Releases page](https://github.com/cristophercervantes/xprobe/releases) to download pre-compiled binaries for various platforms.

## Option 4: Clone and Build
```
# Clone the repository
git clone https://github.com/cristophercervantes/xprobe.git
cd xprobe

# Build and install
make install
```

## 🚀 Usage
### Basic Syntax
```
xprobe -h <host> [options]
xprobe -f <file> [options]
```

## Examples
### Basic host check:
```
xprobe -h example.com -check
```

### Web server scan:
```
xprobe -h example.com -p 80,443,8080,8443
```

### Scan multiple targets from file:
```
xprobe -f targets.txt -p 80,443,22
```

### Full port scan with custom settings:
```
xprobe -h 192.168.1.1 -p 1-1000 -c 500 -t 2s
```

### Verbose scan with detailed output:
```
xprobe -h example.com -v -p 22,80,443
```

### Scan multiple specific ports:
```
xprobe -h target.com -p 21,22,23,25,53,80,110,143,443,993,995
```

## 📌 Command Line Options
```
| Option     | Description                              | Default               |
| ---------- | ---------------------------------------- | --------------------- |
| `-h`       | Target host (IP or domain)               | (required if no file) |
| `-f`       | File containing list of hosts to scan    | (required if no host) |
| `-p`       | Ports to scan (comma-separated or range) | 80,443                |
| `-t`       | Timeout for connections                  | 5s                    |
| `-c`       | Number of concurrent scans               | 100                   |
| `-check`   | Only check if host is alive              | false                 |
| `-v`       | Verbose output                           | false                 |
| `-version` | Show version information                 | false                 |
```

## 📁 File Format

The target file should contain one host (IP or domain) per line.
Lines starting with # are treated as comments and ignored.

Example ```targets.txt```:
```
# Important servers
example.com
192.168.1.1
github.com

# Internal services
10.0.0.5
10.0.0.6
```

## 🎯 Port Specification

• xprobe supports multiple formats for specifying ports:

• Single port: ```80```

• Comma-separated: ```80,443,8080```

• Range: ```1-1000```

• Combination: ```80,443,8000-9000```

## 📊 Output Explanation
xprobe provides clear, formatted output:
```
PORT     STATUS    SERVICE       HTTP STATUS  RESPONSE TIME
-----------------------------------------------------------
80       OPEN      HTTP          200          15.42ms
443      OPEN      HTTPS         301          23.17ms
22       OPEN      SSH                       -
8080     CLOSED    HTTP
```

• PORT: The scanned port number

• STATUS: Whether the port is OPEN or CLOSED

• SERVICE: Common service associated with the port

• HTTP STATUS: HTTP status code (if applicable)

• RESPONSE TIME: Connection response time in milliseconds

## 🔬 Advanced Usage
### Scanning Multiple Hosts
```
# Create a targets file
echo -e "example.com\ngithub.com\n192.168.1.1" > targets.txt

# Scan all targets
xprobe -f targets.txt -p 80,443,22
```

### Integrating with Other Tools
```
# Scan and filter for only open ports
xprobe -h example.com -p 1-1000 | grep OPEN

# Save results to a file
xprobe -h example.com -p 1-1000 > scan_results.txt

# Use in combination with other tools
xprobe -h example.com -p 80,443 | awk '{print $1,$2}' | grep OPEN

# Generate target list from subnet
nmap -sL 192.168.1.0/24 | grep "Nmap scan" | awk '{print $5}' > targets.txt
xprobe -f targets.txt -p 80,443 -check
```

## ⚡ Performance Tips

• Adjust Concurrency: Use -c flag to increase concurrent scans for faster results

• Timeout Settings: Reduce timeout with -t for internal networks, increase for unreliable connections

• Target Specific Ports: Instead of full ranges, target likely ports for faster results

• Verbose Mode: Use -v for debugging, disable for production scans

• Batch Scanning: For large target lists, split into multiple files for parallel execution

## 🛠️ Troubleshooting
### Common Issues
Permission denied errors:
```
# ICMP scanning may need sudo
sudo xprobe -h example.com -check
```

### Host appears down:

• Check network connectivity

• Verify the host is reachable

• Try verbose mode: ```xprobe -h example.com -v```

### Slow scanning:
```
xprobe -h target.com -c 50 -t 10s
```
File not found:
• Ensure file path is correct

• Check permissions: ``` chmod +r targets.txt```

### 🐞 Debugging
Use verbose mode for detailed information:
```
xprobe -h example.com -v -p 80,443
```

## 📂 Development
### Building from Source
```
git clone https://github.com/Christopher/xprobe.git
cd xprobe
go build -o xprobe .
```

## Project Structure
```
xprobe/
├── main.go      # Main application code
├── go.mod       # Go module definition
├── Makefile     # Build automation
└── README.md    # This file
```

## 🤝 Contributing

### Contributions are welcome!

1. Fork the repository

2. Create a feature branch:
```
 git checkout -b feature-name
```

3. Commit your changes:
```
git commit -am 'Add feature'
```

4. Push the branch:
```
git push origin feature-name
```

5. Submit a pull request

# 📜 Changelog

All notable changes to this project will be documented in this file.

---

### v1.2 (Current) - *2025-09-09*

* Added support for installation via `go install`.
* Created the first official, stable release with version tagging.
* Updated repository information to its new official home.
* Enhanced documentation for better clarity.

---

### v1.1

* Added support for scanning multiple targets from a text file.
* Improved the output formatting for multi-target scans.
* Implemented file input handling with support for comments (lines starting with `#`).

## 📜 License

xprobe is released under the MIT License.
See the LICENSE file for details.

## 👨‍💻 Author

Developed by Christopher - Cyber Security Professional

## ⚠️ Disclaimer

xprobe is designed for ethical security testing and network administration.
Always ensure you have proper authorization before scanning any network or system.
The authors are not responsible for any misuse of this tool.

## 📌 Roadmap

### Future planned enhancements for xprobe:

• UDP port scanning support

• SSL/TLS certificate information

• XML/JSON output formats

• Nmap-style service version detection

• Integration with vulnerability databases

• Graphical User Interface (GUI)

• API for integration with other tools

Output to file option

CIDR notation support

# 🎉 Happy Scanning! 🚀






