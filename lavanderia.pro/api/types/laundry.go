package types

type Laundry struct {
	ID   string  `json:"id" bson:"_id,omitempty"`
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
