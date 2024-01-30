package greeting

type Language int8

const (
	English Language = iota
	Italian
	German
)

func Greet(l Language) string {
	g := map[Language]string{
		English: "Hello!",
		Italian: "Ciao!",
		German:  "Hallo!",
	}
	return g[l]
}
