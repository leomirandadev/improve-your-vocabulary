package entities

import "time"

type Meaning struct {
	ID        uint64     `db:"id" json:"id"`
	WordID    uint64     `db:"word_id" json:"word_id"`
	Meaning   string     `db:"meaning" json:"meaning"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type MeaningRequest struct {
	WordID  uint64 `db:"word_id" json:"word_id"`
	Meaning string `db:"meaning" json:"meaning"`
}
