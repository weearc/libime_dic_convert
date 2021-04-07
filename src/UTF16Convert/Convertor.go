package UTF16Convert

import (
"bytes"
"fmt"
"io/ioutil"
"log"
"strings"

"golang.org/x/text/encoding/unicode"
"golang.org/x/text/transform"
)
// function from https://stackoverflow.com/questions/15783830/how-to-read-utf16-text-file-to-string-in-golang

// Similar to ioutil.ReadFile() but decodes UTF-16.  Useful when
// reading data from MS-Windows systems that generate UTF-16BE files,
// but will do the right thing if other BOMs are found.
func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an transformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	return decoded, err
}

func Replacement(filename string)(string, error)  {
	data, err := ReadFileUTF16(filename)
	if err != nil {
		log.Fatal(err)
	}
	final := strings.Replace(string(data), "\r\n", "\n", -1)
	return final, nil

}

func main() {
	filename := "test_file.txt"
	final, err := Replacement(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(final)

}