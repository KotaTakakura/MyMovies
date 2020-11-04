package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type ResetPasswordEmail struct{}

func NewResetPasswordEmail()*ResetPasswordEmail{
	return &ResetPasswordEmail{}
}

func (r ResetPasswordEmail)Send(email model.UserEmail,token model.UserPasswordRememberToken)error{
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
						"パスワードを再発行します。以下のURLにアクセスして新しいパスワードを設定してください。\nhttps:" +
							"//www.frommymovies.com/reset?token=" + string(token)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("MyMoviesパスワード再設定のご案内"),
			},
		},
		Source: aws.String("reset@mail.frommymovies.com"),
	}

	_, sendMailErr := svc.SendEmail(input)
	if sendMailErr != nil {
		return sendMailErr
	}
	return nil
}