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

#Manual Installation
```
## Build the binary
go build -o xprobe .

## Make it executable
chmod +x xprobe

## Install to system path (optional)
sudo mv xprobe /usr/local/bin/
```
