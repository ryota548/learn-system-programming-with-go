package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(file, "Write with os.Stdout at %d\n", rand.Intn(10))
	fmt.Fprintf(file, "Write with os.Stdout at %s\n", "Hello golang")
	fmt.Fprintf(file, "Write with os.Stdout at %f\n", rand.Float64()*10)
}
