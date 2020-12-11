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

package text

import (
	v "github.com/golangee/gotrino"
	h "github.com/golangee/gotrino-html"
)

type Text struct {
	text string
	rs   []v.Renderable
	v.View
}

func NewText(text string, rs ...v.Renderable) *Text {
	return &Text{
		text: text,
		rs:   rs,
	}
}

func (c *Text) Render() v.Node {
	return h.Span(append([]v.Renderable{h.Text(c.text)}, c.rs...)...)
}

func (c *Text) SetText(text string) {
	c.text = text
	c.Invalidate()
}

func (c *Text) Text() string {
	return c.text
}
