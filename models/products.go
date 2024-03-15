package models


type Rating struct {
	Rate  float64 `json:"rate" bson:"rate"`
	Count int     `json:"count" bson:"count"`
}

type Product struct {
	ID          string   `json:"id,omitempty" bson:"_id"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Quantity    int      `json:"quantity" bson:"quantity"`
	Colors      []string `json:"colors" bson:"colors"`
	Sizes       []uint   `json:"sizes" bson:"sizes"`
	Images      []string `json:"images" bson:"images"`
	Category    string   `json:"category" bson:"category"`
	Price       float32  `json:"price" bson:"price"`
	Rating      Rating   `json:"rating" bson:"rating"`
}
