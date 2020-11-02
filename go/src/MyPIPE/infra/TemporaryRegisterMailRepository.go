package infra

import (
	"MyPIPE/domain/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type TemporaryRegisterMailRepository struct {}

func NewTemporaryRegisterMailRepository()*TemporaryRegisterMailRepository{
	return &TemporaryRegisterMailRepository{}
}

func (t TemporaryRegisterMailRepository)Send(mail *model.TemporaryRegisterMail)error{
	sess := session.Must(session.NewSession())
	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				//aws.String(string(mail.To)),
				aws.String("complaint@simulator.amazonses.com"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(
						"現在、仮登録の状態です。以下のURLにアクセスして情報を入力し、本登録を完了してください\nhttp:" +
						"//drp25i52zjwj0.cloudfront.net/register?token=" + string(mail.Token)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("MyMovies仮登録のご案内"),
			},
		},
		Source:        aws.String(string(mail.From)),
	}

	_,sendMailErr := svc.SendEmail(input)
	if sendMailErr != nil{
		return sendMailErr
	}
	return nil
}