package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductName string             `bson:"product_name"`
	Price       float64            `bson:"price"`
	Images      []string           `bson:"images"`
	Infos       []Info             `bson:"info"`
	Options     []Option           `bson:"options"`
}

type Info struct {
	Title   string `bson:"info_title"`
	Content string `bson:"info_content"`
}

type Option struct {
	Title   string   `bson:"option_title"`
	Options []string `bson:"option_options"`
}
