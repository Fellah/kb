package markdown

import (
	// Project packages.
	"github.com/fellah/kb/cache"

	// Third side packages.
	"github.com/russross/blackfriday"

	// Base packages.
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const README = "README.md"

const INDEX = "SUMMARY.md"

// Base directory.
var baseDir string

// Set base directory.
func SetBaseDir(dir string) {
	baseDir = dir
	if !strings.HasSuffix(baseDir, "/") {
		baseDir += "/"
	}
}

// Get index file HTML content.
func GetIndex() []byte {
	return GetFile(baseDir + INDEX)
}

// Get readme file HTML content.
func GetReadme() []byte {
	return GetFile(baseDir + README)
}

// Get file content and parse it from Markdown to HTML.
func GetFile(f string) []byte {
	html := cache.Get(f)
	if html == nil {
		html = GetFileNoCache(f)
	}

	return html
}

// Get file content and parse it from Markdown to HTML without using cache.
func GetFileNoCache(f string) []byte {
	if !filepath.HasPrefix(f, baseDir) {
		f = baseDir + f
	}

	md, err := readFile(f)
	if err != nil {
		log.Println(err)
		return nil
	}

	html := blackfriday.MarkdownBasic(md)

	cache.Set(f, html)

	return html
}

// Read file content.
func readFile(f string) ([]byte, error) {
	d, err := ioutil.ReadFile(f)
	return d, err
}

// Walk by base directory and get html from markdown files.
func Walk(path string, fi os.FileInfo, err error) error {
	if err != nil {
		log.Println(err)
		return err
	}

	if filepath.Ext(path) == ".md" {
		GetFileNoCache(path)
	}

	return nil
}
