package hash

import (
	"github.com/plandem/xlsx/internal/ml"
	"strconv"
	"strings"
)

//NumberFormat return string with all values of number format
func NumberFormat(format *ml.NumberFormat) string {
	if format == nil {
		format = &ml.NumberFormat{}
	}

	return strings.Join([]string{
		strconv.FormatInt(int64(format.ID), 10),
		format.Code,
	}, ":")
}
