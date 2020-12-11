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

package button

import (
	v "github.com/golangee/gotrino"
	h "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/icon"
)

// A Button in the meaning of a material design text button.
// See also https://material.io/components/buttons#anatomy.
type Button struct {
	v.View
	content v.Renderable
	action  func()
}

func NewButton(action func()) *Button {
	c := &Button{
		action: action,
	}
	return c
}

func (c *Button) SetContent(content v.Renderable) *Button {
	c.content = content
	c.Invalidate()
	return c
}

func (c *Button) Render() v.Node {
	return h.Button(
		h.Style("min-width", "64px"),
		h.Class("text-left hover:bg-primary bg-opacity-10 bg-transparent text-primary focus:outline-none p-2 rounded text-center"), // not w-full
		h.AddClickListener(c.action),
		c.content,
	)
}

func NewIconTextButton(ico string, text string, action func()) *Button {
	return NewButton(action).SetContent(h.Span(icon.NewIcon(ico), h.Span(h.Class("pr-2")), h.Text(text)))
}

func NewTextButton(text string, action func()) *Button {
	return NewButton(action).SetContent(h.Span(h.Text(text)))
}
