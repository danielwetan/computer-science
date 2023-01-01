package hello

const englishHelloPrefix = "Hello, "
const indonesianHelloPrefix = "Halo, "
const turkishHelloPrefix = "Merhaba, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "turkish":
		prefix = turkishHelloPrefix
	case "indonesian":
		prefix = indonesianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
