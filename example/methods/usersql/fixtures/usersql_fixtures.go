package fixtures

import (
     "encoding/json"


     "github.com/gokit/sqlkit/example/methods"

)


// json fixtures ...
var (
	UserJSON = `{


    "public_id":	"f0ywplalulqu9kqybapmfdkamh6qd4",

    "name":	"Bobby Boyd"

}`
)

// LoadUserJSON returns a new instance of a methods.User.
func LoadUserJSON(content string) (methods.User, error) {
	var elem methods.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return methods.User{}, err
	}

	return elem, nil
}

