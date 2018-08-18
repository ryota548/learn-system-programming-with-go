package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %d\n", rand.Intn(10))
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %s\n", "Hello golang")
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %f\n", rand.Float64()*10)
}
