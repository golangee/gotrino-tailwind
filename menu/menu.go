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

package menu

import (
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/modal"
)

// Menu provides the basic frame and design for a material design popup menu.
type Menu struct {
	View
	content Renderable
}

func NewMenu(content Renderable) *Menu {
	return &Menu{content: content}
}

func (c *Menu) Render() Node {
	return Div(Class("w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5"),
		Style("max-width", "90vw"),
		Div(Class("py-1"), Role("menu"), AriaOrientation("vertical"),
			c.content,
		),
	)
}

// MenuItem provides the basic component for a simple entry in a Menu.
type MenuItem struct {
	View
	content Renderable
}

func NewMenuItem(content Renderable) *MenuItem {
	return &MenuItem{content: content}
}

func (c *MenuItem) Render() Node {
	return A(Class("block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"), Role("menuitem"),
		c.content,
	)
}

// ShowPopup tries to popup the content intelligent around the given anchor, considering window size and other
// alignment rules. The Popup is closed automatically when clicked outside. The anchor must denote a valid ID.
func ShowPopup(anchorID string, menuContent Renderable) {
	modal.ShowOverlay(modal.NewOverlay().Put(anchorID, menuContent))
}
