package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func handleConnections(port string, wg *sync.WaitGroup) {
	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
LOOP:
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("serving remote address %s\n\n", c.RemoteAddr().String())
		for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
				break
			}
			c.Write([]byte("Hi There from server\n"))
			c.Close()
			break LOOP
		}
	}
	fmt.Println("closing listner")
	wg.Done()

}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go handleConnections("1337", &wg)

	wg.Add(1)
	go handleConnections("1338", &wg)

	wg.Add(1)
	go handleConnections("1339", &wg)

	wg.Add(1)
	go handleConnections("1340", &wg)

	wg.Wait()
	fmt.Println("server done handling connections")

}
