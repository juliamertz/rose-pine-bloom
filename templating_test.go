package main

import "testing"

func TestRenderTemplate(t *testing.T) {
	var values Values = Values{
		Id:          "rose-pine",
		Name:        "Rosé Pine",
		Key:         "main",
		Description: "All natural something something classy minimalist",
		Palette:     &MoonPalette,
		Variant: Variants{
			Main: &MainPalette,
			Moon: &MoonPalette,
			Dawn: &DawnPalette,
		},
	}

	content := `
   {{- define "color"}}
		 (
				 red: {{ div .RGB.R 255 | printf "%.8f" }},
				 green: {{ div .RGB.G 255 | printf "%.8f" }},
				 blue: {{ div .RGB.B 255 | printf "%.8f" }},
				 alpha: 1.0,
		 )
	 {{end -}}

	{{template "color" .Variant.Dawn.Love}}
`

	_, err := RenderTemplate("my.templ", content, &values)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	// t.Fatalf("whatt!?: %v", *out)
}


