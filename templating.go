package main

import (
	"errors"
	"fmt"
	"strings"
	"text/template"
)

type Variants struct {
	Main *Palette
	Moon *Palette
	Dawn *Palette
}

type Values struct {
	Id          string
	Name        string
	Key         string
	Description string
	Palette     *Palette
	Variant     Variants
}

func RenderTemplate(name string, content string, values *Values) (*string, error) {
	funcs := template.FuncMap{
		"format": formatColor,
		"div":    div,
		"hex":    func() ColorFormat { return FormatHex },
		"rgb":    func() ColorFormat { return FormatRGB },
		"hsl":    func() ColorFormat { return FormatHSL },
	}

	tmpl, err := template.New(name).Funcs(funcs).Parse(content)
	if err != nil {
		msg := fmt.Sprintf("error parsing template: %v", err)
		return nil, errors.New(msg)
	}

	var b strings.Builder
	tmpl.ExecuteTemplate(&b, name, values)

	buf := b.String()
	return &buf, nil
}

func div(a, b any) float64 {
	toF := func(v any) float64 {
		switch n := v.(type) {
		case uint8:
			return float64(n)
		case uint16:
			return float64(n)
		case int:
			return float64(n)
		case float64:
			return n
		default:
			return 0
		}
	}
	return toF(a) / toF(b)
}
