package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductName string             `bson:"item_name"`
	Price       float64            `bson:"price"`
	Images      []string           `bson:"item_photos"`
	Infos       []Info             `bson:"info"`
	Variants    []Variant          `bson:"options"`
	Categories  []Category         `bson:"categories"`
}

type Info struct {
	Title   string `bson:"info_title"`
	Content string `bson:"info_content"`
}

type Variant struct {
	Title    string   `bson:"option_title"`
	Variants []string `bson:"option_options"`
}
