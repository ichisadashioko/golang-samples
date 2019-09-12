package main

import "fmt"

func main() {
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}

    // the init and post statements are optional.
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
	fmt.Println(sum)
}
