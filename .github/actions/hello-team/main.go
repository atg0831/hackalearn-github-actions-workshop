package main

import (
	"fmt"
	"os"
)

func main() {
	leader := os.Getenv("INPUT_LEADER")
	member := os.Getenv("INPUT_MEMBER")

	fmt.Printf("Hello %s", leader)
	fmt.Printf("Hello %s", member)
}
