package io_files

import (
	"fmt"
	"io"
	"os"
)

func RunBasic() {
	file, err := os.Create("test.txt")       // create a file
	fileRead, errRead := os.Open("test.txt") // open a file

	fmt.Println("file", file, "error", err)
	fmt.Println("fileRead", fileRead, "errorRead", errRead)

	f1, err := os.OpenFile("sometext.txt", os.O_CREATE, 0666) // OpenFile open or create file if it doesn't exist

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // close the program
	}

	defer file.Close() // necessary close the file after the work
	defer fileRead.Close()
	defer f1.Close()

	fmt.Println(file.Name()) // get file name
}

func RunWrite() *os.File {
	text := "Test message 123"

	file, err := os.OpenFile("text.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Error during program run: ", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println("Error during write data: ", err)
		os.Exit(1)
	}

	fmt.Println("Write is done.")

	return file
}

func RunRead() {
	RunWrite() // for set data

	file, err := os.Open("text.txt")

	if err != nil {
		fmt.Println("Error during open file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 64)

	for { // read file by 64 bit chank
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}

		fmt.Println(string(data[:n]))
	}
}

func RunIoCopy() {
	RunWrite() // for set data

	file, err := os.Open("text.txt")

	if err != nil {
		fmt.Println("Error during the open file :", err)
		panic(1)
	}

	defer file.Close()

	io.Copy(os.Stdout, file) // transfer data from file flow to out flow
}

func RunPrintToFile() {
	file, errOpen := os.OpenFile("text.txt", os.O_WRONLY, 0666)

	if errOpen != nil {
		fmt.Printf("Error during open file %s", errOpen)
		os.Exit(1)
	}

	defer file.Close()

	_, err := fmt.Fprintf(file, "Some new text %d and %s and %f \n", 123, "STRING", 123.456)

	if err != nil {
		fmt.Printf("Error during write to file - %s", err)
		os.Exit(1)
	}

	fmt.Println("Write was successful.")
}

func RunFormattingOut() {
	func() {
		// init data
		var name string = "Tom"
		var age int = 24

		file, err := os.OpenFile("person.txt", os.O_WRONLY, 0666)

		if err != nil {
			fmt.Printf("Error during open file process - %s", err)
			os.Exit(1)
		}

		defer file.Close()

		fmt.Fprintln(file, name)
		fmt.Fprintln(file, age)

		fmt.Println("Write to file complete.")
	}()

	file, err := os.OpenFile("person.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("Error during open file process - %s", err)
		os.Exit(1)
	}

	defer file.Close()

	var name string
	var age int

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}
