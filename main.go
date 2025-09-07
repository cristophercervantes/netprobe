package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// xprobe - Advanced Network Scanning Tool
// Author: Christopher

type ScanResult struct {
	Host     string
	Port     int
	Open     bool
	Protocol string
	Service  string
	Status   int
	Error    string
	ResponseTime time.Duration
}

type Scanner struct {
	Timeout    time.Duration
	Concurrent int
	UserAgent  string
	Verbose    bool
}

func NewScanner(timeout time.Duration, concurrent int, verbose bool) *Scanner {
	return &Scanner{
		Timeout:    timeout,
		Concurrent: concurrent,
		UserAgent:  "xprobe/1.0",
		Verbose:    verbose,
	}
}

func (s *Scanner) CheckHost(host string) bool {
	if s.Verbose {
		fmt.Printf("[+] Checking if host %s is alive...\n", host)
	}

	// Try ICMP ping first
	if s.pingHost(host) {
		if s.Verbose {
			fmt.Printf("[+] Host %s is alive (ICMP response)\n", host)
		}
		return true
	}

	// If ICMP is blocked, try common ports
	commonPorts := []int{80, 443, 22, 21, 25, 53, 110, 143, 993, 995}
	results := s.ScanPorts(host, commonPorts)

	for _, result := range results {
		if result.Open {
			if s.Verbose {
				fmt.Printf("[+] Host %s is alive (port %d open)\n", host, result.Port)
			}
			return true
		}
	}

	if s.Verbose {
		fmt.Printf("[-] Host %s appears to be down\n", host)
	}
	return false
}

func (s *Scanner) pingHost(host string) bool {
	// Try with ICMP first
	conn, err := net.DialTimeout("ip4:icmp", host, s.Timeout)
	if err == nil {
		conn.Close()
		return true
	}

	// If ICMP fails, try TCP "ping" on common ports
	conn, err = net.DialTimeout("tcp", net.JoinHostPort(host, "80"), s.Timeout)
	if err == nil {
		conn.Close()
		return true
	}

	conn, err = net.DialTimeout("tcp", net.JoinHostPort(host, "443"), s.Timeout)
	if err == nil {
		conn.Close()
		return true
	}

	return false
}

func (s *Scanner) ScanPorts(host string, ports []int) []ScanResult {
	if s.Verbose {
		fmt.Printf("[+] Scanning %d ports on %s...\n", len(ports), host)
	}

	var results []ScanResult
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, s.Concurrent)
	resultChan := make(chan ScanResult, len(ports))

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			result := s.ScanPort(host, p)
			resultChan <- result
		}(port)
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})
	
	return results
}

func (s *Scanner) ScanPort(host string, port int) ScanResult {
	start := time.Now()
	result := ScanResult{
		Host: host,
		Port: port,
	}

	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, s.Timeout)
	if err != nil {
		result.Open = false
		result.Error = err.Error()
		result.ResponseTime = time.Since(start)
		return result
	}
	defer conn.Close()
	result.Open = true
	result.ResponseTime = time.Since(start)

	// Try to determine service
	result.Service = guessService(port)

	// If it's an HTTP port, try to get status code
	if port == 80 || port == 443 || port == 8080 || port == 8443 {
		s.getHTTPStatus(&result)
	}

	return result
}

func (s *Scanner) getHTTPStatus(result *ScanResult) {
	scheme := "http"
	if result.Port == 443 || result.Port == 8443 {
		scheme = "https"
	}

	url := fmt.Sprintf("%s://%s:%d", scheme, result.Host, result.Port)
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: s.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Create request with context and user agent
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result.Error = err.Error()
		return
	}
	req.Header.Set("User-Agent", s.UserAgent)

	start := time.Now()
	resp, err := client.Do(req)
	result.ResponseTime = time.Since(start)
	
	if err != nil {
		result.Error = err.Error()
		return
	}
	defer resp.Body.Close()

	result.Status = resp.StatusCode
	result.Protocol = "HTTP"
}

func guessService(port int) string {
	services := map[int]string{
		20:    "FTP Data",
		21:    "FTP",
		22:    "SSH",
		23:    "Telnet",
		25:    "SMTP",
		53:    "DNS",
		80:    "HTTP",
		110:   "POP3",
		143:   "IMAP",
		443:   "HTTPS",
		465:   "SMTPS",
		587:   "SMTP Submission",
		993:   "IMAPS",
		995:   "POP3S",
		3306:  "MySQL",
		3389:  "RDP",
		5432:  "PostgreSQL",
		6379:  "Redis",
		27017: "MongoDB",
	}

	if service, exists := services[port]; exists {
		return service
	}
	return "Unknown"
}

func parsePorts(portStr string) ([]int, error) {
	var ports []int
	
	if portStr == "" {
		return nil, fmt.Errorf("no ports specified")
	}
	
	// Handle comma-separated ports
	if strings.Contains(portStr, ",") {
		portList := strings.Split(portStr, ",")
		for _, p := range portList {
			port, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				return nil, fmt.Errorf("invalid port: %s", p)
			}
			if port < 1 || port > 65535 {
				return nil, fmt.Errorf("port out of range: %d", port)
			}
			ports = append(ports, port)
		}
		return ports, nil
	}
	
	// Handle port ranges
	if strings.Contains(portStr, "-") {
		rangeParts := strings.Split(portStr, "-")
		if len(rangeParts) != 2 {
			return nil, fmt.Errorf("invalid port range: %s", portStr)
		}
		
		start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start port: %s", rangeParts[0])
		}
		
		end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end port: %s", rangeParts[1])
		}
		
		if start < 1 || end > 65535 || start > end {
			return nil, fmt.Errorf("invalid port range: %d-%d", start, end)
		}
		
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
		return ports, nil
	}
	
	// Single port
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %s", portStr)
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("port out of range: %d", port)
	}
	
	return []int{port}, nil
}

func printBanner() {
	fmt.Println(`
  __  __           _                  
  \ \/ /_ __ _  _ | |__  ___ _ _ _ __ 
   \  /| '_ \ || || '_ \/ -_) '_| '_ \
   /_/ | .__/\_,_|_.__/\___|_| | .__/
       |_|                     |_|    
          
    Advanced Network Scanning Tool
          Author: Christopher
          Version: 1.0
	`)
}

func main() {
	var (
		host      = flag.String("h", "", "Host to scan (IP or domain)")
		portRange = flag.String("p", "80,443", "Ports to scan (e.g., 80,443 or 1-100)")
		timeout   = flag.Duration("t", 5*time.Second, "Timeout for connections")
		concurrent= flag.Int("c", 100, "Number of concurrent scans")
		checkOnly = flag.Bool("check", false, "Only check if host is alive, don't scan ports")
		verbose   = flag.Bool("v", false, "Verbose output")
		version   = flag.Bool("version", false, "Show version information")
	)
	flag.Parse()

	if *version {
		fmt.Println("xprobe v1.0 - Advanced Network Scanning Tool")
		fmt.Println("Author: Christopher")
		return
	}

	if *host == "" {
		printBanner()
		fmt.Println("Error: host is required")
		fmt.Println("Usage: xprobe -h <host> [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !*verbose {
		printBanner()
	}

	ports, err := parsePorts(*portRange)
	if err != nil {
		fmt.Printf("Error parsing ports: %v\n", err)
		os.Exit(1)
	}

	scanner := NewScanner(*timeout, *concurrent, *verbose)

	if *verbose {
		fmt.Printf("[+] Starting scan against %s\n", *host)
	}

	if alive := scanner.CheckHost(*host); !alive {
		if !*verbose {
			fmt.Printf("[-] Host %s appears to be down\n", *host)
		}
		os.Exit(1)
	}

	if *checkOnly {
		if !*verbose {
			fmt.Printf("[+] Host %s is alive\n", *host)
		}
		os.Exit(0)
	}

	results := scanner.ScanPorts(*host, ports)

	// Print results
	fmt.Printf("\nScan results for %s:\n", *host)
	fmt.Println("PORT     STATUS    SERVICE       HTTP STATUS  RESPONSE TIME")
	fmt.Println("-----------------------------------------------------------")
	
	openCount := 0
	for _, result := range results {
		status := "CLOSED"
		if result.Open {
			status = "OPEN"
			openCount++
		}
		
		httpStatus := ""
		if result.Status > 0 {
			httpStatus = fmt.Sprintf("%d", result.Status)
		}
		
		responseTime := fmt.Sprintf("%.2fms", float64(result.ResponseTime.Microseconds())/1000)
		
		fmt.Printf("%-8d %-10s %-12s %-12s %s\n", 
			result.Port, status, result.Service, httpStatus, responseTime)
	}
	
	fmt.Printf("\nSummary: %d ports scanned, %d open, %d closed\n", 
		len(results), openCount, len(results)-openCount)
}
