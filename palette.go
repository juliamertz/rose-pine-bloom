package main

import (
	"iter"
	"reflect"
	"strings"
)

type Variant int

const (
	Main Variant = iota
	Moon
	Dawn
)

type Palette struct {
	Base          Color
	Surface       Color
	Overlay       Color
	Muted         Color
	Subtle        Color
	Text          Color
	Love          Color
	Gold          Color
	Rose          Color
	Pine          Color
	Foam          Color
	Iris          Color
	HighlightLow  Color
	HighlightMed  Color
	HighlightHigh Color
}

type VariantMeta struct {
	id          string
	name        string
	appearance  string
	description string
	colors      *Palette
}

var accents = []string{
	"love", "gold", "rose", "pine", "foam", "iris",
}

var (
	MainPalette = Palette{
		Base: Color{
			HSL: hsl(249, 22, 12),
			RGB: rgb(25, 23, 36),
		},
		Surface: Color{
			HSL: hsl(247, 23, 15),
			RGB: rgb(31, 29, 46),
		},
		Overlay: Color{
			HSL: hsl(245, 25, 18),
			RGB: rgb(38, 35, 58),
		},
		Muted: Color{
			HSL: hsl(249, 12, 47),
			RGB: rgb(110, 106, 134),
		},
		Subtle: Color{
			HSL: hsl(248, 15, 61),
			RGB: rgb(144, 140, 170),
		},
		Text: Color{
			HSL: hsl(245, 50, 91),
			RGB: rgb(224, 222, 244),
		},
		Love: Color{
			HSL: hsl(343, 76, 68),
			RGB: rgb(235, 111, 146),
			On:  "text",
		},
		Gold: Color{
			HSL: hsl(35, 88, 72),
			RGB: rgb(246, 193, 119),
			On:  "surface",
		},
		Rose: Color{
			HSL: hsl(2, 55, 83),
			RGB: rgb(235, 188, 186),
			On:  "surface",
		},
		Pine: Color{
			HSL: hsl(197, 49, 38),
			RGB: rgb(49, 116, 143),
			On:  "text",
		},
		Foam: Color{
			HSL: hsl(189, 43, 73),
			RGB: rgb(156, 207, 216),
			On:  "surface",
		},
		Iris: Color{
			HSL: hsl(267, 57, 78),
			RGB: rgb(196, 167, 231),
			On:  "surface",
		},
		HighlightLow: Color{
			HSL: hsl(244, 18, 15),
			RGB: rgb(33, 32, 46),
		},
		HighlightMed: Color{
			HSL: hsl(247, 15, 28),
			RGB: rgb(64, 61, 82),
		},
		HighlightHigh: Color{
			HSL: hsl(245, 13, 36),
			RGB: rgb(82, 79, 103),
		},
	}

	MoonPalette = Palette{
		Base: Color{
			HSL: hsl(246, 24, 17),
			RGB: rgb(35, 33, 54),
		},
		Surface: Color{
			HSL: hsl(248, 24, 20),
			RGB: rgb(42, 39, 63),
		},
		Overlay: Color{
			HSL: hsl(248, 21, 26),
			RGB: rgb(57, 53, 82),
		},
		Muted: Color{
			HSL: hsl(249, 12, 47),
			RGB: rgb(110, 106, 134),
		},
		Subtle: Color{
			HSL: hsl(248, 15, 61),
			RGB: rgb(144, 140, 170),
		},
		Text: Color{
			HSL: hsl(245, 50, 91),
			RGB: rgb(224, 222, 244),
		},
		Love: Color{
			HSL: hsl(343, 76, 68),
			RGB: rgb(235, 111, 146),
			On:  "text",
		},
		Gold: Color{
			HSL: hsl(35, 88, 72),
			RGB: rgb(246, 193, 119),
			On:  "surface",
		},
		Rose: Color{
			HSL: hsl(2, 66, 75),
			RGB: rgb(234, 154, 151),
			On:  "surface",
		},
		Pine: Color{
			HSL: hsl(197, 48, 47),
			RGB: rgb(62, 143, 176),
			On:  "text",
		},
		Foam: Color{
			HSL: hsl(189, 43, 73),
			RGB: rgb(156, 207, 216),
			On:  "surface",
		},
		Iris: Color{
			HSL: hsl(267, 57, 78),
			RGB: rgb(196, 167, 231),
			On:  "surface",
		},
		HighlightLow: Color{
			HSL: hsl(245, 22, 20),
			RGB: rgb(42, 40, 62),
		},
		HighlightMed: Color{
			HSL: hsl(247, 16, 30),
			RGB: rgb(68, 65, 90),
		},
		HighlightHigh: Color{
			HSL: hsl(249, 15, 38),
			RGB: rgb(86, 82, 110),
		},
	}

	DawnPalette = Palette{
		Base: Color{
			HSL: hsl(32, 57, 95),
			RGB: rgb(250, 244, 237),
		},
		Surface: Color{
			HSL: hsl(35, 100, 98),
			RGB: rgb(255, 250, 243),
		},
		Overlay: Color{
			HSL: hsl(25, 36, 92),
			RGB: rgb(242, 233, 225),
		},
		Muted: Color{
			HSL: hsl(254, 9, 61),
			RGB: rgb(152, 147, 165),
		},
		Subtle: Color{
			HSL: hsl(249, 13, 52),
			RGB: rgb(121, 117, 147),
		},
		Text: Color{
			HSL: hsl(248, 19, 40),
			RGB: rgb(87, 82, 121),
		},
		Love: Color{
			HSL: hsl(343, 35, 55),
			RGB: rgb(180, 99, 122),
			On:  "surface",
		},
		Gold: Color{
			HSL: hsl(35, 81, 56),
			RGB: rgb(234, 157, 52),
			On:  "surface",
		},
		Rose: Color{
			HSL: hsl(2, 55, 67),
			RGB: rgb(215, 130, 126),
			On:  "surface",
		},
		Pine: Color{
			HSL: hsl(197, 53, 34),
			RGB: rgb(40, 105, 131),
			On:  "surface",
		},
		Foam: Color{
			HSL: hsl(189, 30, 48),
			RGB: rgb(86, 148, 159),
			On:  "surface",
		},
		Iris: Color{
			HSL: hsl(267, 22, 57),
			RGB: rgb(144, 122, 169),
			On:  "surface",
		},
		HighlightLow: Color{
			HSL: hsl(25, 35, 93),
			RGB: rgb(244, 237, 232),
		},
		HighlightMed: Color{
			HSL: hsl(10, 9, 86),
			RGB: rgb(223, 218, 217),
		},
		HighlightHigh: Color{
			HSL: hsl(315, 4, 80),
			RGB: rgb(206, 202, 205),
		},
	}
)

const description = "All natural pine, faux fur and a bit of soho vibes for the classy minimalist"

var (
	MainVariantMeta = VariantMeta{
		id:          "rose-pine",
		name:        "Rosé Pine",
		appearance:  "dark",
		description: description,
		colors:      &MainPalette,
	}

	MoonVariantMeta = VariantMeta{
		id:          "rose-pine-moon",
		name:        "Rosé Pine Moon",
		appearance:  "dark",
		description: description,
		colors:      &MoonPalette,
	}

	DawnVariantMeta = VariantMeta{
		id:          "rose-pine-dawn",
		name:        "Rosé Pine Dawn",
		appearance:  "light",
		description: description,
		colors:      &DawnPalette,
	}
)

var variants = []VariantMeta{
	MainVariantMeta,
	MoonVariantMeta,
	DawnVariantMeta,
}

func (p *Palette) Iter() iter.Seq2[string, *Color] {
	return func(yield func(string, *Color) bool) {
		v := reflect.ValueOf(p).Elem()
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			name := strings.ToLower(field.Name)
			color := v.Field(i).Addr().Interface().(*Color)

			if !yield(name, color) {
				return
			}
		}
	}
}

func (p *Palette) Get(role string) (*Color, bool) {
	v := reflect.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if strings.EqualFold(field.Name, role) {
			return v.Field(i).Addr().Interface().(*Color), true
		}
	}

	return nil, false
}
