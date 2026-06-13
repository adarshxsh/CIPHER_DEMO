package controllers

import (
	// "bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Reciever() {
	fmt.Println("You are on port: 8090 to recieve file ")

	server, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("Error listening on port 8090:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server is listening on port 8090...")

	conn, err := server.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	// Ensure the output directory exists
	err = os.MkdirAll("output", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create the destination image file
	outFile, err := os.Create("output/img.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// io.Copy automatically reads from the network stream in chunks 
	// and writes to your file until the sender closes the connection (EOF).
	written, err := io.Copy(outFile, conn)
	if err != nil {
		fmt.Println("Error saving stream to file:", err)
		return
	}

	fmt.Printf("File saved successfully! Bytes received: %d\n", written)
}
