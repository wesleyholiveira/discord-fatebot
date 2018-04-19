package parser

import (
	"errors"
	"fmt"
	"github.com/wesleyholiveira/discord-fatebot/model"
	t "github.com/wesleyholiveira/discord-fatebot/parser/template"
	s "github.com/wesleyholiveira/discord-fatebot/sanitizer"
	"regexp"
	"strings"
)

type Parser struct {
	HPage model.FG
}

var name string

func New(m model.FG, n string) *Parser {
	name = n
	return &Parser{HPage: m}
}

func Change(n string) {
	name = n
}

func (p *Parser) Parse() error {
	p.HPage[name].Servant.Name = p.HPage[name].Title
	p.redirect()

	err := p.template()
	if err != nil {
		return err
	}

	desc, err := p.getDescription()
	if err != nil {
		return err
	}

	p.HPage[name].Servant.Description = s.RemoveAll(desc)
	return nil
}

func (p *Parser) redirect() {
	page := p.HPage[name]
	text := page.Revision.Text

	re := regexp.MustCompile(s.RedirectRegexp)
	matches := re.FindStringSubmatch(text)

	if len(matches) > 1 {
		redirectTo := s.RemoveDoubleBrackets(matches[1])
		p.HPage[redirectTo].Servant.Name = s.Title(redirectTo)
		p.HPage[name] = p.HPage[redirectTo]
		Change(redirectTo)
	}
}

func (p *Parser) template() error {
	re := regexp.MustCompile(s.RedirectRegexp)
	text := p.HPage[name].Revision.Text
	matches := re.FindAllString(text, len(text))

	if len(matches) < 1 {
		return errors.New("Nenhum template foi encontrado")
	}

	for _, match := range matches {
		match = s.RemoveDoubleCurlyBraces(match)
		params := strings.Split(match, "|")

		if templateFunc, ok := t.Funcs[strings.ToLower(params[0])]; ok {
			p.HPage[name].Revision.Text = templateFunc(params[1:])
		}
	}
	return nil
}

func (p *Parser) getDescription() (string, error) {
	page := p.HPage[name]
	text := page.Revision.Text

	re := regexp.MustCompile(`{{[N|n]ihongo\|`)
	find := re.FindString(text)

	starts := strings.Index(text, fmt.Sprintf("%s%s", find, page.Servant.Name))
	ends := strings.Index(text, "==")
	if starts < 0 {
		return "", errors.New("Não foi possível obter a descrição")
	}
	if ends < 0 {
		ends = len(text)
	}
	return text[starts:ends], nil
}
