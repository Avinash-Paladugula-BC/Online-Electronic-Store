package models

import "gorm.io/gorm"

//	{
//		"name": "Wireless Bluetooth Headphones",
//		"category": "Audio",
//		"price": 2499.99,
//		"brand": "SoundPulse",
//		"quantity": 50
//	 }
type Product struct {
	gorm.Model
	Name     string   `gorm:"column:name; uniqueIndex;not null" json:"name"`
	Category string   `gorm:"column:category; not null" json:"category"`
	Price    float64  `gorm:"column:price; not null" json:"price"`
	Brand    string   `gorm:"column:brand"  json:"brand"`
	Quantity uint     `gorm:"column:quantity; not null" json:"quantity"`
	OrderId  uint     `gorm:"column:order_id" json:"order_id"`
	Reviews  []Review `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"  json:"reviews"`
}

// {
// 	"name": "Avinash Paladugula",
// 	"email": "avinash@example.com",
// 	"password": "avinash",
// 	"role": "customer",
// 	"phone_number": "9876543210",
// 	"address": "Hyderabad, Telangana, India"
// }

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
	Role        UserRole `gorm:"column:role;type:varchar(20);check:role IN('admin','customer');default:customer" json:"role"`
	PhoneNumber string   `gorm:"column:phone_number;not null" json:"phone_number"`
	Address     string   `gorm:"column:address;" json:"address"`
	Orders      []Order  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orders"`
	Reviews     []Review `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
}

//	{
//		"shipping_address": "Hyderabad, Telangana, India",
//		"user_id": 1,
//		"product_id": 1
//	}
type Order struct {
	gorm.Model
	ShippingAddress string `gorm:"column:shipping_address;not null" json:"shipping_address"`
	UserID          uint   `gorm:"column:user_id;not null" json:"user_id"`
	// ProductID       uint   `gorm:"column:user_id;not null" json:"product_id"`
	Products []Product `gorm:foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:products`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	// PaymentID       uint       `gorm:"column:payment_id;not null" json:"payment_id"`
	// Payment         Payment    `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment"`
	//##################### multiple devlivery for single order doesn't exist
	// Deliveries []Delivery `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"deliveries"`
}

// {
// 	"order_id": 1,
// 	"user_id": 1,
// 	"payment_status": "Paid",
// 	"payment_method": "Credit Card",
// 	"transaction_id": "789456123"
// }

type Payment struct {
	gorm.Model
	OrderID uint `gorm:"column:order_id;not null" json:"order_id"` //refers to the orderid in the the order struct
	UserID  uint `gorm:"column:user_id;not null" json:"user_id"`
	// PaymentStatus bool   `gorm:"column:payment_status;not null" json:"payment_status"`
	PaymentStatus string `gorm:"column:payment_status;not null" json:"payment_status"`
	PaymentMethod string `gorm:"column:payment_method;not null" json:"payment_method"`
	TransactionID string `gorm:"column:transaction_id;not null" json:"transaction_id"`
}

//	{
//		"rating": 5,
//		"review": "Excellent sound quality and battery life!",
//		"user_id": 1,
//		"product_id": 1
//	}
type Review struct {
	gorm.Model
	Rating    uint    `gorm:"column:rating;not null" json:"rating"`
	Review    string  `gorm:"column:review" json:"review"`
	UserID    uint    `gorm:"column:user_id;uniqueIndex:review_idx" json:"user_id"` //removed not null constraint
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ProductID uint    `gorm:"column:product_id;uniqueIndex:review_idx" json:"product_id"` //removed not null constraint
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
}

//	{
//		"delivery_status": "Delivered",
//		"order_id": 1
//	}
type Delivery struct {
	gorm.Model
	DeliveryStatus string `gorm:"column:delivery_status;not null" json:"delivery_status"`
	OrderID        uint   `gorm:"column:order_id;not null" json:"order_id"`
	Order          Order  `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
}

func (Delivery) TableName() string {
	return "deliveries"
}
