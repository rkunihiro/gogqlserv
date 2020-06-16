package entity

type User struct {
	ID   uint64 `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (p *User) TableName() string {
	return "User"
}
