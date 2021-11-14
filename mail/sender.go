package mail

import (
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

type OptMessage func(*Message)

func WithMessageHTML(subject string, optsBody ...OptionBody) OptMessage {
	return func(message *Message) {
		template := newTemplate()
		body, _ := template.generateHTML(optsBody...)
		message.body = body
		message.mailMessage.SetBody("text/html", body)
		message.mailMessage.SetHeader("Subject", subject)
	}
}

func WithImage(imageUrl string) OptMessage {
	return func(message *Message) {
		matches := "</dl>"
		image := fmt.Sprintf(`<img style="-webkit-user-select: none;margin: auto;cursor: zoom-in;" src="%s">`, imageUrl)
		index := strings.LastIndex(message.body, matches)
		message.body = message.body[:index+len(matches)+1] + image + message.body[index+len(matches)+1:]
		message.mailMessage.SetBody("text/html", message.body)
	}
}

func WithMessagePlainText(subject string, optsBody ...OptionBody) OptMessage {
	return func(message *Message) {
		template := newTemplate()
		body, _ := template.generatePlainText(optsBody...)
		message.mailMessage.SetBody("text/plain", body)
		message.mailMessage.SetHeader("Subject", subject)
	}
}

func WithAttachment(filename string) OptMessage {
	return func(message *Message) {
		message.mailMessage.Attach(filename)
	}
}

func WithAddAlternative(filename string) OptMessage {
	return func(message *Message) {
		message.mailMessage.AddAlternative("text/html", filename)
	}
}

func WithFrom(from string) OptMessage {
	return func(message *Message) {
		message.mailMessage.SetHeader("From", from)
	}
}

func WithTo(to []string) OptMessage {
	return func(message *Message) {
		if len(to) > 0 {
			message.mailMessage.SetHeader("To", to...)
		}
	}
}

func WithBcc(bcc []string) OptMessage {
	return func(message *Message) {
		if len(bcc) > 0 {
			message.mailMessage.SetHeader("bcc", bcc...)
		}
	}
}

func WithCc(cc []string) OptMessage {
	return func(message *Message) {
		if len(cc) > 0 {
			message.mailMessage.SetHeader("cc", cc...)
		}
	}
}

func defaultBody() string {
	template := newTemplate()
	body, _ := template.generatePlainText()
	return body
}

func defaultContentType() string {
	return "text/plain"
}

func NewMessage() *Message {
	return &Message{
		mailMessage: gomail.NewMessage(),
	}
}

type Message struct {
	// MIME type. Eg: text/plain, text/html...
	mailMessage *gomail.Message
	body        string
}

func (m *Message) applyMessage(opts ...OptMessage) {
	//m.mailMessage.SetBody(defaultContentType(), defaultBody())
	for _, opt := range opts {
		opt(m)
	}
}
