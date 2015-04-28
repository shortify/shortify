package shortify

import (
	"gopkg.in/gorp.v1"
	"time"
)

const tokenQuery = "SELECT id, token, url, created_at FROM redirects WHERE token = ?"
const urlQuery = "SELECT id, token, url, created_at FROM redirects WHERE url = ?"
const encodingSeed = int64(10000)

type Redirect struct {
	model
	Token string `db:"token" json:"token"`
	Url   string `db:"url" json:"url"`
}

func NewRedirect(url string) *Redirect {
	redir := new(Redirect)
	redir.Url = url

	return redir
}

func FindOrCreateRedirect(url string) (Redirect, error) {
	var redir Redirect
	err := db.selectOne(&redir, urlQuery, url)
	if err != nil {
		redir = *NewRedirect(url)
		err = redir.Save()
	}

	return redir, err
}

func FindRedirectByToken(token string) (Redirect, error) {
	var redir Redirect
	err := db.selectOne(&redir, tokenQuery, token)
	return redir, err
}

func (self *Redirect) Save() error {
	if self.isNew() {
		if err := db.insert(self); err != nil {
			return err
		}

		self.Token = ShortifyEncoder.Encode(self.Id + encodingSeed)
	}

	_, err := db.update(self)
	return err
}

func (self *Redirect) PreInsert(sqlExec gorp.SqlExecutor) error {
	self.CreatedAt = time.Now()
	return nil
}
