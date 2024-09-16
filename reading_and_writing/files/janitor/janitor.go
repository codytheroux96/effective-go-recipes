/* A utility that goes over a given directory and compresses any log file that has .log suffix that is over a month old.
After compressing, you compare the content of the compressed file with the original file and if it matches, delete the original file
*/

package janitor

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"time"
)

// gzCompress compresses src to dest with gzip compress
func gzCompress(src, dest string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	w := gzip.NewWriter(out)
	defer w.Close()

	//Update metadata, must be before io.Copy
	w.Name = src
	info, err := file.Stat()
	if err == nil {
		w.ModTime = info.ModTime()
	}

	if _, err := io.Copy(w, file); err != nil {
		os.Remove(dest)
		return err
	}

	return nil
}

// shouldCompress checks if a file is older than a given time span (in this case, 30 days)
func shouldCompress(path string, maxAge time.Duration) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.Printf("warning: %q: can't get info: %s", path, err)
		return false
	}

	if info.IsDir() {
		return false
	}

	return time.Since(info.ModTime()) >= maxAge
}

// filesToCompress will return a list of files that are older than a given time span in a directory
func filesToCompress(dir string, maxAge time.Duration) ([]string, error) {
	root := os.DirFS(dir)
	logFiles, err := fs.Glob(root, "*.log")
	if err != nil {
		return nil, err
	}

	var names []string
	for _, src := range logFiles {
		name := path.Join(dir, src)
		if shouldCompress(name, maxAge) {
			names = append(names, name)
		}
	}
	return names, nil
}

// Checking the file was compressed without issues by comparing the file SHA1 signature
func fileSHA1(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	var r io.Reader = file
	if path.Ext(filename) == ".gz" {
		var err error
		r, err = gzip.NewReader(r)
		if err != nil {
			return "", err
		}
	}

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := fmt.Sprintf("%x", w.Sum(nil))
	return sig, nil
}

// Func to compare SHA1 signatures
func sameSig(file1, file2 string) (bool, error) {
	sig1, err := fileSHA1(file1) 
	if err != nil {
		return false, err
	}

	sig2, err := fileSHA1(file2)
	if err != nil {
		return false, nil
	}

	return sig1 == sig2, nil
}
