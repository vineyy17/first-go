package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = sha1sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

/*

if file names ends with .gz

      $ cat http.log.gz| gunzip | sha1sum

else

    $ cat http.log.gz| sha1sum


*/

func sha1sum(fileName string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close() // defers are called in reverse order
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {

		gz, err := gzip.NewReader(file)

		if err != nil {
			return "", err
		}

		defer gz.Close()
		r = gz

	}

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
