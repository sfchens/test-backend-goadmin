package smtp_email

import (
	"bytes"
	"crypto/tls"
	"csf/library/easy_config"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"io"
	"mime"
	"net/smtp"
	"net/textproto"
	"path/filepath"
)

const (
	DisabledText = "email is disabled"
)

type Dialer struct {
	// Host represents the host of the SMTP server.
	Host string
	// Port represents the port of the SMTP server.
	Port int
	// Username is the username to use to authenticate to the SMTP server.
	Username string
	// Password is the password to use to authenticate to the SMTP server.
	Password string
}

func NewDialer(host string, port int, username, password string) *Dialer {
	return &Dialer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (d *Dialer) Addr() string {
	return fmt.Sprintf("%s:%d", d.Host, d.Port)
}

func (d *Dialer) PlainAuth() smtp.Auth {
	return smtp.PlainAuth("", d.Username, d.Password, d.Host)
}

type SmtpEmail struct {
	Email *email.Email
}

func NewEmail() *SmtpEmail {
	e := email.NewEmail()
	return &SmtpEmail{
		Email: e,
	}
}

func (e *SmtpEmail) SetFrom(email string) *SmtpEmail {
	e.Email.From = email
	return e
}

func (e *SmtpEmail) AddTo(emails ...string) *SmtpEmail {
	for _, mail := range emails {
		e.Email.To = append(e.Email.To, mail)
	}

	return e
}

func (e *SmtpEmail) AddCc(emails ...string) *SmtpEmail {
	for _, mail := range emails {
		e.Email.Cc = append(e.Email.Cc, mail)
	}
	return e
}

func (e *SmtpEmail) AddBcc(emails ...string) *SmtpEmail {
	for _, mail := range emails {
		e.Email.Bcc = append(e.Email.Bcc, mail)
	}
	return e
}

func (e *SmtpEmail) AddReplyTo(emails ...string) *SmtpEmail {
	for _, mail := range emails {
		e.Email.ReplyTo = append(e.Email.ReplyTo, mail)
	}
	return e
}

func (e *SmtpEmail) SetSubject(subject string) *SmtpEmail {
	e.Email.Subject = subject
	return e
}

func (e *SmtpEmail) SetHtml(message []byte) *SmtpEmail {
	e.Email.HTML = message
	return e
}

func (e *SmtpEmail) SetText(message []byte) *SmtpEmail {
	e.Email.Text = message
	return e
}

func (e *SmtpEmail) SetSender(email string) *SmtpEmail {
	e.Email.Sender = email
	return e
}

func (e *SmtpEmail) Attach(r io.Reader, filename string, c string) (err error) {
	_, err = e.Email.Attach(r, filename, c)
	return err
}

func (e *SmtpEmail) AttachSteam(r io.Reader, filename string) (err error) {
	var buffer bytes.Buffer
	if _, err = io.Copy(&buffer, r); err != nil {
		return
	}

	ct := mime.TypeByExtension(filepath.Ext(filename))
	basename := filepath.Base(filename)

	at := &email.Attachment{
		Filename:    basename,
		ContentType: ct,
		Header:      textproto.MIMEHeader{},
		Content:     buffer.Bytes(),
	}
	e.Email.Attachments = append(e.Email.Attachments, at)

	return nil
}

// AttachFile is used to attach content to the email.
func (e *SmtpEmail) AttachFile(filename string) (err error) {
	_, err = e.Email.AttachFile(filename)
	return err
}

// Send an email
func (e *SmtpEmail) Send() error {
	if IsClosed() {
		return errors.New(DisabledText)
	}
	dialer := NewDialer(
		easy_config.Config.Email.Host,
		easy_config.Config.Email.Port,
		easy_config.Config.Email.Username,
		easy_config.Config.Email.Password,
	)

	addr := dialer.Addr()
	a := dialer.PlainAuth()

	return e.Email.Send(addr, a)
}

// SendWithTLS sends an email over tls with an optional TLS config.
func (e *SmtpEmail) SendWithTLS() error {
	if IsClosed() {
		return errors.New(DisabledText)
	}

	dialer := NewDialer(
		easy_config.Config.Email.Host,
		easy_config.Config.Email.Port,
		easy_config.Config.Email.Username,
		easy_config.Config.Email.Password,
	)

	addr := dialer.Addr()
	a := dialer.PlainAuth()

	return e.Email.SendWithTLS(addr, a, &tls.Config{ServerName: dialer.Host})
}

func IsClosed() bool {
	return easy_config.Config.Email.Disabled
}
