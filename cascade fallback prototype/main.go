package main

// A cascaded fallback with a hard timeout.

import (
	"fmt"

	"os"
	"time"
)

func input_read(ch chan string) {
	var buffer string
	fmt.Println("Enter the string : ")
	fmt.Scanf("%s", &buffer)
	ch <- buffer
}

func file_read(ch chan string) {
	file, err := os.ReadFile("/home/blazex/Documents/go/learning/temp.txt")
	if err != nil {
		fmt.Println("Failed to open the file ")
		return
	}
	ch <- string(file)
}

func main() {
	fmt.Println("\033[1m", "\t time based reading ", "\033[0m")
	inputch := make(chan string)
	filech := make(chan string)
	// 500ms timeout
	timeout := time.After(5000 * time.Millisecond)
	exittime := time.After(10000 * time.Millisecond)
	go input_read(inputch)

	for {
		select {
		case input := <-inputch:
			fmt.Printf("Data readed from stdin : %s\n", input)
			return
		case input := <-filech:
			fmt.Println(input)
			return
		case <-timeout:
			fmt.Println("[TIMEOUT] 5sec esclapsed | running the file reading ")
			go file_read(filech)
		case <-exittime:
			fmt.Println("NO event completed ,Program is exiting ")
			return
		}
	}
}
