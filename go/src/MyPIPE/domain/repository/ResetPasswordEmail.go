package repository

import "MyPIPE/domain/model"

type ResetPasswordEmail interface {
	Send(email model.UserEmail,token model.UserPasswordRememberToken)error
}