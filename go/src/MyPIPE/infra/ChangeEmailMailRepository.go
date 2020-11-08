package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type ChangeEmailMail struct{}

func NewChangeEmailMail() *ChangeEmailMail {
	return &ChangeEmailMail{}
}

func (c ChangeEmailMail) Send(email model.UserEmail, token model.UserEmailChangeToken) error {
	sess := session.Must(session.NewSession())
	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(string(email)),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data: aws.String(
						"メールアドレスの変更をを完了するために、以下のURLにアクセスしてください。\nURLの有効期限は30分間です。\n" +
							"このメールに心当たりのない場合は、URLにアクセスせず破棄してください。\nhttps:" +
							"//www.frommymovies.com/email?token=" + string(token)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("MyMoviesメールアドレス変更のご案内"),
			},
		},
		Source: aws.String("info@mail.frommymovies.com"),
	}

	_, sendMailErr := svc.SendEmail(input)
	if sendMailErr != nil {
		return sendMailErr
	}
	return nil
}
