package display

import (
	"fmt"
	"strings"
)

// SectionF contains information for a section
type SectionF struct {
	Icon   string   // Icon contains the title for the section emoji
	Header string   // Header is the main text of a section
	Body   []string // Body provides additional details to a section
}

// Section builds a section into a string
func Section(section SectionF) string {
	var str strings.Builder
	switch {
	case section.Header != "" && section.Icon != "":
		str.WriteString(
			fmt.Sprintf("%s %s\n", Emoji(section.Icon), Title(section.Header)),
		)
	case section.Header != "":
		str.WriteString(fmt.Sprintf("%s\n", Indent(Title(section.Header))))
	}
	for _, line := range section.Body {
		str.WriteString(fmt.Sprintf("%s\n", Indent(Secondary(line))))
	}
	return str.String()
}

// Indent adds a fixed left padding to the line
func Indent(line string) string {
	return fmt.Sprintf("  %s", line)
}
