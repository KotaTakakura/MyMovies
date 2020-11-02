package model

type TemporaryRegisterMailFrom string

func NewTemporaryRegisterMailFrom(from string)(TemporaryRegisterMailFrom,error){
	return TemporaryRegisterMailFrom(from),nil
}

type TemporaryRegisterMailTo string

func NewTemporaryRegisterMailTo(to string)(TemporaryRegisterMailTo,error){
	return TemporaryRegisterMailTo(to),nil
}

type TemporaryRegisterMailToken string

func NewTemporaryRegisterMailToken(token string)(TemporaryRegisterMailToken,error){
	return TemporaryRegisterMailToken(token),nil
}

type TemporaryRegisterMail struct{
	From	TemporaryRegisterMailFrom
	To	UserEmail
	Token	UserToken
}

func NewTemporaryRegisterMail(to UserEmail,token UserToken)*TemporaryRegisterMail{
	return &TemporaryRegisterMail{
		From:    TemporaryRegisterMailFrom("info@frommymovies.com"),
		To:      to,
		Token:	token,
	}
}
