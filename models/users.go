package models

import (
	"api/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	tmp := &User{}

	err := GetDB().Table("users").Where("login = ?", user.Login).First(tmp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error!"), false
	}
	if tmp.Login != "" {
		return utils.Message(false, "User with this login already exists"), false
	}
	return utils.Message(false, "Requirement passed!"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, err := user.Validate();
		!err {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return utils.Message(false, "Failed to create new user")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(GetTokenPassword()))
	user.Token = tokenString

	user.Password = ""
	response := utils.Message(true, "Account has been created!")
	response["user"] = user
	return response
}

func Login(login, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("users").Where("login = ?", login).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "There is no user with this Login!")
		}
		return utils.Message(false, "Connection error!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid Login/Password. Try again!")
	}

	user.Password = ""

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(GetTokenPassword()))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetUser(u uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Login == "" {
		return nil
	}

	user.Password = ""
	return user
}
