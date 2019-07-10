// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFontUnderline(t *testing.T) {
	type Element struct {
		Property primitives.UnderlineType `xml:"property,omitempty"`
	}

	list := map[string]primitives.UnderlineType{
		"single":           styles.UnderlineTypeSingle,
		"double":           styles.UnderlineTypeDouble,
		"singleAccounting": styles.UnderlineTypeSingleAccounting,
		"doubleAccounting": styles.UnderlineTypeDoubleAccounting,
		"none":             styles.UnderlineTypeNone,
		"underline-a":      primitives.UnderlineType("underline-a"),
	}

	for s, v := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Element{Property: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Element><property val="%s"></property></Element>`, s), string(encoded))

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
