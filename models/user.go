package models

type User struct {
}

func (m *User) TableName() string {
	return GetUserTableName()
}
