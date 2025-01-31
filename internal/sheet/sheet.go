package sheet

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/tj/front"
)

// frontmatter is an un-exported helper struct used in parsing cheatsheets
type frontmatter struct {
	Tags   []string
	Syntax string
}

// Sheet encapsulates sheet information
type Sheet struct {
	Title    string
	Path     string
	Text     string
	Tags     []string
	Syntax   string
	ReadOnly bool
}

// New initializes a new Sheet
func New(
	title string,
	path string,
	tags []string,
	readOnly bool,
) (Sheet, error) {

	// read the cheatsheet file
	markdown, err := ioutil.ReadFile(path)
	if err != nil {
		return Sheet{}, fmt.Errorf("failed to read file: %s, %v", path, err)
	}

	// parse the front-matter
	var fm frontmatter
	text, err := front.Unmarshal(markdown, &fm)
	if err != nil {
		return Sheet{}, fmt.Errorf("failed to parse front-matter: %v", err)
	}

	// merge the sheet-specific tags into the cheatpath tags
	tags = append(tags, fm.Tags...)

	// sort strings so they pretty-print nicely
	sort.Strings(tags)

	// initialize and return a sheet
	return Sheet{
		Title:    title,
		Path:     path,
		Text:     strings.TrimSpace(string(text)) + "\n",
		Tags:     tags,
		Syntax:   fm.Syntax,
		ReadOnly: readOnly,
	}, nil
}
