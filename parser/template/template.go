package template

var Funcs map[string]func(params []string) string

func nihongo(params []string) string {
	return "Teste"
}

func init() {
	Funcs = make(map[string]func(params []string) string)
	Funcs["nihongo"] = nihongo
}
