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

package modal

import (
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

// A DialogCard represents a material design-like card view, providing a minimal width, shadow and font.
type DialogCard struct {
	View
	title   string
	content gotrino.Renderable
}

func NewDialogCard(title string, content Renderable) *DialogCard {
	return &DialogCard{title: title, content: content}
}

func (c *DialogCard) Render() Node {
	return Div(Class("rounded-md shadow-xl bg-white m-auto p-6 pb-2"),
		Style("min-width", "240px"),
		AddClickListener(nil), // intentionally block any click events
		If(c.title != "",
			P(Class("text-xl font-medium"),
				Text(c.title),
			),
			nil,
		),

		c.content,
	)
}
