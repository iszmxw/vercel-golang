package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Scanner 结构体
type Scanner struct {
	IP          string        // 目标IP
	PortRange   string        // 端口范围，如 "20-100"
	Timeout     time.Duration // 超时时间
	Concurrency int           // 并发数
}

// 解析端口范围
func parsePortRange(portRange string) []int {
	parts := strings.Split(portRange, "-")
	startPort, endPort := 1, 65535 // 默认扫描 1-65535

	if len(parts) == 2 {
		startPort, _ = strconv.Atoi(parts[0])
		endPort, _ = strconv.Atoi(parts[1])
	}

	var ports []int
	for port := startPort; port <= endPort; port++ {
		ports = append(ports, port)
	}
	return ports
}

// 端口扫描任务
func (s *Scanner) scanPort(port int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", s.IP, port)
	conn, err := net.DialTimeout("tcp", address, s.Timeout)
	if err == nil {
		results <- address
		conn.Close()
	}
}

// 启动扫描
func (s *Scanner) Run() []string {
	var wg sync.WaitGroup
	results := make(chan string, 1000)
	sem := make(chan struct{}, s.Concurrency) // 并发控制
	ports := parsePortRange(s.PortRange)

	// 启动扫描任务
	for _, port := range ports {
		wg.Add(1)
		sem <- struct{}{} // 控制并发数量
		go func(port int) {
			defer func() { <-sem }() // 释放并发锁
			s.scanPort(port, &wg, results)
		}(port)
	}

	// 关闭通道
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集扫描结果
	var openPorts []string
	for res := range results {
		openPorts = append(openPorts, res)
	}

	return openPorts
}
