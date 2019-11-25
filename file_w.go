package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	text1 := "Hello Gold!"
	file1, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file1.Close()
	file1.WriteString(text1)

	fmt.Println("Done.")

	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		fmt.Print(string(data[:n]))

	}
}
