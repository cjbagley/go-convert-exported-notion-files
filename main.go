package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
    readFile := flag.Arg(0)
    if readFile == "" {
        exit("Please specify a directory")

    }
}

func exit(message string) {
    fmt.Println(message)
    os.Exit(1)
}
