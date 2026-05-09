package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: go-portscan <host> <port-range>\n")
		os.Exit(1)
	}
	ip := os.Args[1]
	if net.ParseIP(ip) == nil {
		fmt.Fprintf(os.Stderr, "Invalid IP: %s\n", ip)
		os.Exit(1)
	}
	ports := strings.Split(os.Args[2], "-")
	if len(ports) != 2 {
		fmt.Fprintf(os.Stderr, "Port range must follow the format X-X")
		os.Exit(1)
	}
	firstPort, err := strconv.Atoi(ports[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Port %v is not an integer", ports[0])
		os.Exit(1)
	}
	lastPort, err := strconv.Atoi(ports[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Port %v is not an integer", ports[1])
		os.Exit(1)
	}
	if firstPort < 1 || lastPort > 65535 || firstPort > lastPort {
		fmt.Fprintf(os.Stderr, "Your port range is incorrect. Ports should be between 1 and 65535 with first port not being larger than last port")
		os.Exit(1)
	}

	fmt.Printf("Scanning %v (ports %v-%v)...\n", ip, firstPort, lastPort)

	for currentPort := firstPort; currentPort <= lastPort; currentPort++ {
		timeout := time.Second
		portString := strconv.Itoa(currentPort)
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort(ip, portString), timeout)
		if conn != nil {
			conn.Close()
			fmt.Printf("[open] %v/tcp\n", currentPort)
		}

	}
}
