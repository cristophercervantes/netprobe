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




