package node

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var KnownPeers = make(map[string]bool)
var mutex sync.Mutex

func StartNode(args []string) {
	port := "8080"
	isBootstrap := false
	usePruned := false

	for _, arg := range args {
		if arg == "--bootstrap" {
			isBootstrap = true
		}
		if strings.HasPrefix(arg, "--port=") {
			port = strings.TrimPrefix(arg, "--port=")
		}
		if strings.HasPrefix(arg, "--use-pruned") {
			usePruned = true
		}
	}

	if usePruned && !CanUsePruned() {
		fmt.Println("❌ No se puede usar modo pruned aún – la cadena no ha superado los 6GB")
		return
	}

	fmt.Printf("📡 Iniciando nodo en puerto %s\n", port)
	listener, _ := net.Listen("tcp", ":"+port)

	go startServer(listener)
	go discoverNodes(port)

	select {}
}
