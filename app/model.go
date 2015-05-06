package app

import "time"

type model struct {
	Id        int64     `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func (self model) isNew() bool {
	return self.Id == 0
}
