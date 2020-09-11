package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"time"
)

type UserID uint64

func NewUserID(userId uint64) UserID {
	return UserID(userId)
}

type UserName string

func NewUserName(userName string) (UserName, error) {
	err := validation.Validate(userName,
		validation.Required,
		validation.RuneLength(1, 200),
	)
	if err != nil {
		return UserName(""), err
	}
	return UserName(userName), nil
}

type UserPassword string

func NewUserPassword(userPassword string) (UserPassword, error) {
	err := validation.Validate(userPassword,
		validation.Required,
		validation.Match(regexp.MustCompile("^[0-9a-zA-Z]*$")),
		validation.RuneLength(8, 20),
	)
	if err != nil {
		return UserPassword(""), err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(userPassword), 12)
	return UserPassword(fmt.Sprintf("%s", hash)), nil
}

type UserEmail string

func NewUserEmail(userEmail string) (UserEmail, error) {
	err := validation.Validate(userEmail,
		validation.Required,
		is.Email,
	)
	if err != nil {
		return UserEmail(""), err
	}
	return UserEmail(userEmail), nil
}

type UserToken string

func NewUserToken(userToken string) (UserToken, error) {
	err := validation.Validate(userToken,
		validation.Required,
	)
	if err != nil {
		return UserToken(""), err
	}
	return UserToken(userToken), nil
}

type User struct {
	ID        UserID   `json:"id" gorm:"primaryKey"`
	Name      UserName `json:"name"`
	Password  UserPassword
	Email     UserEmail `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Token     UserToken `json:"token"`
	Movies    []Movie
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetBirthday(stringBirthday string) error{
	err := validation.Validate(stringBirthday,
		validation.Required,
		validation.Date("2006-01-02"),
	)
	if err != nil {
		return err
	}
	birthday, _ := time.Parse("2006-01-02", stringBirthday)
	u.Birthday = birthday
	return nil
}

func (u *User) CalcAge() (int, error) {
	// 現在日時を数値のみでフォーマット (YYYYMMDD)
	dateFormatOnlyNumber := "20060102" // YYYYMMDD

	now := time.Now().Format(dateFormatOnlyNumber)
	birthday := u.Birthday.Format(dateFormatOnlyNumber)

	// 日付文字列をそのまま数値化
	nowInt, err := strconv.Atoi(now)
	if err != nil {
		return 0, err
	}
	birthdayInt, err := strconv.Atoi(birthday)
	if err != nil {
		return 0, err
	}

	// (今日の日付 - 誕生日) / 10000 = 年齢
	age := (nowInt - birthdayInt) / 10000
	return age, nil
}

func (u User) CheckPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(pass))
	if err != nil {
		return false
	}
	return true
}

func (u *User) EmptyToken() {
	u.Token = ""
}

func (u User) TemporaryRegisteredWithinOneHour() bool{
	duration := time.Now().Sub(u.UpdatedAt)
	return int(duration.Minutes()) < 60
}
