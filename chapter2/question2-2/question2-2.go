package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	records := [][]string{
		{"one", "two", "three"},
		{"first", "second", "third"},
	}
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()

	writer1 := csv.NewWriter(os.Stdout)
	writer2 := csv.NewWriter(file)
	for _, record := range records {
		if err := writer1.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		if err := writer2.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	writer1.Flush()
	writer2.Flush()
}
