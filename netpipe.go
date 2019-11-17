package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

var tcpPort = flag.Int64("l", 0, "TCP port to listen")

func printUsage() {
	port := "<PORT>"
	if *tcpPort != 0 {
		port = strconv.FormatInt(*tcpPort, 10)
	}
	fmt.Printf("Usage: %s -l %s <COMMAND> [<ARGS>]\n", os.Args[0], port)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *tcpPort == 0 {
		printUsage()
	}
	if len(flag.Args()) < 1 {
		printUsage()
	}
	command := flag.Args()[0]
	args := flag.Args()[1:]

	addr := ":" + strconv.FormatInt(*tcpPort, 10)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go pipe(conn, command, args)
	}
}

func pipe(conn net.Conn, command string, args []string) {
	defer conn.Close()

	cmd := exec.Command(command, args...)
	cmd.Stdin = conn
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println("error:", err)
	}
}
