package entities

import "time"

type Word struct {
	ID        uint64     `db:"id" json:"id"`
	Word      string     `db:"word" json:"word"`
	Meanings  []Meaning  `json:"meanings"`
	UserID    uint64     `db:"user_id" json:"user_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type WordRequest struct {
	Word   string `db:"word" json:"word"`
	UserID uint64 `db:"user_id" json:"user_id"`
}
