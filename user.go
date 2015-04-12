package main

import (
	"code.google.com/p/go-uuid/uuid"
	"crypto/sha1"
	"fmt"
	"gopkg.in/gorp.v1"
	"strings"
	"time"
)

const getUsersQuery = "SELECT id, name, password_hash, created_at FROM users ORDER BY name"
const getSingleUserQuery = "SELECT id, name, password_hash, created_at FROM users WHERE name = ?"

type User struct {
	Id           int64     `db:"id"`
	Name         string    `db:"name"`
	Password     string    `db:"-"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func NewUser(name string) *User {
	user := User{Id: 0, Name: name}
	user.Password = strings.Replace(uuid.New(), "-", "", -1)
	return &user
}

func GetUser(name string) (User, error) {
	var user User
	err := DbSelectOne(&user, getSingleUserQuery, name)
	return user, err
}

func GetUsers() ([]User, error) {
	var users []User
	_, err := DbSelectAll(&users, getUsersQuery)
	return users, err
}

func IsValidUser(name string, password string) bool {
	var user User
	if err := DbSelectOne(&user, getSingleUserQuery, name); err != nil {
		return false
	}

	return hashString(password) == user.PasswordHash
}

func (self *User) isNew() bool {
	return self.Id == 0
}

func hashString(plaintext string) string {
	hash := sha1.New()
	hash.Write([]byte(plaintext))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (self *User) Save() error {
	if self.isNew() {
		return DbInsert(self)
	}

	_, err := DbUpdate(self)
	return err
}

func (self *User) PreInsert(sqlExec gorp.SqlExecutor) error {
	self.CreatedAt = time.Now()
	self.PasswordHash = hashString(self.Password)
	return nil
}
