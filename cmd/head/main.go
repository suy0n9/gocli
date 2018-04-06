package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func run() error {
	var (
		number  int
		verbose bool
	)
	flag.IntVar(&number, "n", 10, "number flag")
	flag.BoolVar(&verbose, "v", false, "print file")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		return fmt.Errorf("need args")
	}

	file := args[0]
	if verbose {
		fmt.Printf("==> %s <==\n", file)
	}
	fp, err := os.Open(file)

	if err != nil {
		return fmt.Errorf("not read file: %s", file)
	}
	defer fp.Close()

	i := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		i++
		if i == number {
			break
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

}
