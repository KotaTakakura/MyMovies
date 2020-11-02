package repository

import "MyPIPE/domain/model"

type TemporaryRegisterMailRepository interface {
	Send(mail *model.TemporaryRegisterMail) error
}
