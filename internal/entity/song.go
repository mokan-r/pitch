package entity

type Song struct {
	Id       string `json:"id,omitempty" db:"uid"`
	Title    string `json:"title,omitempty" db:"title"`
	Duration int64  `json:"duration,omitempty" db:"duration"`
}
