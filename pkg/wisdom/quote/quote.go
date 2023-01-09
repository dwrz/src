package quote

import (
	"strings"

	"code.dwrz.net/src/pkg/text"
)

type Quote struct {
	Author  string              `json:"author"`
	Comment string              `json:"comment"`
	Source  string              `json:"source"`
	Tags    map[string]struct{} `json:"tags"`
	Text    string              `json:"text"`
}

func (q *Quote) Render(width int) string {
	var str strings.Builder

	str.WriteString(text.Wrap(q.Text, width))

	if q.Author != "" || q.Source != "" {
		str.WriteString("\n\n")
	}
	if q.Author != "" {
		str.WriteString(text.Wrap(q.Author, width))
	}
	if q.Author != "" && q.Source != "" {
		str.WriteString(", ")
	}
	if q.Source != "" {
		str.WriteString(text.Wrap(q.Source, width))
	}
	if q.Comment != "" {
		str.WriteString("\n\n")
		str.WriteString(text.Wrap(q.Comment, width))
	}
	if len(q.Tags) > 0 {
		str.WriteString("\nTags: ")
		var tags string
		for tag := range q.Tags {
			tags += tag + " "
		}
		str.WriteString(text.Wrap(tags, width))
	}

	return str.String()
}
