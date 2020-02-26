	
	/*
# Authors : [rajeev.chinnikanu@colorado.edu]
# Name   : Programming Assignment 1.
# Purpose: UDPServer
# Date   : Feb2020
# Version : Go 1.13
*/


package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		CONNECT := arguments[1]
		fmt.Println(CONNECT)
	}
	// net.ListenUDP()
	// s, err := net.ResolveUDPAddr("udp4", CONNECT)
	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0:3333")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 2048)
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		Ti2 := strconv.FormatInt(time.Now().UnixNano(), 10)
		Ti3 := string(buffer[0:n])
		fmt.Println("Ti3 -> ", Ti3)
		fmt.Println("Ti2 >> ", Ti2)
		_ = addr
		_ = err

		Ti1 := strconv.FormatInt(time.Now().UnixNano(), 10)
		reader := Ti3 + "???" + Ti2 + "???" + Ti1
		// fmt.Println(reader)
		// fmt.Println(">> Ti1:", Ti1)
		data := []byte(reader)
		_, err = connection.WriteToUDP(data, addr)
	}
}

