// main entry point for our server.

package main

import "fmt"

func main() {
	fmt.Println("Uhh this is not ready yet but soon.")
	fmt.Println(hello("Yo"))
	for i := 0; i < 10; i++ {

		fmt.Printf("%d * 7428 = %d", i, i*7428)
	}
}

func hello(name string) string {
	return fmt.Sprintf("Hello there %s", name)
}
