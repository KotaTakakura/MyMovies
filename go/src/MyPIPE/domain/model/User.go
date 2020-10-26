package model

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

type UserID uint64

func NewUserID(userId uint64) (UserID, error) {
	err := validation.Validate(userId,
		validation.Required,
	)
	if err != nil {
		return UserID(0), err
	}
	return UserID(userId), nil
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

type UserProfileImage struct{
	Name string
	FileHeader multipart.FileHeader
	File multipart.File
}

func NewUserProfileImage(fileHeader multipart.FileHeader,file multipart.File)(*UserProfileImage,error){
	extension := filepath.Ext(fileHeader.Filename)
	if !(extension == ".jpg" || extension == ".JPG" || extension == ".png" || extension == ".PNG" || extension == ".bmp" || extension == ".BMP" || extension == ".gif" || extension == ".GIF"){
		return nil,errors.New("Image File Only.")
	}
	if fileHeader.Size > 2000000 {
		return nil,errors.New("Too Large File.")
	}
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	return &UserProfileImage{
		Name:       timestamp + extension,
		FileHeader: fileHeader,
		File:       file,
	},nil
}

type User struct {
	ID               UserID   `json:"id" gorm:"primaryKey"`
	Name             UserName `json:"name"`
	Password         UserPassword
	Email            UserEmail `json:"email"`
	Birthday         time.Time `json:"birthday"`
	ProfileImageName string    `json:"profile_image_name"`
	Token            UserToken `json:"token"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func NewUser(email UserEmail, birthday time.Time) *User {
	return &User{
		Email:    email,
		Birthday: birthday,
	}
}

func (u *User) SetBirthday(stringBirthday string) error {
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

func (u *User) ChangeName(name UserName) error {
	u.Name = name
	return nil
}

func (u *User) ChangeEmail(email UserEmail) error {
	u.Email = email
	return nil
}

func (u *User) ChangePassword(password UserPassword) error {
	u.Password = password
	return nil
}

func (u *User) EmptyToken() {
	u.Token = ""
}

func (u User) TemporaryRegisteredWithinOneHour() bool {
	duration := time.Now().Sub(u.UpdatedAt)
	return int(duration.Minutes()) < 60
}

type IEvaluate interface {
	Evaluate(user *User, movieID MovieID) error
}

func (u *User) Evaluate(ievaluate IEvaluate, movieID MovieID) error {
	err := ievaluate.Evaluate(u, movieID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) SetProfileImage(profileImage UserProfileImage) error {
	u.ProfileImageName = profileImage.Name
	return nil
}

func (u *User) SetNewToken() error {
	u.Token = UserToken(uuid.New().String())
	return nil
}

func (u *User) Register(name UserName, password UserPassword, birthday time.Time) error {
	duration := time.Now().Sub(u.UpdatedAt)
	if int(duration.Minutes()) > 60 {
		return errors.New("Invalid Token.")
	}
	u.Token = ""
	u.Name = name
	u.Password = password
	u.Birthday = birthday
	return nil
}
