package shared

import (
	"os"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Attachment struct {
	File     []byte
	Filename string
	MimeType string
}

func SendEmail(subject, body string, to, cc []string, attachment *Attachment) error {
	smtpClient, err := ConnectToMailServer()
	if err != nil {
		return err
	}

	sender := os.Getenv("SENDER_EMAIL")

	email := mail.NewMSG()
	email.SetFrom(sender).
		AddTo(to...).
		AddCc(cc...).
		SetSubject(subject).
		SetBody(mail.TextPlain, body)
	// email.SetBodyData(mail.TextHTML, []byte(htmlBody))

	if attachment != nil {
		email.AddAttachmentData(attachment.File, attachment.Filename, attachment.MimeType)
		// email.Attach(&mail.File{FilePath: "../../files/template.xlsx", Name: "template.xlsx", Inline: true})
	}

	if email.Error != nil {
		return email.Error
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return nil
}

func ConnectToMailServer() (*mail.SMTPClient, error) {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	username := os.Getenv("EMAIL_USERNAME")
	pwd := os.Getenv("EMAIL_PASSWORD")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = portInt
	server.Username = username
	server.Password = pwd

	server.Encryption = mail.EncryptionNone
	server.Authentication = mail.AuthNone
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}

	return smtpClient, nil
}
