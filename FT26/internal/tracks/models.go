package tracks

import "go.mongodb.org/mongo-driver/bson/primitive"

type Circuit struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Country     string             `bson:"country" json:"country"`
	YearBuilt   int                `bson:"yearBuilt" json:"yearBuilt"`
	LengthKM    float64            `bson:"lengthKm" json:"lengthKm"`
	Corners     int                `bson:"corners" json:"corners"`
	Description string             `bson:"description" json:"description"`
}
