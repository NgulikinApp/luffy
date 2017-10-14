package product

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Price string        `bson:"price" json:"price"`
	Image Image         `bson:"image" json:"image"`
}

type Image struct {
	URL    string `bson:"url" json:"url"`
	Width  int64  `bson:"width" json:"width"`
	Height int64  `bson:"height" json:"height"`
	Mime   string `bson:"mime" json:"mime"`
}

type Filter struct {
	Tag string
}
