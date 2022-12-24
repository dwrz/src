package entry

import (
	"fmt"
	"strings"
	"time"

	"code.dwrz.net/src/pkg/dqs/category"
	"code.dwrz.net/src/pkg/dqs/diet"
	"code.dwrz.net/src/pkg/dqs/user"
)

const (
	DateFormat        = "20060102"
	dateDisplayFormat = "2006-01-02"
)

type Entry struct {
	BodyFat    float64                      `json:"bodyFat"`
	Categories map[string]category.Category `json:"categories"`
	Date       time.Time                    `json:"date"`
	Diet       diet.Diet                    `json:"diet"`
	Height     float64                      `json:"height"`
	Note       string                       `json:"note"`
	Weight     float64                      `json:"weight"`
}

func New(date time.Time, u *user.User) *Entry {
	var e = &Entry{
		Categories: u.Diet.Template(),
		Date:       date,
		Diet:       u.Diet,
	}

	var currentDate = time.Now().Format(DateFormat)
	if currentDate == date.Format(DateFormat) {
		e.BodyFat = u.BodyFat
		e.Height = u.Height
		e.Weight = u.Weight
	}

	return e
}

// Key returns the key used to retrieve an entry from the store.
func (e *Entry) Key() string {
	return e.Date.Format(DateFormat)
}

func (e *Entry) Score() (total float64) {
	for _, c := range e.Categories {
		total += c.Score()
	}

	return total
}

func (e *Entry) Category(c string) (*category.Category, error) {
	// Try expanding an abbreviation.
	category, exists := e.Categories[category.Abbreviations[c]]
	if !exists {
		// Check if a lowercase full category name was used.
		category, exists = e.Categories[strings.Title(c)]
		if !exists {
			// Check the full, capitalized name.
			category, exists = e.Categories[c]
			if !exists {
				return nil, fmt.Errorf(
					"category %s not found", c,
				)
			}
		}
	}

	return &category, nil
}
