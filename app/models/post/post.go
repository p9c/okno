package post

import (
	"github.com/p9c/okno/app/jdb"
	"html/template"
	"time"

	stripmd "github.com/writeas/go-strip-markdown"
	"gopkg.in/russross/blackfriday.v2"
)

// Post contains articles and pages used by the CMS
type Post struct {
	ID             string
	Title          string
	Content        []byte
	ContentPreview []byte
	Slug           string
	IsDraft        bool
	Published      bool
	ItemType       string
	Template       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type item struct {
	*jdb.JDB
}

// GetHTMLContent returns the post's markdown content as HTML
func (p *Post) GetHTMLContent() template.HTML {
	str := string(blackfriday.Run(p.Content))
	return template.HTML(str)
}

// GetContentPreview strips markdown from the content string, trims and returns it
func (p *Post) GetContentPreview(max int) string {
	preview := stripmd.Strip(string(p.Content))
	if len(preview) > max {
		preview = preview[:max]
	}
	return preview
}
