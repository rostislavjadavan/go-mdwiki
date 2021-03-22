package storage

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

var defaultFilePermission fs.FileMode = 0755
var defaultDirectoryPermission fs.FileMode = 0755

var renderer markdown.Renderer
var filenameValidationRegexp *regexp.Regexp

func init() {
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer = html.NewRenderer(opts)

	filenameValidationRegexp, _ = regexp.Compile("^[\\w-\\\\.]+[^\\W]{1}$")
}

func ToMarkdown(content []byte) string {
	contentString := strings.NewReplacer("\r\n", "\n").Replace(string(content))
	return string(markdown.ToHTML([]byte(contentString), nil, renderer))
}

func UriToPage(uri string) string {
	page := strings.TrimSpace(uri)
	if page == "" {
		page = "home"
	}
	if !strings.HasSuffix(page, ".md") {
		page = page + ".md"
	}
	return page
}

func ValidateFilename(filename string) bool {
	return filenameValidationRegexp.MatchString(filename)
}

func fsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
