package hello

const prefixHello = "Hello, "
const prefixHelloFrench = "Bonjour, "
const prefixHelloPortuguese = "Olá, "
const prefixHelloSpanish = "Hola, "
const period = "."

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name + period
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = prefixHelloSpanish
	case "portuguese":
		prefix = prefixHelloPortuguese
	case "french":
		prefix = prefixHelloFrench
	default:
		prefix = prefixHello
	}
	return
}
