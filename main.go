package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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

	var wg sync.WaitGroup // Wait group used to not exit prematurely before all goroutines finish
	var openPorts []int   // Array of open ports that we can sort and display nicely later
	var mu sync.Mutex     // Mutex used to prevent goroutines from overwriting eachother in arrays

	timeout := time.Second

	for currentPort := firstPort; currentPort <= lastPort; currentPort++ {
		portString := strconv.Itoa(currentPort)

		wg.Go(func() {
			conn, _ := net.DialTimeout("tcp", net.JoinHostPort(ip, portString), timeout)
			if conn != nil {
				conn.Close()
				mu.Lock()
				openPorts = append(openPorts, currentPort)
				mu.Unlock()
			}
		})
	}
	wg.Wait()

	sort.Ints(openPorts)
	for _, p := range openPorts {
		fmt.Printf("[open] %d/tcp\n", p)
	}
}
