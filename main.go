package main

import (
	"encoding/xml"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wesleyholiveira/discord-fatebot/model"
	"github.com/wesleyholiveira/discord-fatebot/parser"
	"io/ioutil"
	"os"
)

func ConvertToMap(x *model.XMLResult) model.FG {
	vMap := make(model.FG)
	for i, page := range x.Page {
		vMap[page.Title] = &x.Page[i]
	}
	return vMap
}

func main() {
	var name string
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Nome da página é obrigatório")
	}

	name = args[1]
	v := new(model.XMLResult)
	file, err := ioutil.ReadFile("dump-fate/typemoon_dump.xml")
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(file, v)
	if err != nil {
		log.Fatal(err)
	}

	a := ConvertToMap(v)
	if _, ok := a[name]; !ok {
		log.Fatal("Página não encontrada")
	}

	p := parser.New(a, name)
	err = p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", a[name].Servant)
}
