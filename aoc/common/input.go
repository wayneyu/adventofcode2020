package common

import (
    "bufio"
    "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_input(path string) []string {

    file, err := os.Open(path)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Read numbers from input file
    out := make([]string, 0)
    for scanner.Scan() {
        out = append(out, scanner.Text())
    }

    return out
}