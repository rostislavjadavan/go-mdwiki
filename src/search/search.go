package search

import (
	"github.com/rostislavjadavan/mdwiki/src/search/fuzzy"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	stripmd "github.com/writeas/go-strip-markdown"
	"sort"
	"strings"
	"time"
)

var charsAroundSearchMatch int = 50

type Result struct {
	Query string
	Pages []pageSearchResult
}

type pageSearchResult struct {
	Filename string
	ModTime  time.Time
	Score    int
	Preview  string
}

func Search(query string, s *storage.Storage) (*Result, error) {
	pages, err := s.ListPages()
	if err != nil {
		return nil, err
	}

	query = strings.TrimSpace(query)
	result := Result{
		Query: query,
		Pages: make([]pageSearchResult, 0),
	}

	if query == "" {
		return &result, nil
	}

	for _, page := range pages {
		content, _ := s.LoadRawPageContent(page.Filename)
		found, score, preview := searchInMdContent(query, content)
		if found {
			result.Pages = append(result.Pages, pageSearchResult{
				Filename: page.Filename,
				ModTime:  page.ModTime,
				Score:    score,
				Preview:  preview,
			})
		}
	}

	sort.Slice(result.Pages, func(i, j int) bool {
		return result.Pages[i].Score > result.Pages[j].Score
	})

	return &result, nil
}

func searchInMdContent(query string, markdownContent string) (bool, int, string) {
	content := stripmd.Strip(markdownContent)
	words := strings.Fields(content)
	matches := fuzzy.Find(query, words)
	if len(matches) > 0 {
		score := 0
		preview := ""
		for _, match := range matches {
			score += match.Score
			preview += highlightMatchedContent(match, content)
		}
		if score > 0 {
			return true, score, preview
		}
	}
	return false, 0, ""
}

func highlightMatchedContent(match fuzzy.Match, content string) string {
	if match.Score <= 0 {
		return ""
	}

	index := strings.Index(content, match.Str)
	if index == 1 {
		return ""
	}

	indexStart := index - charsAroundSearchMatch
	if indexStart < 0 {
		indexStart = 0
	}
	indexEnd := index + charsAroundSearchMatch
	if indexEnd > len(content)-1 {
		indexEnd = len(content) - 1
	}
	content = content[indexStart:indexEnd]
	content = strings.Replace(content, match.Str, "<b class=\"highlight\">"+match.Str+"</b>", 1)
	return content + "\n"
}
