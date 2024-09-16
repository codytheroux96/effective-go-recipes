/* A utility that goes over a given directory and compresses any log file that has .log suffix that is over a month old.
After compressing, you compare the content of the compressed file with the original file and if it matches, delete the original file
*/

package janitor

import (
	"compress/gzip"
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