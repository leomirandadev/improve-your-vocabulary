package entities

type Word struct {
	ID   uint64 `db:"id" json:"id"`
	Word string `db:"word" json:"word"`
}

type WordRequest struct {
	Word string `db:"word" json:"word"`
}
