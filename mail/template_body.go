package mail

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
)

type OptionBody func(template *Template)

func WithName(name string) OptionBody {
	return func(template *Template) {
		template.email.Body.Name = name
	}
}

func WithIntros(intros []string) OptionBody {
	return func(template *Template) {
		template.email.Body.Intros = intros
	}
}

func WithOuttros(outtros []string) OptionBody {
	return func(template *Template) {
		template.email.Body.Outros = outtros
	}
}

func WithTitle(title string) OptionBody {
	return func(template *Template) {
		template.email.Body.Title = title
	}
}

func WithGreeting(greeting string) OptionBody {
	return func(template *Template) {
		template.email.Body.Greeting = greeting
	}
}

func WithSignature(signature string) OptionBody {
	return func(template *Template) {
		template.email.Body.Signature = signature
	}
}

func WithDictionary(dictionary map[string]string) OptionBody {
	return func(template *Template) {
		for key, value := range dictionary {
			template.email.Body.Dictionary = append(template.email.Body.Dictionary, EmailEntry{
				{
					Key:   key,
					Value: value,
				},
			}...)
		}
	}
}

func WithFreeMarkDown(data string) OptionBody {
	return func(template *Template) {
		template.email.Body.FreeMarkdown = hermes.Markdown(data)
	}
}

func WithDictionaryLink(key, url, name string) OptionBody {
	return func(template *Template) {
		value := fmt.Sprintf("%s\" "+
			"rel=\"noopener\" style=\"text-decoration: underline; color: #0068A5;\" "+
			"target=\"_blank\">%s</a>"+" <br>", url, name)
		template.email.Body.Dictionary = append(template.email.Body.Dictionary, EmailEntry{
			{
				Key:   key,
				Value: value,
			},
		}...)

	}
}

func WithAction(actions []EmailButton) OptionBody {
	return func(template *Template) {
		for _, action := range actions {
			template.email.Body.Actions = append(template.email.Body.Actions, hermes.Action{
				Button: hermes.Button(action),
			})
		}
	}
}

func WithLogo(logo string) OptionBody {
	return func(template *Template) {
		template.template.Product.Logo = logo
	}
}

func WithLink(link string) OptionBody {
	return func(template *Template) {
		template.template.Product.Link = link
	}
}

func defaultSignature() string {
	return "Thanks"
}

func defaultGreeting() string {
	return "Hello our customer"
}

type emailTemplate hermes.Email

func NewOptionBody(bodyOptions ...OptionBody) []OptionBody {
	bodyOptions = append([]OptionBody{WithSignature(defaultSignature()), WithGreeting(defaultGreeting())},
		bodyOptions...,
	)
	return bodyOptions
}

func (t *Template) applyOption(opts ...OptionBody) {
	for _, opt := range opts {
		opt(t)
	}
}
