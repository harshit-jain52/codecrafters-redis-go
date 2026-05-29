package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err.Error())
			return
		}
		fmt.Println("Received:", string(buf[:n]))

		received_data, err := decodeRESPArray(string(buf[:n]))
		if err != nil {
			fmt.Println("Error decoding RESP array:", err.Error())
			return
		}
		fmt.Println("Decoded data:", received_data)

		if received_data[0] == "PING" {
			_, err = conn.Write([]byte("+PONG\r\n"))
		} else if received_data[0] == "ECHO" && len(received_data) > 1 {
			response := encodeBulkString(received_data[1])
			_, err = conn.Write([]byte(response))
		}
		
		if err != nil {
			fmt.Println("Error writing to connection:", err.Error())
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Server listening on port 6379")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handleClient(conn)
	}
}