package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func  main()  {

	//flag for line option
	lines := flag.Bool("l", false, "count number of lines")
	flag.Parse()

	result := count(os.Stdin, *lines)
	fmt.Println(result)
	
}
func count(r io.Reader, lineCount bool) int{
	//read r (stream of bytes)
	scanner := bufio.NewScanner(r)
	if !lineCount{
		// change from lines to words
	scanner.Split(bufio.ScanWords)
	}
	
	counter := 0
	for scanner.Scan() {
		counter += 1
	}
	err := scanner.Err()
	if err != nil {
		return 0
	}


	return counter
}
