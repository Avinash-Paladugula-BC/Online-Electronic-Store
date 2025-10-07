package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string   `gorm:"column:name; uniqueIndex;not null" json:"name"`
	Category string   `gorm:"column:category; not null" json:"category"`
	Price    float64  `gorm:"column:price; not null" json:"price"`
	Brand    string   `gorm:"column:brand"  json:"brand"`
	Quantity uint     `gorm:"column:quantity; not null" json:"quantity"`
	Reviews  []Review `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"  json:"reviews"`
}

type User struct {
	gorm.Model
	Name        string   `gorm:"column:name;not null"  json:"name"`
	Mail        string   `gorm:"column:mail;unique;not null" json:"mail"`
	Password    string   `gorm:"column:password;not null" json:"password"`
	Role        string   `gorm:"column:role; default:customer" json:"role"`
	PhoneNumber string   `gorm:"column:phone_number;not null" json:"phone_number"`
	Address     string   `gorm:"column:address;" json:"address"`
	Orders      []Order  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	Reviews     []Review `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
}

type Order struct {
	gorm.Model
	ShippingAddress string     `gorm:"column:shipping_address;not null" json:"shipping_address"`
	UserID          uint       `gorm:"column:user_id;not null" json:"user_id"`
	User            User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	PaymentID       uint       `gorm:"column:payment_id;not null" json:"payment_id"`
	Payment         Payment    `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment"`
	Deliveries      []Delivery `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"deliveries"`
}

type Payment struct {
	gorm.Model
	OrderID       uint   `gorm:"column:order_id;not null" json:"order_id"`
	UserID        uint   `gorm:"column:user_id;not null" json:"user_id"`
	PaymentStatus bool   `gorm:"column:payment_status;not null" json:"payment_status"`
	PaymentMethod string `gorm:"column:payment_method;not null" json:"payment_method"`
	TransactionID uint   `gorm:"column:transaction_id;not null" json:"transaction_id"`
}

type Review struct {
	gorm.Model
	Rating    uint    `gorm:"column:rating;not null" json:"rating"`
	Review    string  `gorm:"column:review" json:"review"`
	UserID    uint    `gorm:"column:user_id;not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ProductID uint    `gorm:"column:product_id" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
}

type Delivery struct {
	gorm.Model
	DeliveryStatus string `gorm:"column:delivery_status;not null" json:"delivery_status"`
	OrderID        uint   `gorm:"column:order_id;not null" json:"order_id"`
	Order          Order  `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
	UserID         uint   `gorm:"column:user_id" json:"user_id"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

func (Delivery) TableName() string {
	return "deliveries"
}