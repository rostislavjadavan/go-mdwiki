package storage

import (
	"github.com/rostislavjadavan/mdwiki/src/config"
	"os"
	"path"
)

type Storage struct {
	config *config.AppConfig
}

type Page struct {
	Filename   string
	Name       string
	Version    int64
	Content    string
	RawContent string
}

func CreateStorage(config *config.AppConfig) (*Storage, error) {
	s := &Storage{config: config}
	initStorage(s)

	return s, nil
}

var starterHomeMdContent = `# Welcome to your personal wiki!

This is your homepage and you can edit it by using links in the top panel.

## Markdown syntax and extensions

- https://www.markdownguide.org/basic-syntax/
- https://github.com/gomarkdown/markdown#extensions

`

func initStorage(s *Storage) error {
	err := os.MkdirAll(path.Join(s.config.Storage, "pages"), defaultDirectoryPermission)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(s.config.Storage, "trash"), defaultDirectoryPermission)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(s.config.Storage, "versions"), defaultDirectoryPermission)
	if err != nil {
		return err
	}

	homeMd, err := fsExists(path.Join(s.config.Storage, "pages", "home.md"))
	if err != nil {
		return err
	}
	if !homeMd {
		page, err := s.PageCreate("home.md")
		if err != nil {
			return err
		}
		s.PageContentUpdate(starterHomeMdContent, page)
	}
	return nil
}
