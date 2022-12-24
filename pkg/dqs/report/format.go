package report

import (
	"fmt"
	"strings"

	"code.dwrz.net/src/pkg/color"
)

func (r *Report) Format() string {
	var str strings.Builder

	fmt.Fprintf(&str, "%sEntries%s\n", color.BrightBlack, color.Reset)
	fmt.Fprintf(&str, "Count: %d\n", len(r.Entries))
	if len(r.Entries) > 0 {
		fmt.Fprintf(
			&str,
			"First: %s\n",
			r.Entries[0].Date.Format("2006-01-02"),
		)
		fmt.Fprintf(
			&str,
			"Latest: %s\n",
			r.Entries[len(r.Entries)-1].Date.Format("2006-01-02"),
		)
		str.WriteString("\n")
	}

	if len(r.BodyFat.Values) > 0 {
		fmt.Fprintf(
			&str,
			"%sBody Fat%s\n", color.BrightBlack, color.Reset,
		)
		fmt.Fprintf(
			&str,
			"Max: %.2f on %v\n",
			r.BodyFat.Max.Value,
			r.BodyFat.Max.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(
			&str,
			"Min: %.2f on %v\n",
			r.BodyFat.Min.Value,
			r.BodyFat.Min.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(&str, "Average: %.2f\n", *r.BodyFat.Average)
		fmt.Fprintf(&str, "Median: %.2f\n", *r.BodyFat.Median)
		str.WriteString("\n")
	}

	if len(r.DQS.Values) > 0 {
		fmt.Fprintf(&str, "%sDQS%s\n", color.BrightBlack, color.Reset)
		fmt.Fprintf(
			&str,
			"Max: %.2f on %v\n",
			r.DQS.Max.Value,
			r.DQS.Max.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(
			&str,
			"Min: %.2f on %v\n",
			r.DQS.Min.Value,
			r.DQS.Min.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(&str, "Average: %.2f\n", *r.DQS.Average)
		fmt.Fprintf(&str, "Median: %.2f\n", *r.DQS.Median)
		str.WriteString("\n")
	}

	if len(r.Weight.Values) > 0 {
		fmt.Fprintf(
			&str, "%sWeight%s\n", color.BrightBlack, color.Reset,
		)
		fmt.Fprintf(
			&str,
			"Max: %.2f on %v\n",
			r.Weight.Max.Value,
			r.Weight.Max.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(
			&str,
			"Min: %.2f on %v\n",
			r.Weight.Min.Value,
			r.Weight.Min.Date.Format("2006-01-02"),
		)
		fmt.Fprintf(&str, "Average: %.2f\n", *r.Weight.Average)
		fmt.Fprintf(&str, "Median: %.2f\n", *r.Weight.Median)
		str.WriteString("\n")
	}

	if len(r.More) > 0 || len(r.Less) > 0 {
		fmt.Fprintf(
			&str,
			"%sRecommendations%s\n", color.BrightBlack, color.Reset,
		)
	}
	if len(r.More) > 0 {
		fmt.Fprintf(&str, "You should eat more:\n\n")
		for i := 0; i < 3 && i < len(r.More); i++ {
			rec := r.More[i]
			fmt.Fprintf(
				&str,
				"%s: %d lost points (%.2f per entry).\n",
				rec.Name,
				rec.Points,
				float64(rec.Points)/float64(len(r.Entries)),
			)
		}
	}
	str.WriteString("\n")
	if len(r.Less) > 0 {
		fmt.Fprintf(&str, "You should eat less:\n")
		for i := 0; i < 2 && i < len(r.Less); i++ {
			rec := r.Less[i]
			fmt.Fprintf(
				&str,
				"%s: %d lost points (%.2f per entry).\n",
				rec.Name,
				rec.Points,
				float64(rec.Points)/float64(len(r.Entries)),
			)
		}
	}

	return str.String()
}
