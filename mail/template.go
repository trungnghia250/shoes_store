package mail

import (
	"github.com/matcornic/hermes/v2"
)

type EmailEntry []hermes.Entry
type EmailMarkdown hermes.Markdown
type EmailButton hermes.Button
type EmailColumns hermes.Columns
type EmailEntryTable [][]hermes.Entry

type EmailTable struct {
	Data    EmailEntryTable // Contains data
	Columns EmailColumns
}

type EmailAction struct {
	Instructions string
	Button       EmailButton
}

type Template struct {
	template *hermes.Hermes
	email    emailTemplate
}

func newTemplate() *Template {
	template := &Template{
		template: &hermes.Hermes{
			Product: hermes.Product{
				Name:      "Amazing shoes",
				Copyright: "Copyright © 2021 Amazing shoes",
			},
		},
		email: emailTemplate{
			Body: hermes.Body{},
		},
	}
	return template
}

func (t *Template) generateHTML(optionBody ...OptionBody) (string, error) {
	// Generate the HTML template
	t.applyOption(NewOptionBody(optionBody...)...)
	res, err := t.template.GenerateHTML(hermes.Email(t.email))
	return res, err
}

func (t *Template) generatePlainText(optionBody ...OptionBody) (string, error) {
	// Generate the plaintext template
	t.applyOption(NewOptionBody(optionBody...)...)
	res, err := t.template.GeneratePlainText(hermes.Email(t.email))
	return res, err
}
