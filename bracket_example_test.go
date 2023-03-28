package tag_test

import (
	"fmt"

	"github.com/rvflash/tag"
)

func ExampleFind() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(string(tag.Find([]byte(text))))
	// Output:
	// happy
}

func ExampleFindAll() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Print(tag.FindAll([]byte(text)))
	// Output:
	// [[104 97 112 112 121] [111 116 104 101 114 115] [104 97 112 112 121]]
}

func ExampleFindAllString() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(tag.FindAllString(text))
	// Output:
	// [happy others happy]
}

func ExampleFindString() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(tag.FindString(text))
	// Output:
	// happy
}

func ExampleReplaceAll() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(string(tag.ReplaceAll([]byte(text), []byte("?"))))
	// Output:
	// Whoever is ? will make ? ? too.
}

func ExampleReplaceAllString() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	bold := tag.Must("<b>", "</b>")
	fmt.Println(tag.ReplaceAllString(text, bold.String()))
	// Output:
	// Whoever is <b>happy</b> will make <b>others</b> <b>happy</b> too.
}

func ExampleReplaceAllFunc() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	values := tag.Any{
		"happy":  "nasty",
		"others": "anyone",
	}
	fmt.Println(string(tag.ReplaceAllFunc([]byte(text), values.Func)))
	// Output:
	// Whoever is nasty will make anyone nasty too.
}

func ExampleReplaceAllStringFunc() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	values := tag.Any{
		"happy":  "nasty",
		"others": "anyone",
	}
	fmt.Println(tag.ReplaceAllStringFunc(text, values.StringFunc))
	// Output:
	// Whoever is nasty will make anyone nasty too.
}

func ExampleTemplateAll() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(string(tag.TemplateAll([]byte(text))))
	// Output:
	// Whoever is {{.StringFunc "happy"}} will make {{.StringFunc "others"}} {{.StringFunc "happy"}} too.
}

func ExampleTemplateAllString() {
	const text = `Whoever is [happy] will make [others] [happy] too.`
	fmt.Println(tag.TemplateAllString(text))
	// Output:
	// Whoever is {{.StringFunc "happy"}} will make {{.StringFunc "others"}} {{.StringFunc "happy"}} too.
}
