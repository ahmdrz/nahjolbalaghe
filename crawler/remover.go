package main

import "strings"

func RemoveDiacritics(text string) string {
	characters := map[string]string{
		"ْ": "",
		"ٌ": "",
		"ٍ": "",
		"ً": "",
		"ُ": "",
		"ِ": "",
		"َ": "",
		"ّ": "",
		"":  "",
		"ؤ": "و",
		"ي": "ی",
		"ة": "ه",
		"إ": "ا",
		"أ": "ا",
		"ك": "ک",
	}

	for key, val := range characters {
		text = strings.Replace(text, key, val, -1)
	}

	return text
}
