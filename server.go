package main

import (
	"net"
	"fmt"
	"bufio"
	"time"
	"log"
	"os"
)

// Create the session table structure.
type sessions struct {
	clientaddr []string
	uuid  	   []string
}

// Define the Protocol Handshake headers.
var (
	CH = []byte("0x001")
	SH = []byte("0x002")
	CA = []byte("0x003")
	SA = []byte("0x004")
)

func logger(flag string, msg string) string {
	// set log output.
	f, err := os.OpenFile("server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Error, Failed to open log file.")
	}
	defer f.Close()
	log.SetOutput(f)
	// Set the flags to Switch.
	switch flag {
	case "INFO":
		msg := "INFO: " + msg
		log.Println(msg)
	case "WARNING":
		msg := "WARNING: " + msg
		log.Println(msg)
	case "ERROR":
		msg := "ERROR: " + msg
		log.Println(msg)
	}
	fmt.Println(msg)
	return msg
}

func handle(conn net.Conn) string {
	// Announce the connection & Log the Users IP address.
	logger("INFO","Established Connection with Client " + string(conn.RemoteAddr().String()))
	// Set Connection TimeOut Value.
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("Fatal Error.", err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		logger("msg", "Got data from Client, Payload = " + ln)
		header := string(ln)[0:5]
		fmt.Println("Got header from payload: " + header)
	}
	defer conn.Close()

	return logger("INFO","Connection Closed with Client " + string(conn.RemoteAddr().String()))
}

// Main Program Loop.
func main() {
	lsnr, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	defer lsnr.Close()
	for {
		conn, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}
