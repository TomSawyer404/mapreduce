package main

import (
	"fmt"
	"log"
	"map-reduce/pkg/mapreduce"
	"os/exec"
	"time"
)

func main() {
	// go run ./examples/main.go -race
	cmd := exec.Command("fortune")
	buf, err := cmd.Output()
	if err != nil {
		log.Fatalln("cmd.Output() ->", err)
	}

	input := string(buf)
	fmt.Printf("\x1b[31m")
	fmt.Println(input)
	fmt.Printf("\x1b[0m")

	t := time.Now()
	fmt.Println(mapreduce.Run(input))
	elapesed := time.Since(t)
	fmt.Println("time elapesed:", elapesed)
}
