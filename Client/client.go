	
	/*
# Authors : [rajeev.chinnikanu@colorado.edu, srsh9018@colorado.edu ]
# Name   : Programming Assignment 1.
# Purpose: UDPClient
# Date   : Feb2020
# Version : Go 1.13
*/
	
	package main

	import (
		"fmt"
		"net"
		"os"
		"strings"
		"time"
		"strconv"
		"encoding/json"
	)

	func create_file(f string) *os.File{
		file, err := os.Create(f)
		if err!=nil {
			panic(err)
		}
		return file
		
	}

	func main() {
		arguments := os.Args
		fmt.Println("hello")
		Exit_Condition := false 
		if len(arguments) < 2 {
			return
		}
		fmt.Println(arguments[1], arguments[2])
		var CONNECT string = arguments[1]
		var File_1 string = arguments[2]
		s, err := net.ResolveUDPAddr("udp4", CONNECT)
		connect, err := net.DialUDP("udp4", nil, s)

		if err != nil {
			fmt.Println("Some error %v", err)
			return
		}

		file_object := create_file(File_1)
		file_object.Close()

		defer connect.Close()

		writer_client := func()  {

			for {
				if Exit_Condition == true{
					return
				}
	
				Ti3 := strconv.FormatInt(time.Now().UnixNano(), 10)
				fmt.Println(">> Ti3:", Ti3)
				// text, _ := reader.ReadString('\n')
				data := []byte(Ti3)
				_, err = connect.Write(data)
				time.Sleep(1 * time.Minute)
		
				
			}
			
		}

		
		reader_client := func ()  {
			// var list_time_values [] string
			type clocks_values struct {
				Ti3 int64
				Ti2 int64
				Ti1 int64
				Ti int64
			}
			clock_map := make(map[string]clocks_values)
			// array_clock_map := make([]int64, 3)
			defer fmt.Println(clock_map)
			i := 1
			start := time.Now()
			defer func() {
				// result is accessed after it was set to 6 by the return statement
				Exit_Condition = true
			}() 
			
			for {
				var file, err = os.OpenFile(File_1, os.O_RDWR, 0644)
				fmt.Println("Reader inititaed")
				buffer := make([]byte, 2048)
				n, _, err := connect.ReadFromUDP(buffer) // _ is addr
				s := strings.Split(string(buffer[0:n]), "???")
				Ti3, err := strconv.ParseInt(s[0],10,64)
				Ti2, err := strconv.ParseInt(s[1],10,64)
				Ti1, err := strconv.ParseInt(s[2],10,64)
				if err != nil{
					fmt.Println(err)
				}
				Ti := time.Now().UnixNano()
				key := time.Now().String()
				clock_map[key] = clocks_values{Ti3, Ti2, Ti1, Ti}
				jsonString, _ := json.MarshalIndent(clock_map,"	","		")
				file.Write(jsonString)
				err = file.Sync()
				fmt.Println(clock_map)
				i += 1
				elapsed := time.Since(start)
				if elapsed.Minutes() >= 120{
                                        err = file.Sync()
					file.Close()
					return
				}
				file.Close()


			}
				
		}

		go reader_client()
		writer_client()
		


	}

