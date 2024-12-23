package server

import (
	"TCP-echo-server/config"
	"io"
	"log"
	"net"
	"strconv"
)

func readCommand(c net.Conn) (string, error) {
	// read the command from the client
	// Max Read size is 1024 bytes
	cmd := make([]byte, 1024)
	n, err := c.Read(cmd[:])
	if err != nil {
		return "", err
	}
	return string(cmd[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func RunSyncTCPServer() { // It is a synchronous server, it will handle one connection at a time
	log.Println("Starting a synchronous TCP server on", config.Host, config.Port)

	var con_clients int = 0 // hold current number of connected clients to the server -> 1 always as it is a synchronous server

	// listening to the configured host and port
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	// Start an infininte loop to accept incoming connections
	for {
		// blocking call: waiting for the new client to connnect
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		// increment the number of connected clients
		con_clients += 1
		log.Println("Client connected. Total clients:", con_clients, "Address:", c.RemoteAddr())

		// Another infinite loop to read data from the client
		for {
			// over the socket continuously read the command and print it out
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("Client disconnected. Total clients:", con_clients, "Address:", c.RemoteAddr())
				if err == io.EOF { // client disconnected on its own
					break
				}
				log.Println("Error reading command. Error:", err)
			}
			log.Println("Command received:", cmd)
			if err =  respond(cmd, c); err != nil {
				log.Println("Error responding to client. Error:", err)
			}
		}
	}

}
