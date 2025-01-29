package entity

type Role struct {
	RoleID   int    `json:"roleID" gorm:"type:smallint;primary_key;autoIncrement"`
	RoleName int    `json:"roleName" gorm:"type:varchar(30);not null"`
	Users    []User `json:"users" gorm:"foreignKey:RoleID"`
}
