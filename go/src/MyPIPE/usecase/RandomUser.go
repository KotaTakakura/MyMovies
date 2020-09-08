package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"math/rand"
	"time"
	"fmt"
)

type RandomUser struct{
	userRepository	repository.UserRepository
}

func NewRandomUser(u repository.UserRepository) *RandomUser{
	return &RandomUser{
		userRepository: u,
	}
}

func (r RandomUser) getRandomId() int{
	rand.Seed(time.Now().UnixNano())
	var randomId int
	randomId = rand.Intn(100)
	randomId = 1
	return randomId
}

func (r RandomUser) GetRandomIdUser() *model.User{
	var randomId int
	randomId = r.getRandomId()
	var user *model.User
	fmt.Println("||||||||||||||")
	user = r.userRepository.FindById(randomId)
	fmt.Println("||||||||||||||")
	var userAge int
	var err error
	userAge, err = user.CalcAge()
	if err != nil{
		fmt.Println("NOOOOO")
	}
	fmt.Println(userAge)
	return user
}
