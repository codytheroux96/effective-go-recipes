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
	"io"
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
		if code[0] == '3'{
			nRedirects++
		}
	}

	if err := s.Err(); err != nil {
		return -1, -1, err
	}
	return nLines, nRedirects, nil
}