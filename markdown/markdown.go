package markdown

// Third side packages.
import "github.com/russross/blackfriday"

// Golang packages.
import (
	"io/ioutil"
	"strings"
)

// Base directory.
var baseDir string

// Set base directory.
func SetBaseDir(dir string) {
	baseDir = dir
	if !strings.HasSuffix(baseDir, "/") {
		baseDir += "/"
	}
}

// Get readme file HTML content.
func GetReadme() ([]byte, error) {
	html, err := GetFile("README.md")
	return html, err
}

// Get file content and parse it from Markdown to HTML.
func GetFile(f string) ([]byte, error) {
	md, err := readFile(f)
	if err != nil {
		return md, err
	}

	html := blackfriday.MarkdownBasic(md)
	return html, nil
}

// Read file content.
func readFile(f string) ([]byte, error) {
	d, err := ioutil.ReadFile(baseDir + f)
	return d, err
}
