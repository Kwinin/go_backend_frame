package utils

import "time"

type BaseModel struct {
	Id        int       `gorm:"primary_key;comment:'id 自增'" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime;NOT NULL;comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;NOT NULL;comment:'更新时间'" json:"updated_at"`
}
