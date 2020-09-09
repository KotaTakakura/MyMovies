package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type User struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Password  string
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	Token     string    `json:"token"`
	Movies    []Movie
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser() *User {
	return &User{}
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

func (u *User) SetHashedPassword(pass string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 12)
	u.Password = fmt.Sprintf("%s", hash)
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
