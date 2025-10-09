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

type UserRole string
const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)
type User struct {
	gorm.Model
	Name        string   `gorm:"column:name;not null"  json:"name"`
	Email       string   `gorm:"column:email;unique;not null" json:"email"`
	Password    string   `gorm:"column:password;not null" json:"password"`
	Role        UserRole   `gorm:"column:role;type:varchar(20);default:customer" json:"role"`
	PhoneNumber string   `gorm:"column:phone_number;not null" json:"phone_number"`
	Address     string   `gorm:"column:address;" json:"address"`
	Orders      []Order  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orders"`
	Reviews     []Review `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
}

type Order struct {
	gorm.Model
	ShippingAddress string `gorm:"column:shipping_address;not null" json:"shipping_address"`
	UserID          uint   `gorm:"column:user_id;not null" json:"user_id"`
	ProductID       uint   `gorm:"column:user_id;not null" json:"product_id"`
	User            User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}

type Payment struct {
	gorm.Model
	OrderID       uint   `gorm:"column:order_id;not null" json:"order_id"`
	UserID        uint   `gorm:"column:user_id;not null" json:"user_id"`
	PaymentStatus string   `gorm:"column:payment_status;not null" json:"payment_status"`
	PaymentMethod string `gorm:"column:payment_method;not null" json:"payment_method"`
	TransactionID uint   `gorm:"column:transaction_id;not null" json:"transaction_id"`
}

type Review struct {
	gorm.Model
	Rating    uint    `gorm:"column:rating;not null" json:"rating"`
	Review    string  `gorm:"column:review" json:"review"`
	UserID    uint    `gorm:"column:user_id;uniqueIndex:review_idx" json:"user_id"` //removed not null constraint
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ProductID uint    `gorm:"column:product_id;uniqueIndex:review_idx" json:"product_id"` //removed not null constraint
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
}

type Delivery struct {
	gorm.Model
	DeliveryStatus string `gorm:"column:delivery_status;not null" json:"delivery_status"`
	OrderID        uint   `gorm:"column:order_id;not null" json:"order_id"`
	Order          Order  `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
}

func (Delivery) TableName() string {
	return "deliveries"
}