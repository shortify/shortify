package shortify

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
	model
	Name         string `db:"name"`
	Password     string `db:"-"`
	PasswordHash string `db:"password_hash"`
}

func NewUser(name string) *User {
	user := new(User)
	user.Name = name
	user.Password = newPassword()

	return user
}

func GetUser(name string) (User, error) {
	var user User
	err := shortifyDb.selectOne(&user, getSingleUserQuery, name)
	return user, err
}

func GetUsers() ([]User, error) {
	var users []User
	_, err := shortifyDb.selectAll(&users, getUsersQuery)
	return users, err
}

func IsValidUser(name string, password string) bool {
	var user User
	if err := shortifyDb.selectOne(&user, getSingleUserQuery, name); err != nil {
		return false
	}

	return hashString(password) == user.PasswordHash
}

func hashString(plaintext string) string {
	hash := sha1.New()
	hash.Write([]byte(plaintext))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func newPassword() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}

func (self *User) ResetPassword() error {
	self.Password = newPassword()
	self.PasswordHash = hashString(self.Password)
	return self.Save()
}

func (self *User) Save() error {
	if self.isNew() {
		return shortifyDb.insert(self)
	}

	_, err := shortifyDb.update(self)
	return err
}

func (self *User) PreInsert(sqlExec gorp.SqlExecutor) error {
	self.CreatedAt = time.Now()
	self.PasswordHash = hashString(self.Password)
	return nil
}
