package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		var source string
		fmt.Print("Input a word:")
		_, err := fmt.Scan(&source)
		if err != nil {
			fmt.Println("Incorrect input", err)
			continue
		}

		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Translation:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}

//server and main directories are related - they should be ran simultaneously using 2 different terminals
//this directory is not related to the manga_crud project
