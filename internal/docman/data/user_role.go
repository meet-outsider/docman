package data

type UserRole struct {
	UserID uint `gorm:"primary_key"`
	RoleID uint `gorm:"primary_key"`
}
