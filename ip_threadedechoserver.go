/* Threaded Echo Server
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1201"
	tcpAddr, err := net.ResolveTcpAddr("ipv4", service)
	checkError(err)

	listener, err := listener.Accept()
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// gorutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//close conn
	defer conn.Close()

	var buf [512]byte
	for {
		//read limit 512
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		// write n bytes; read
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
