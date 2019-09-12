package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		fmt.Printf("You just enter: %s\n", fmt.Sprintf("%#v", text))

		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello")
		} else if strings.Compare("exit", text) == 0 {
			break
		}
	}
}
