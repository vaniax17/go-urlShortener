package postgres

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"` //Hashed
}
