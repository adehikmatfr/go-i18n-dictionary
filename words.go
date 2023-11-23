package dictionary

const (
	Hallo      Keyword = "hallo"
	Goodbye    Keyword = "goodbye"
	HalloNamef Keyword = "hallo_name"
)

var Words = map[Lang]map[Keyword]string{
	EN: {
		Hallo:      "Hello",
		Goodbye:    "Goodbye",
		HalloNamef: "Hallo {0}",
	},
	JP: {
		Hallo:      "こんにちは",
		Goodbye:    "さようなら",
		HalloNamef: "こんにちは {0}",
	},
}
