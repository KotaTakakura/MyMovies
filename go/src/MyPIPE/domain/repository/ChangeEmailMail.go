package repository

import "MyPIPE/domain/model"

type ChangeEmailMail interface {
	Send(email model.UserEmail, token model.UserEmailChangeToken) error
}
