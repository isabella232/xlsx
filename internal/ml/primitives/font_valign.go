// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FontVAlignType is a type to encode XSD ST_VerticalAlignRun
type FontVAlignType ml.Property

//MarshalXML marshal FontVAlignType
func (t *FontVAlignType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal FontVAlignType
func (t *FontVAlignType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
