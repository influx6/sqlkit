package methods

// User contains user data.
// @sql_methods
type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
}

func (u User) Fields() (map[string]interface{}, error) {
	return nil, nil
}

func (u *User) Consume(data map[string]interface{}) error {
	return nil
}
