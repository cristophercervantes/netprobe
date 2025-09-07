# xprobe - Advanced Network Scanning Tool

![xprobe](https://img.shields.io/badge/xprobe-Network%2520Scanner-blue)
![Go](https://img.shields.io/badge/Go-1.21%252B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Platform](https://img.shields.io/badge/Platform-Linux-lightgrey)

**xprobe** is a powerful, high-performance network scanning tool written in Go.  
It provides comprehensive network reconnaissance capabilities with a focus on **speed, accuracy, and ease of use.**

---

## 🔥 New in Version 1.1
- **Multiple Target Support**: Scan multiple IPs/domains from a text file  
- **Improved Output Formatting**: Better organized results for multiple targets  
- **File Input Handling**: Support for comments in target files (lines starting with `#`)  

---

## ✨ Features
- 🚀 **Fast Concurrent Scanning**: Utilizes Go's goroutines for high-speed parallel scanning  
- 🌐 **Multiple Protocol Support**: TCP port scanning, HTTP/HTTPS service detection  
- 📊 **Comprehensive Results**: Port status, service detection, HTTP status codes, response times  
- 🔍 **Host Discovery**: Multiple methods to determine host availability (ICMP, TCP)  
- ⚡ **Performance Metrics**: Measures and displays connection response times  
- 🎯 **Flexible Targeting**: Support for IP addresses, domain names, and custom port ranges  
- 📝 **Detailed Reporting**: Clean, formatted output with scan summaries  
- 🔧 **Configurable**: Adjustable timeouts, concurrency levels, and verbosity  
- 📁 **Batch Scanning**: Scan multiple targets from a file  

---

## ⚙️ Installation

### Prerequisites
- Go 1.21 or later  
- Linux (Ubuntu/Debian recommended)  

### Quick Install
```bash
# Clone the repository
git clone https://github.com/Christopher/xprobe.git
cd xprobe

# Build and install
make install
```

## Manual Installation
```
## Build the binary
go build -o xprobe .

## Make it executable
chmod +x xprobe

## Install to system path (optional)
sudo mv xprobe /usr/local/bin/
```

## Using Makefile
```
make build    # Build the binary
make install  # Install system-wide
make clean    # Remove built binaries
make test     # Run tests (if available)
make version  # Display version information
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






