package node

import (
	"fmt"
	"net"
	"time"
)

func discoverNodes(basePort string) {
	ip := net.ParseIP("192.168.1.0")

	for i := 1; i <= 254; i++ {
		host := fmt.Sprintf("%d.%d.%d.%d:%s",
			ip[0], ip[1], ip[2], byte(i), basePort)

		go func(addr string) {
			conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Write([]byte("PING\n"))
			conn.Close()

			if isNodeMalicious(addr) {
				fmt.Printf("❌ Nodo malicioso detectado: %s\n", addr)
				return
			}

			mutex.Lock()
			KnownPeers[addr] = true
			mutex.Unlock()
		}(host)
	}
}
