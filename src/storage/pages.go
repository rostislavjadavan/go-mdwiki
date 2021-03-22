package storage

import (
	"github.com/rostislavjadavan/go-mdwiki/src/config"
	"github.com/rostislavjadavan/go-mdwiki/src/ui"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

type Storage struct {
	config *config.AppConfig
}

type Page struct {
	Filename   string
	Content    string
	RawContent string
}

func CreateStorage(config *config.AppConfig) (*Storage, error) {
	return &Storage{config: config}, nil
}

func (s *Storage) CreateNewPage(filename string) (*Page, error) {
	if !strings.HasSuffix(filename, ".md") {
		filename = filename + ".md"
	}
	err := ioutil.WriteFile(path.Join(s.config.Storage, "pages", filename), []byte("# "+filename), defaultFilePermission)
	if err != nil {
		return nil, err
	}
	return s.LoadPage(filename)
}

func (s *Storage) LoadPage(filename string) (*Page, error) {
	content, err := ioutil.ReadFile(path.Join(s.config.Storage, "pages", filename))
	if err != nil {
		return nil, err
	}
	return &Page{
		Filename:   filename,
		Content:    ToMarkdown(content),
		RawContent: string(content[:]),
	}, nil
}

func (p *Page) Render() string {
	out, err := ui.Render(p.Content, nil)
	if err != nil {
		return err.Error()
	}
	return out
}

func (s *Storage) DeletePage(page *Page) error {
	return os.Remove(path.Join(s.config.Storage, "pages", page.Filename))
}

func (s *Storage) UpdatePageContent(content string, page *Page) error {
	err := ioutil.WriteFile(path.Join(s.config.Storage, "pages", page.Filename), []byte(content), defaultFilePermission)
	if err != nil {
		return err
	}
	page.RawContent = content
	page.Content = ToMarkdown([]byte(content))
	return nil
}

func (s *Storage) RenamePage(newFilename string, page *Page) error {
	err := os.Rename(path.Join(s.config.Storage, "pages", page.Filename), path.Join(s.config.Storage, "pages", newFilename))
	if err != nil {
		return err
	}
	page.Filename = newFilename
	return nil
}

type PageInfo struct {
	Filename string
	ModTime  time.Time
}

func (s *Storage) ListPages() ([]PageInfo, error) {
	dir, err := os.Open(path.Join(s.config.Storage, "pages"))
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
