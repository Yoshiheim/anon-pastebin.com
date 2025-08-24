package models

type Paste struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}
