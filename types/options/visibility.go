// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for VisibilityType
const (
	_ primitives.VisibilityType = iota
	Visible
	Hidden
	VeryHidden
)

func init() {
	primitives.FromVisibilityType = map[primitives.VisibilityType]string{
		Visible:    "visible",
		Hidden:     "hidden",
		VeryHidden: "veryHidden",
	}

	primitives.ToVisibilityType = make(map[string]primitives.VisibilityType, len(primitives.FromVisibilityType))
	for k, v := range primitives.FromVisibilityType {
		primitives.ToVisibilityType[v] = k
	}
}
