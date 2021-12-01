package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	yearPtr := flag.String("year", "2042", "A year of advent of code")
	dayPtr := flag.String("day", "00", "A day of advent of code")

	flag.Parse()

	fmt.Println("Year:", *yearPtr)
	fmt.Println("Day:", *dayPtr)

	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	puzzleDir := filepath.Join(currentDir, *yearPtr, *dayPtr)
	filePath := filepath.Join(puzzleDir, "solution.go")
	inputPath := filepath.Join(puzzleDir, "INPUT")

	fmt.Println("go", "run", filePath, inputPath)
	output, err :=  exec.Command("go", "run", filePath, " ", inputPath).Output()
	if err != nil {
		fmt.Println(err)
		log.Fatal(output)
		//panic(err)
	}



	fmt.Println("Solution:", string(output))
}
