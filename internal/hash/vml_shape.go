package hash

import (
	"github.com/plandem/ooxml/drawing/vml"
	"strconv"
	"strings"
)

//Vml return string with all required values of vml.Shape to build unique index
func Vml(shape *vml.Shape) Key {
	if shape == nil {
		shape = &vml.Shape{}
	}

	return Key(strings.Join([]string{
		shape.Type,
		strconv.FormatInt(int64(shape.Spt), 10),
		strconv.FormatInt(int64(shape.ClientData.Column), 10),
		strconv.FormatInt(int64(shape.ClientData.Row), 10),
	}, ":"))
}
