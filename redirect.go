package main

import "time"

const tokenQuery = "SELECT id, token, url, created_at FROM redirects WHERE token = ?"
const urlQuery = "SELECT id, token, url, created_at FROM redirects WHERE url = ?"
const encodingSeed = int64(10000)

type Redirect struct {
	Id        int64     `db:"id"`
	Token     string    `db:"token"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}

func NewRedirect(url string) *Redirect {
	return &Redirect{0, "", url, time.Now()}
}

func (self *Redirect) isNew() bool {
	return self.Id == 0
}

func FindOrCreateRedirect(url string) (Redirect, error) {
	var redir Redirect
	err := DbSelectOne(&redir, urlQuery, url)
	if err != nil {
		redir = *NewRedirect(url)
		err = redir.Save()
	}

	return redir, err
}

func FindRedirectByToken(token string) (Redirect, error) {
	var redir Redirect
	err := DbSelectOne(&redir, tokenQuery, token)
	return redir, err
}

func (self *Redirect) Save() error {
	if self.isNew() {
		if err := DbInsert(self); err != nil {
			return err
		}

		self.Token = Base62Encode(self.Id + encodingSeed)
	}

	_, err := DbUpdate(self)
	return err
}
