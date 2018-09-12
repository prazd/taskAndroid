package mongo

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"
)

const CONN = "mongodb://mongo:27017"

func SetInfo(login, pass string) bool {
	session, err := mgo.Dial(CONN)
	if err != nil {
		log.Println(err)
	}
	c := session.DB("users").C("info")

	type User struct {
		Login string
	}
	var user User

	c.Find(bson.M{"login": login}).One(&user)

	if len(user.Login) != 0 {
		return false
	}

	err = c.Insert(bson.M{"login": login, "pass": hashAndSalt([]byte(pass))})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GetPassword(login string) string {
	session, err := mgo.Dial(CONN)
	if err != nil {
		log.Println(err)
	}
	c := session.DB("users").C("info")

	type User struct {
		Login    string
		Password string
	}
	var user User
	err = c.Find(bson.M{"login": login}).One(&user)
	if err != nil {
		log.Println(err)
	}
	if len(user.Password) == 0 {
		return "not found"
	} else {
		return user.Password
	}

}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
