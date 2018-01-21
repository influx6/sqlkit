package fixtures

import (
     "encoding/json"


     "github.com/gokit/sqlkit/example/api"

)


// json fixtures ...
var (
	UserJSON = `{


    "public_id":	"oyhvza9uue5fkfsmne6qa7t3sjq84n",

    "name":	"Raymond Hernandez"

}`
)

// LoadUserJSON returns a new instance of a api.User.
func LoadUserJSON(content string) (api.User, error) {
	var elem api.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return api.User{}, err
	}

	return elem, nil
}

