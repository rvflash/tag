// Package tag provides methods to deal with tags pattern in any bytes or string content.
// A tag by default is a bracket tag.
// It is expected surrounded by brackets like `[content]`, where `content` is the tag value.
// A dedicated Tag structure allows to deal with your own tag (`<b>content</b>`, `(content)`, etc.).
package tag

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

const value = `(.*?)`

func quote(begin, end string) string {
	return regexp.QuoteMeta(begin) + value + regexp.QuoteMeta(end)
}

// Make tries to return a new instance of Tag based on these begin and end values.
func Make(begin, end string) (*Tag, error) {
	rgx, err := regexp.Compile(quote(begin, end))
	if err != nil {
		return nil, err
	}
	return &Tag{
		begin: begin,
		end:   end,
		regex: rgx,
	}, nil
}

// Must returns a new instance of Tag based on these begin and end values.
// It panics if it failed to make it.
func Must(begin, end string) *Tag {
	tag, err := Make(begin, end)
	if err != nil {
		panic(fmt.Sprintf(`tag: make with %q as beginning and ending with %q: %s`, begin, end, err))
	}
	return tag
}

// Tag represents a tag.
type Tag struct {
	begin string
	end   string
	regex *regexp.Regexp
}

// Find returns a slice holding the tag value of the leftmost match in b.
// A return value of nil indicates no match.
func (t Tag) Find(b []byte) []byte {
	return t.trim(t.regex.Find(b))
}

// FindAll is the 'All' version of Find.
// It returns a slice of all tag values found.
// A return value of nil indicates no match.
func (t Tag) FindAll(b []byte) [][]byte {
	res := t.regex.FindAll(b, -1)
	for k, v := range res {
		res[k] = t.trim(v)
	}
	return res
}

func (t Tag) trim(b []byte) []byte {
	return bytes.TrimSuffix(bytes.TrimPrefix(b, []byte(t.begin)), []byte(t.end))
}

// FindAllString is the 'All' version of FindString.
// It returns a slice of all tag values.
// A return value of nil indicates no match.
func (t Tag) FindAllString(s string) []string {
	res := t.regex.FindAllString(s, -1)
	for k, v := range res {
		res[k] = t.trimString(v)
	}
	return res
}

// FindString returns a string holding the tag value of the leftmost match in s.
// If there is no match, the return value is an empty string.
func (t Tag) FindString(s string) string {
	return t.trimString(t.regex.FindString(s))
}

func (t Tag) trimString(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, t.begin), t.end)
}

// ReplaceAll returns a copy of `in`, replacing all tags with the replacement text `by`.
// Inside repl, `$` signs are interpreted as in Expand,
// so for instance `$1` represents the value of the first tag.
func (t Tag) ReplaceAll(in, by []byte) []byte {
	return t.regex.ReplaceAll(in, by)
}

// ReplaceAllFunc returns a copy of `in` in which all tags have been replaced
// by the return value of function  `by` applied to the matched tag value.
func (t Tag) ReplaceAllFunc(in []byte, by func([]byte) []byte) []byte {
	return t.regex.ReplaceAllFunc(in, func(b []byte) []byte {
		return by(t.trim(b))
	})
}

// ReplaceAllString returns a copy of `in`, replacing all tags with the replacement string `by`.
// Inside repl, `$` signs are interpreted as in Expand,
// so for instance `$1` represents the value of the first tag.
func (t Tag) ReplaceAllString(in, by string) string {
	return t.regex.ReplaceAllString(in, by)
}

// ReplaceAllStringFunc returns a copy of `in` in which all tags have been replaced
// by the return value of function `by` applied to the matched tag value.
func (t Tag) ReplaceAllStringFunc(in string, by func(string) string) string {
	return t.regex.ReplaceAllStringFunc(in, func(s string) string {
		return by(t.trimString(s))
	})
}

// String returns a generic placeholder representation of the tag.
// It implements the fmt.Stinger interface.
func (t Tag) String() string {
	return fmt.Sprintf("%s${1}%s", t.begin, t.end)
}

// TemplateAll returns a copy of `in` in which all tags have been replaced
// to build a bytes representation ready to be used by the standard html or text template packages,
// giving access to the values of Any map.
func (t Tag) TemplateAll(in []byte) []byte {
	return t.ReplaceAll(in, []byte(template.String()))
}

// TemplateAllString returns a new string based on `in` in which all tags have been replaced
// to build a string representation ready to be used by the standard html or text template packages,
// giving access to the values of Any map.
func (t Tag) TemplateAllString(in string) string {
	return t.ReplaceAllString(in, template.String())
}

// Any is map of string containing any values.
// The func method attaches to it allow to use this list of values as string replacements.
type Any map[string]any

// Func provides the bytes replacement function as expected by ReplaceAllFunc.
// It returns the bytes representation of the value behind this tag text.
// Example, with a bracket's tag like this one: `[content]`, the value is `content`.
func (m Any) Func(tag []byte) []byte {
	return []byte(m.StringFunc(string(tag)))
}

// StringFunc provides the string replacement function as expected by ReplaceAllStringFunc.
// It returns the string representation of the value behind this tag string.
// Example, with a bracket's tag like this one: `[content]`, the value is `content`.
func (m Any) StringFunc(tag string) string {
	return fmt.Sprintf("%v", m[tag])
}
