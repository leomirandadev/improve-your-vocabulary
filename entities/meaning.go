package entities

type Meaning struct {
	ID      uint64 `db:"id" json:"id"`
	WordID  uint64 `db:"word_id" json:"word_id"`
	Meaning string `db:"meaning" json:"meaning"`
}

type MeaningRequest struct {
	WordID  uint64 `db:"word_id" json:"word_id"`
	Meaning string `db:"meaning" json:"meaning"`
}
