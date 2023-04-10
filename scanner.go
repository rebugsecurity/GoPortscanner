// lets get started shall we?

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "example.com"
	ports := []int{21,22,53,80,443,8080}

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, time.Second)

		if err != nil {
			fmt.Printf("%d: closed\n", port)
			continue
		}

		conn.Close()
		fmt.Printf("%d: open\n", port)
	}
}