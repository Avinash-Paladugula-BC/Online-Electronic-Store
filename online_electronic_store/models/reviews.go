package models

// import "gorm.io/gorm"

// type Review struct {
// 	gorm.Model
// 	Rating    uint    `gorm:"column:rating;not null" json:"rating"`
// 	Review    string  `gorm:"column:review" json:"review"`
// 	UserID    uint    `gorm:"column:user_id;uniqueIndex:review_idx" json:"user_id"` //removed not null constraint
// 	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
// 	ProductID uint    `gorm:"column:product_id;uniqueIndex:review_idx" json:"product_id"` //removed not null constraint
// 	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product"`
// }
