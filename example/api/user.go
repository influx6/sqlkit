package api

// User contains user data.
// @sqlapi
type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
}
