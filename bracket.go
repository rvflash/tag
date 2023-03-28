package tag

var (
	bracket  = Must("[", "]")
	template = Must(`{{.StringFunc "`, `"}}`)
)

// Find implements the Tag.Find method on the default bracket tag.
func Find(b []byte) []byte {
	return bracket.Find(b)
}

// FindAll implements the Tag.FindAll method on the default bracket tag.
func FindAll(b []byte) [][]byte {
	return bracket.FindAll(b)
}

// FindAllString implements the Tag.FindAllString method on the default bracket tag.
func FindAllString(s string) []string {
	return bracket.FindAllString(s)
}

// FindString implements the Tag.FindString method on the default bracket tag.
func FindString(s string) string {
	return bracket.FindString(s)
}

// ReplaceAll implements the Tag.ReplaceAll method on the default bracket tag.
func ReplaceAll(in, by []byte) []byte {
	return bracket.ReplaceAll(in, by)
}

// ReplaceAllFunc implements the Tag.ReplaceAllFunc method on the default bracket tag.
func ReplaceAllFunc(in []byte, by func([]byte) []byte) []byte {
	return bracket.ReplaceAllFunc(in, by)
}

// ReplaceAllString implements the Tag.ReplaceAllString method on the default bracket tag.
func ReplaceAllString(in, by string) string {
	return bracket.ReplaceAllString(in, by)
}

// ReplaceAllStringFunc implements the Tag.ReplaceAllStringFunc method on the default bracket tag.
func ReplaceAllStringFunc(in string, by func(string) string) string {
	return bracket.ReplaceAllStringFunc(in, by)
}

// TemplateAll implements the Tag.TemplateAll method on the default bracket tag.
func TemplateAll(in []byte) []byte {
	return bracket.TemplateAll(in)
}

// TemplateAllString implements the Tag.TemplateAllString method on the default bracket tag.
func TemplateAllString(in string) string {
	return bracket.TemplateAllString(in)
}
