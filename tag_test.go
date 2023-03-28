package tag_test

import (
	"strings"
	"testing"
	"text/template"

	"github.com/matryer/is"
	"github.com/rvflash/tag"
)

const (
	happyKey  = "happy"
	happyVal  = "sad"
	othersKey = "others"
	othersVal = "everyone"

	src = `Whoever is [happy] will make [others] [happy] too.`
	dst = `Whoever is sad will make everyone sad too.`
)

func TestTemplateAllString(t *testing.T) {
	t.Parallel()
	are := is.New(t)
	tpl, err := template.New("").Parse(tag.TemplateAllString(src))
	are.NoErr(err) // unexpected parse error
	buf := new(strings.Builder)
	err = tpl.Execute(buf, tag.Any{
		happyKey:  happyVal,
		othersKey: othersVal,
	})
	are.NoErr(err)               // unexpected execute error
	are.Equal(dst, buf.String()) // mismatch result
}
