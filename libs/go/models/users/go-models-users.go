package go_models_users

type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	//Role float64 `json:"salary"`
	//Age    float64 `json:"age"`
}
