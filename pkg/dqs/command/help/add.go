package help

import (
	"fmt"
	"sort"
	"strings"

	"code.dwrz.net/src/pkg/dqs/category"
)

var add = strings.ReplaceAll(portions, "%s", "add")

func Add() string {
	var str strings.Builder

	str.WriteString(add)

	var abbreviations = [][2]string{}
	for abbreviation, name := range category.Abbreviations {
		abbreviations = append(abbreviations, [2]string{
			name, abbreviation,
		})
	}

	sort.Slice(abbreviations, func(i, j int) bool {
		return abbreviations[i][0] < abbreviations[j][0]
	})

	for _, a := range abbreviations {
		str.WriteString(fmt.Sprintf("%-5s %s\n", a[1], a[0]))
	}

	return str.String()
}
