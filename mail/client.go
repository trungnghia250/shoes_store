package mail

import (
	"errors"
	"gopkg.in/gomail.v2"
)

type Client struct {
	Credential *Credential
	Host       string
	Port       int
}

type Credential struct {
	User string
	Pass string
	cc   []string
	bcc  []string
	to   []string
}

func Send(msgOpt ...OptMessage) error {
	c := &Client{
		Credential: &Credential{
			User: "storeshop211@gmail.com",
			Pass: "storeshop",
			cc:   []string{"huynhtrungnghia250@gmail.com"},
		},
		Host: "smtp.gmail.com",
		Port: 465,
	}
	msg := NewMessage()
	opts := []OptMessage{WithFrom(c.Credential.User), WithTo(c.Credential.to), WithCc(c.Credential.cc), WithBcc(c.Credential.bcc)}
	opts = append(opts, msgOpt...)
	msg.applyMessage(opts...)
	if len(msg.mailMessage.GetHeader("From")) == 0 || len(msg.mailMessage.GetHeader("To")) == 0 {
		return errors.New("error : From or To not found")
	}
	d := gomail.NewPlainDialer(c.Host, c.Port, c.Credential.User, c.Credential.Pass)
	if msg.mailMessage == nil {
		return nil
	}
	err := d.DialAndSend(msg.mailMessage)
	return err
}
