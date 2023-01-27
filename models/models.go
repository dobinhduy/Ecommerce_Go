package models

import (
	"time"
	"go.mongodb.org/mongo-diver/bson/primitive"
)

type User struct{
	ID  				primitive.ObjectID 		`json:"_id" bson:"_id"`
	First_Name			*string					`json:"first_name"		 validate:"required,min =2, max = 30"`
	Last_Name			*string					`json:"last_name" 		 validate:"required,min =2, max = 30"`
	Password 			*string					`json:"password" 		 validate:"required,min =6"`
	Email				*string					`json:"email" 			 validate:"email,required"`
	Phone				*string					`json:"phone" 			 validate:"required"`
	Token				*string					`json:"token"`
	Refresh_Token		*string					`json:"refresh_token"`
	Create_At			time.Time				`json:"create_at"`
	Update_At			time.Duration			`json:"update_at"`
	User_ID				*string					`json:"user_id"`
	UserCart			[]ProductUser			`json:"user_cart" bson:"usercart"`
	Address_Details		[]Address				`json:"address" bson:"address"`
	Order_Status 		[]Order					`json:"orders" bson:"orders"`
}
type Product struct{
	Product_ID         	primitive.ObjectID      `bson:"_id"`
	Product_Name		*string					`json:"product_name"`
	Price				*uint64					`json:"price"`
	Rating				*string					`json:"rating"`
	Image				*string					`json:"image"`
}
type ProductUser struct{
	Product_ID			primitive.ObjectID		`bson:"_id"`
	Product_Name		*string					`json:"product_name" bson:"product_name"`
	Price				int						`json:"price" bson:"price"`
	Rating				*uint					`json:"rating" bson:"rating"`
	Image				*string					`json:"image" bson: "image"`
}
type Address struct{
	Address_ID			primitive.ObjectID		`bson:"_id"`
	House				*string					`json:"house_name" bson:"house_name"`
	Street				*string					`json:"street_name" bson:"street_name"`
	City				*string					`json:"city_name" bson:"city_name"`
	PinCode				*string					`json:"pin_code" bson:"pin_code"`
}
type Order struct{
	Order_ID 			primitive.ObjectID		`bson:"_id"`
	Order_Cart			[]ProductUser			`json:"order_list" bson:"order_list"`
	Order_At			time.Time				`json:"order_at" bson:"order_at"`
	Price				int						`json:"total_price" bson:"total_price"`
	Discount			*int					`json:"discount" bson:"discount"`
	Payment_Method		Payment					`json:"payment_method" bson:"payment_method"`
}
type Payment struct{
	Digital bool
	COD	bool
}