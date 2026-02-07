package hello

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		return "Hello, world!"
	}

	prefixes := map[string]string{
		"english": "Hello",
		"spanish": "Hola",
		"italian": "Ciao",
	}

	greetingPrefix, ok := prefixes[language]
	var greeting string

	if ok {
		greeting = fmt.Sprintf("%s, %s!", greetingPrefix, name)
	} else {
		greeting = fmt.Sprintf("%s, %s!", prefixes["english"], name)
	}

	return greeting
}
