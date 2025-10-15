package models

// import "gorm.io/gorm"

// type User struct {
// 	gorm.Model
// 	Name        string   `gorm:"column:name;not null"  json:"name"`
// 	Mail        string   `gorm:"column:mail;unique;not null" json:"mail"`
// 	Password    string   `gorm:"column:password;not null" json:"password"`
// 	Role        string   `gorm:"column:role; default:customer" json:"role"`
// 	PhoneNumber string   `gorm:"column:phone_number;not null" json:"phone_number"`
// 	Address     string   `gorm:"column:address;" json:"address"`
// 	Orders      []Order  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
// 	Reviews     []Review `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
// }
