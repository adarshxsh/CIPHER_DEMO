package controllers

import (
	"fmt"
	"io"
	"net"
	"os"
)

func FileSelect() *os.File { 
	fmt.Println("Searching for demo.txt")
	loadFile, err := os.Open("test/image.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	fmt.Println("File demo.txt found successfully!")

    return loadFile

}


func FileTransfer() {
	fmt.Println("Starting file transfer...")
    // get the file from the file select function 
    file := FileSelect()

    if file != nil {
		fmt.Println("File found")
		defer file.Close()
	} else {
		fmt.Println("File not found")
		return
	}

    // check for the length of the file      
    // 

    fileInfo, err := file.Stat()
    if err != nil {
        fmt.Println("Error getting file info:", err)
        return
    }
    fmt.Println("File size:", fileInfo.Size())

    // check for the name of the file 
    fmt.Println("File name:", fileInfo.Name())


    // get the port to send the file 
    // for testing bydefault port is 8090

    var port int = 8090

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to server")
	defer conn.Close()

    _, err = io.Copy(conn, file)

    if err != nil { 
        fmt.Println("Error while sending file ")
    }
	// send the file to the server 
	fmt.Println("File sent to server")



	fmt.Println("File transfer completed successfully!")
}

func Sender() { 
	fmt.Println("welcome to port : 8085")

	server, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server is listening on port 8085...")

	for {
        conn, err := server.Accept() // accepting the connection from the reciever 
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        fmt.Println("Client connected:", conn.RemoteAddr())
        // Transfer the file to the connected client
        FileTransfer()
        conn.Close()
        break // Exit after handling one client
}
}
