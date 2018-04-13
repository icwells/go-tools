// Returns input string of uncertain encoding as a utf-8 encoded string

package toutf

import (
	"bufio"
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
)

func getEncoding(b []byte) string {
	// Returns string encoding
	det := chardet.Detector()
	res, err := det.DetectBest(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\t[ERROR] Indentifying target encoding: %v\n", err)
	}
	return res.Charset
}

func Recode(b []byte) string {
	// Returns utf-8 string
	enc := getEncoding(b)
	r := transform.NewReader(b, charmap.enc.NewDecoder())
	sc := buifio.NewScanner(r)
	return string(sc.Scan())
}
