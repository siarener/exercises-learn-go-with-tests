package hello

import "fmt"

const (
	spanish     = "ES"
	english     = "EN"
	french      = "FR"
	hola        = "Hola"
	hello       = "Hello"
	bonjour     = "Bonjour"
	defaultName = "World"
)

var greet = map[string]string{
	spanish: hola,
	english: hello,
	french:  bonjour,
}

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	prefix, ok := greet[language]

	if ok == false {
		return greet[english]
	}
	return prefix
}

// func greetingPrefix(language string) (prefix string) {
// 	switch language {
// 	case french:
// 		prefix = bonjour
// 	case spanish:
// 		prefix = hola
// 	default:
// 		prefix = hello
// 	}
// 	return prefix
// }

func main() {
	fmt.Println(Hello("<Placeholder>", ""))
}
