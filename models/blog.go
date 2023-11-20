package models

import "time"

type Blog struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	File      string    `json:"file"`
	Title     string    `json:"title" gorm:"unique;not null"`
	Author    string    `json:"author" default:"admin"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:nano"`
	UpdateAt  time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
}
