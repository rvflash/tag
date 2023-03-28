# Tag

[![GoDoc](https://godoc.org/github.com/rvflash/tag?status.svg)](https://godoc.org/github.com/rvflash/tag)
[![Build Status](https://github.com/rvflash/tag/workflows/build/badge.svg)](https://github.com/rvflash/tag/actions?workflow=build)
[![Code Coverage](https://codecov.io/gh/rvflash/tag/branch/main/graph/badge.svg)](https://codecov.io/gh/rvflash/tag)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/tag?)](https://goreportcard.com/report/github.com/rvflash/tag)


`tag` is a Go package providing methods to deal with tag pattern in any bytes or string content.
By default, a tag is a bracket tag. It is expected surrounded by brackets like `[content]`, 
where `content` is the tag value. 

A dedicated `Tag` structure allows to deal with your own tag (`<b>content</b>`, `(content)`, or whatever).


## Features

1. Provides methods to find and replace one or more tags in bytes or string content. 
2. Provides by default bracket tag management and offers possibility to deal with any other. 
3. Offers `Any` map to handle any values as string and easy usage of standard template package. 


## Examples

For these examples, we share this constant as base content.

```go
const text = `Whoever is [happy] will make [others] [happy] too.`
```

### Find the first tag.

```go
fmt.Println(tag.FindString(text))
// Output: happy
```

### Find all tags.

```go
fmt.Println(tag.FindAllString(text))
// Output: [happy others happy]
```

### Replace all tags with the same value `?` in this slice of bytes.

```go
fmt.Println(string(tag.ReplaceAll([]byte(text), []byte("?"))))
// Output: Whoever is ? will make ? ? too.
```

### Replace all tags by another one, preserving the tag value: `[content]` will become `<b>content</b>`.

```go
bold := tag.Must("<b>", "</b>")
fmt.Println(tag.ReplaceAllString(text, bold.String()))
// Output: Whoever is <b>happy</b> will make <b>others</b> <b>happy</b> too.
```

## Replace all tags by using a function to retrieve the new content by the tag value.

Here we use the `Any`, an alias to `map[string]any` to provide a list of key / value and 
the expected function to get a content by a tag value. `Any.StringFunc` will do this job.

```go
values := tag.Any{
    "happy":  "nasty",
    "others": "anyone",
}
fmt.Println(tag.ReplaceAllStringFunc(text, values.StringFunc))
// Output: Whoever is nasty will make anyone nasty too.
```

## Transform the content to build a string representation ready to be used by standard `[html|text]/template` packages.

Here we assume that Any will be use when executing the template.

```go
fmt.Println(tag.TemplateAllString(text))
// Output: Whoever is {{.StringFunc "happy"}} will make {{.StringFunc "others"}} {{.StringFunc "happy"}} too.
```

This will allow this usage :

```go
tpl, err := template.New("").Parse(tag.TemplateAllString(src))
if err != nil {
	log.Fatal(err)
}
err = tpl.Execute(os.Stdout, tag.Any{
    "happy":  "sad",
    "others": "everyone",
})
if err != nil {
    log.Fatal(err)
}
// Output: Whoever is sad will make everyone sad too.
```