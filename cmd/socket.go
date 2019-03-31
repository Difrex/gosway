package main

import (
	"bufio"
	"fmt"
	"net"
	"os/user"
)

func ListenCTL() {
	ln, err := net.Listen("tcp", "127.0.0.1:7877")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		s, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
	}
}

func SendToCTL(cmd string) {
	conn, err := net.Dial("tcp", "127.0.0.1:7877")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, cmd+"\n")
	defer conn.Close()
}

func userID() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.Uid
}
