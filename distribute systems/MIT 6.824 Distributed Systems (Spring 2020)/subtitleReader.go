package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	count := 0
	var output []byte
	for s.Scan() {
		count++
		if (count-3)%4 == 0 {
			output = append(output, s.Bytes()...)
			output = append(output, ' ')

		}

	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	// defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write(output)

	}
}
