package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stdin string
	fmt.Scan(&stdin)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{stdin})
}
