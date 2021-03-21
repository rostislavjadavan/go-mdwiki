package main

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

var defaultFilePermission fs.FileMode = 0644

var renderer markdown.Renderer
var filenameValidationRegexp *regexp.Regexp

func init() {
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer = html.NewRenderer(opts)

	filenameValidationRegexp, _ = regexp.Compile("^[\\w-\\\\.]+[^\\W]{1}$")
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

func ToMarkdown(content []byte) string {
	contentString := strings.NewReplacer("\r\n", "\n").Replace(string(content))
	return string(markdown.ToHTML([]byte(contentString), nil, renderer))
}

func CreateNewPage(filename string) (*Page, error) {
	if !strings.HasSuffix(filename, ".md") {
		filename = filename + ".md"
	}
	err := ioutil.WriteFile(path.Join(config.Storage, "pages", filename), []byte("# "+filename), defaultFilePermission)
	if err != nil {
		return nil, err
	}
	return LoadPage(filename)
}

type Page struct {
	Filename   string
	Content    string
	RawContent string
}

func LoadPage(filename string) (*Page, error) {
	content, err := ioutil.ReadFile(path.Join(config.Storage, "pages", filename))
	if err != nil {
		return nil, err
	}
	return &Page{
		Filename:   filename,
		Content:    ToMarkdown(content),
		RawContent: string(content[:]),
	}, nil
}

func (p *Page) updateMarkdown() {
	p.Content = ToMarkdown([]byte(p.RawContent))
}

func (p *Page) Render() string {
	out, err := Render(p.Content, nil)
	if err != nil {
		return err.Error()
	}
	return out
}

func (p *Page) Delete() error {
	return os.Remove(path.Join(config.Storage, "pages", p.Filename))
}

func (p *Page) UpdateContent(content string) error {
	err := ioutil.WriteFile(path.Join(config.Storage, "pages", p.Filename), []byte(content), defaultFilePermission)
	if err != nil {
		return err
	}
	p.RawContent = content
	p.updateMarkdown()
	return nil
}

func (p *Page) Rename(newFilename string) error {
	err := os.Rename(path.Join(config.Storage, "pages", p.Filename), path.Join(config.Storage, "pages", newFilename))
	if err != nil {
		return err
	}
	p.Filename = newFilename
	return nil
}

type PageInfo struct {
	Filename string
	ModTime  time.Time
}

func ListPages() ([]PageInfo, error) {
	dir, err := os.Open(path.Join(config.Storage, "pages"))
	if err != nil {
		return nil, err
	}

	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	list := make([]PageInfo, 0)
	for _, file := range files {
		if !file.IsDir() && ValidateFilename(file.Name()) {
			list = append(list, PageInfo{
				Filename: file.Name(),
				ModTime:  file.ModTime(),
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Filename < list[j].Filename
	})

	return list, nil
}
