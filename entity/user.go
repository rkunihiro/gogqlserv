package entity

// UserID user identify
type UserID uint64

// User user entity
type User struct {
	ID   UserID `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

// TableName binding database table name
func (p *User) TableName() string {
	return "User"
}
