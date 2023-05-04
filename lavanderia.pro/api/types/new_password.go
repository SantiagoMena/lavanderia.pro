package types

type NewPassword struct {
	ID          string `json:"id" bson:"_id,omitempty" uri:"id"`
	Password    string `json:"password,omitempty" bson:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty" bson:"new_password,omitempty"`
}
