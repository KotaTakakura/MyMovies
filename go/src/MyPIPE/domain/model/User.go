package model

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"time"
	"errors"
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

type GoodEvaluate struct{}

func (g GoodEvaluate)Evaluate(user *User,movieId MovieID)error{
	for _,goodMovieIds := range user.GoodMovies{
		if goodMovieIds == movieId{
			return errors.New("Duplicate Good Movie.")
		}
	}
	user.GoodMovies = append(user.GoodMovies,movieId)
	_ = g.UnsetBad(user,movieId)
	return nil
}

func (g GoodEvaluate)UnsetBad(user *User,movieId MovieID)error{
	var result []MovieID
	for _, badMovieIds := range user.BadMovies {
		if badMovieIds != movieId {
			result = append(result, badMovieIds)
		}
	}
	user.BadMovies = result
	return nil
}

type BadEvaluate struct{}

func (g BadEvaluate)Evaluate(user *User,movieId MovieID)error{
	for _,badMovieIds := range user.BadMovies{
		if badMovieIds == movieId{
			return errors.New("Duplicate Bad Movie.")
		}
	}

	user.BadMovies = append(user.BadMovies,movieId)
	_ = g.UnsetGood(user,movieId)
	return nil
}

func (g BadEvaluate)UnsetGood(user *User,movieId MovieID)error{
	var result []MovieID
	for _, goodMovieIds := range user.GoodMovies {
		if goodMovieIds != movieId {
			result = append(result, goodMovieIds)
		}
	}
	user.GoodMovies = result
	return nil
}

type UnsetGoodEvaluate struct{}

func (u UnsetGoodEvaluate)Evaluate(user *User,movieId MovieID)error{
	var result []MovieID
	for _, goodMovieIds := range user.GoodMovies {
		if goodMovieIds != movieId {
			result = append(result, goodMovieIds)
		}
	}
	user.GoodMovies = result
	return nil
}

type UnsetBadEvaluate struct{}

func (u UnsetBadEvaluate)Evaluate(user *User,movieId MovieID)error{
	var result []MovieID
	for _, badMovieIds := range user.BadMovies {
		if badMovieIds != movieId {
			result = append(result, badMovieIds)
		}
	}
	user.BadMovies = result
	return nil
}

func NewEvaluate(evaluation string)(IEvaluate,error){
	err := validation.Validate(evaluation,
		validation.Required,
		validation.In("good", "bad","unset_good","unset_bad"),
	)
	if err != nil {
		return nil, err
	}

	switch evaluation {
	case "good":
		return GoodEvaluate{},nil
	case "bad":
		return BadEvaluate{},nil
	case "unset_good":
		return UnsetGoodEvaluate{},nil
	case "unset_bad":
		return UnsetBadEvaluate{},nil
	}
	return nil,errors.New("Invalid Evaluation.")
}

type User struct {
	ID                UserID   `json:"id" gorm:"primaryKey"`
	Name              UserName `json:"name"`
	Password          UserPassword
	Email             UserEmail `json:"email"`
	Birthday          time.Time `json:"birthday"`
	Token             UserToken `json:"token"`
	Movies            []Movie
	Comments          []Comment
	GoodMovies        []MovieID   `gorm:"-"`
	BadMovies         []MovieID   `gorm:"-"`
	PlayLists         []PlayList
	Follows           []User	`gorm:"many2many:follow_users;"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewUser() *User {
	return &User{}
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

func (u *User) EmptyToken() {
	u.Token = ""
}

func (u User) TemporaryRegisteredWithinOneHour() bool {
	duration := time.Now().Sub(u.UpdatedAt)
	return int(duration.Minutes()) < 60
}

type IEvaluate interface {
	Evaluate(user *User,movieID MovieID) error
}

func (u *User)Evaluate(ievaluate IEvaluate,movieID MovieID)error{
	err := ievaluate.Evaluate(u,movieID)
	if err != nil{
		return err
	}
	return nil
}

//func (u *User)SetGoodMovie(movieId MovieID)error{
//	for _,goodMovieIds := range u.GoodMovies{
//		if goodMovieIds == movieId{
//			return errors.New("Duplicate Good Movie.")
//		}
//	}
//
//	u.GoodMovies = append(u.GoodMovies,movieId)
//	_ = u.UnsetBadMovie(movieId)
//	return nil
//}
//
//func (u *User)UnsetGoodMovie(movieId MovieID)error{
//	var result []MovieID
//	for _, goodMovieIds := range u.GoodMovies {
//		if goodMovieIds != movieId {
//			result = append(result, goodMovieIds)
//		}
//	}
//	u.GoodMovies = result
//	return nil
//}
//
//func (u *User)SetBadMovie(movieId MovieID)error{
//	for _,badMovieIds := range u.BadMovies{
//		if badMovieIds == movieId{
//			return errors.New("Duplicate Bad Movie.")
//		}
//	}
//
//	u.BadMovies = append(u.BadMovies,movieId)
//	_ = u.UnsetGoodMovie(movieId)
//	return nil
//}
//
//func (u *User)UnsetBadMovie(movieId MovieID)error{
//	var result []MovieID
//	for _, badMovieIds := range u.BadMovies {
//		if badMovieIds != movieId {
//			result = append(result, badMovieIds)
//		}
//	}
//	u.BadMovies = result
//	return nil
//}