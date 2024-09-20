/*
- Currently your company emits an HTTP redirect status code to send customers from an old webpage to a new one
- Now you want the ability to check how many customers are still trying to access the old page
- HTTP server log files are being saved in the logs diectory and some are compressed
- You want to see how many HTTP redirect codes you have been throwing
- ***Solution: write a function that gets an io.Reader and returns the total number of lines and the number of lines that have an
HTTP redirect directive
*/

package redirects

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// numRedirects gets an io.Reader and returns the total number of lines and the number of lines that have an HTTP redirect directive
func numRedirects(r io.Reader) (int, int, error) {
	s := bufio.NewScanner(r)
	nLines, nRedirects := 0, 0
	for s.Scan() {
		nLines++
		// Example:
		// 203.252.212.44 - - [01/Aug/1995:03:45:47 -0400]
		// "GET /ksc.html HTTP/1.0" 200 7200
		fields := strings.Fields(s.Text())
		code := fields[len(fields)-2] // code is one before last
		if code[0] == '3' {
			nRedirects++
		}
	}

	if err := s.Err(); err != nil {
		return -1, -1, err
	}
	return nLines, nRedirects, nil
}

func redirect() {
	// Finding out what log files are in the logs directory
	matches, err := filepath.Glob("logs/http-*.log")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// iterating over files and reading them one by one for reading
	nLines, nRedirects := 0, 0
	for _, fileName := range matches {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		var r io.Reader = file
		if strings.HasSuffix(fileName, ".gz") {
			r, err = gzip.NewReader(r)
			if err != nil {
				log.Fatalf("%q - %v", fileName, err)
			}
		}

		nl, nr, err := numRedirects(r)
		if err != nil {
			log.Fatalf("%q - %v", fileName, err)
		}
		nLines += nl
		nRedirects += nr
	}
	fmt.Printf("%d redirects in %d lines\n", nRedirects, nLines)
}
