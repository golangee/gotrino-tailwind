// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package progress

import (
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	s "github.com/golangee/gotrino-html/svg"
	"github.com/golangee/property"
)

type Circle struct {
	View
	visible property.Bool
}

func NewInfiniteCircle() *Circle {
	c := &Circle{}
	c.visible.Attach(c.Invalidate)
	return c
}

func (c *Circle) VisibleProperty() *property.Bool {
	return &c.visible
}

func (c *Circle) Render() Node {
	return s.Svg(Class("text-primary stroke-current wtk-mdc-circular-progress"), s.ViewBox("25 25 50 50"),
		s.Circle(Class("wtk-mdc-circular-progress__path"), s.Cx("50"), s.Cy("50"), s.R("20"), s.Fill("none"), s.StrokeWidth("4"), s.StrokeMiterlimit("10")),
	)
}
