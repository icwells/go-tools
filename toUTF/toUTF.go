// Returns input string of uncertain encoding as a utf-8 encoded string

package toutf

import (
	"fmt"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func getEncoding(b []bytes) (string, error) {
	// Returns string encoding
	det := chardet.Detector()
	res, err := det.DetectBest(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\t[ERROR] Indentifying target encoding: %v\n", err)
	}
	return res.Charset
}

func Recode(b []bytes) (string, error) {
	// Returns utf-8 string
	enc := getEncoding(b)
	r := transform.NewReader(b, charmap.enc.NewDecoder())
	return string(r)
}
