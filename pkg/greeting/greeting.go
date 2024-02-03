package greeting

type Language int8

const (
	English Language = iota
	Italian
	German
	Polish
)

func Greet(l Language) string {
	g := map[Language]string{
		English: "Hello!",
		Italian: "Ciao!",
		German:  "Hallo!",
		Polish:  "Cześć!",
	}
	return g[l]
}
