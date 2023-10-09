package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
    flag.Parse()
    dir := flag.Arg(0)
    
    if dir == "" {
        exit("Please specify a directory")
    }

    isDir, err := isValidDirectory(dir)
    if !isDir || err != nil {
        exit("Not a directory")
    }

    fmt.Println("okay")
}

func exit(message string) {
    fmt.Println(message)
    os.Exit(1)
}

func isValidDirectory(path string) (bool, error) {
    info, err := os.Stat(path)
    if err != nil {
        return false, err
    }

    return info.IsDir(), nil
}
