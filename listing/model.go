package listing

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type Item struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Category    string
	Quantity    int
	Status      bool
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type Category struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type History struct {
	ID        int `gorm:"primaryKey"`
	ItemID    int
	Action    bool
	Quantity  int
 	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}