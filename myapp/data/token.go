package data

import (
	"time"
)

type Token struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	Email     string    `db:"email" json:"email"`
	PlainText string    `db:"-" json:"token"`
	Hash      []byte    `ds:"token_hash" json:"-"`
	CreatedAt time.Time `ds:"created_at" json:"created_at"`
	UpdatedAt time.Time `ds:"updated_at" json:"updated_at"`
	Expires   time.Time `ds:"expiry" json:"expiry"`
}

func (t *Token) Table() string {

	return "tokens"

}
