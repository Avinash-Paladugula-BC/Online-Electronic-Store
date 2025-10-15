package models

// import "gorm.io/gorm"

// type Order struct {
// 	gorm.Model
// 	ShippingAddress string `gorm:"column:shipping_address;not null" json:"shipping_address"`
// 	UserID          uint   `gorm:"column:user_id;not null" json:"user_id"`
// 	ProductID       uint   `gorm:"column:user_id;not null" json:"product_id"`
// 	User            User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
// 	// PaymentID       uint       `gorm:"column:payment_id;not null" json:"payment_id"`
// 	// Payment         Payment    `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment"`
// 	Deliveries []Delivery `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"deliveries"`
// }
