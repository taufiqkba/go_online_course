package sendgrid

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	register "go_online_course/internal/register/dto"
	"html/template"
	"os"
	"path/filepath"
)

type Mail interface {
	SendVerificationEmail(toEmail string, dto register.CrateEmailVerification)
}

type MailImpl struct {
}

func (mi *MailImpl) sendMail(toEmail string, result string, subject string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	to := mail.NewEmail(toEmail, toEmail)

	messages := mail.NewSingleEmail(from, subject, to, "", result)
	client := sendgrid.NewSendClient(os.Getenv("MAIL_KEY"))

	resp, err := client.Send(messages)
	if err != nil {
		fmt.Println(err)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Printf("Email has been success sent to %s\n", toEmail)
	}
}

func (mi *MailImpl) SendVerificationEmail(toEmail string, dto register.CrateEmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/email/verification_email.gohtml")

	result, err := ParseTemplate(templateFile, dto)
	if err != nil {
		fmt.Print(err)
	}

	mi.sendMail(toEmail, result, dto.Subject)
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewMail() Mail {
	return &MailImpl{}
}
