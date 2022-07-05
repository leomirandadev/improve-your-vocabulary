package entities

import "time"

type Word struct {
	ID        uint64     `db:"id" json:"id"`
	Word      string     `db:"word" json:"word"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type WordRequest struct {
	Word string `db:"word" json:"word"`
}
