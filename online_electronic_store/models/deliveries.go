package models

// import "gorm.io/gorm"

// type Delivery struct {
// 	gorm.Model
// 	DeliveryStatus string `gorm:"column:delivery_status;not null" json:"delivery_status"`
// 	OrderID        uint   `gorm:"column:order_id;not null" json:"order_id"`
// 	Order          Order  `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
// 	UserID         uint   `gorm:"column:user_id" json:"user_id"`
// 	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
// }

// func (Delivery) TableName() string {
// 	return "deliveries"
// }
