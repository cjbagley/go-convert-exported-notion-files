package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
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

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

        if info.IsDir() {
            return nil
        }

        rawFilename := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
        if rawFilename == "" || rawFilename == "." {
            return nil
        }

        rawFilename = processFilename(rawFilename, info)

        newExt := filepath.Ext(path)
        if newExt == ".md" {
            newExt = ".org"
        }

        err = os.Rename(path, filepath.Join(filepath.Dir(path), rawFilename + newExt))
        if err != nil {
            return err
        }

		return nil
	})

	if err != nil {
		exit(err.Error())
	}
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

func processFilename(filename string, info os.FileInfo) string {
    var newParts []string
    for _, p := range strings.Split(filename, " ") {
       // Remove the notion hashes
       if utf8.RuneCountInString(p) == 32 {
           continue
       }
       newParts = append(newParts, strings.ToLower(p))
    }

    return strings.ReplaceAll(strings.Join(newParts, "-"), "---", "-") 
}
