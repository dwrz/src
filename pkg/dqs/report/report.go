package report

import (
	"sort"

	"code.dwrz.net/src/pkg/dqs/entry"
	"code.dwrz.net/src/pkg/dqs/stats"
)

type Recommendation struct {
	Name   string
	Points int
}

type Report struct {
	BodyFat stats.Stats
	DQS     stats.Stats
	Entries []*entry.Entry
	Weight  stats.Stats

	Less []*Recommendation
	More []*Recommendation
}

func New(entries []*entry.Entry) *Report {
	var (
		bf  = []stats.DateValue{}
		dqs = []stats.DateValue{}
		w   = []stats.DateValue{}

		unclaimed = map[string]int{}
		negative  = map[string]int{}
	)
	for _, e := range entries {
		bf = append(bf, stats.DateValue{
			Date:  e.Date,
			Value: e.BodyFat,
		})
		dqs = append(dqs, stats.DateValue{
			Date:  e.Date,
			Value: e.Score(),
		})
		w = append(w, stats.DateValue{
			Date:  e.Date,
			Value: e.Weight,
		})

		// Iterate through high quality categories.
		// Collect all unclaimed positive points.
		// Collect all claimed negative points.
		for _, c := range e.Categories {
			for _, p := range c.Portions {
				if p.Points > 0 && !p.Amount.Claimed() {
					unclaimed[c.Name] += p.Points
				}
				if p.Points < 0 && p.Amount.Claimed() {
					negative[c.Name] += p.Points
				}
			}
		}
	}

	// Assemble the recommendations.
	var more = []*Recommendation{}
	var less = []*Recommendation{}
	for name, points := range unclaimed {
		more = append(more, &Recommendation{
			Name:   name,
			Points: points,
		})
	}
	for name, points := range negative {
		less = append(less, &Recommendation{
			Name:   name,
			Points: points,
		})
	}
	sort.Slice(more, func(i, j int) bool {
		return more[i].Points > more[j].Points
	})
	sort.Slice(less, func(i, j int) bool {
		return less[i].Points < less[j].Points
	})

	return &Report{
		BodyFat: stats.New(bf),
		DQS:     stats.New(dqs),
		Weight:  stats.New(w),
		Entries: entries,

		More: more,
		Less: less,
	}
}
