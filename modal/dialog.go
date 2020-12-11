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
	"github.com/golangee/dom"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/button"
	"github.com/golangee/property"
	"time"
)

// Dialog provides a half-translucent fullscreen view, which places any content just on top.
type Dialog struct {
	isOpen  property.Bool
	content Renderable
	View
}

func NewDialog() *Dialog {
	c := &Dialog{}
	// we do not attach the property, because we don't want a full re-rendering
	return c
}

func (c *Dialog) Show() *Dialog {
	c.isOpen.Set(true)
	return c
}

func (c *Dialog) Hide() *Dialog {
	c.isOpen.Set(false)
	return c
}

func (c *Dialog) SetContent(r Renderable) *Dialog {
	c.content = r
	return c
}

func (c *Dialog) Render() Node {
	// semi-transparent content blocking layer
	return Div(
		Div(Class(" z-10 fixed ease-in-out inset-0 bg-black opacity-0 transition-all duration-200"),
			IfCond(&c.isOpen,
				Modifiers(
					Style("visibility", "visible"),
					AddClass("opacity-50"),
				),
				Modifiers(
					Style("visibility", "hidden"),
					RemoveClass("opacity-50"),
				),
			),

			Div(
				Class("absolute inset-0"),
			),
		),

		Div(Class("absolute inset-0 ease-in-out transition-all opacity-0 duration-200 flex z-20"),
			AddClickListener(c.isOpen.Toggle),
			IfCond(&c.isOpen,
				AddClass("opacity-100"),
				RemoveClass("opacity-100"),
			),

			c.content,
		),
	)
}

func (c *Dialog) AddCloseListener(f func()) *Dialog {
	c.isOpen.Observe(func(old, new bool) {
		if new == false {
			f()
		}
	})

	return c
}

// ShowDialog creates a new partially opaque full screen plane and appends it to the body element.
// As soon as the dialog is closed it is removed from the DOM. You cannot show it again.
func ShowDialog(title string, content Renderable) *Dialog {
	dlg := NewDialog()
	dlg.SetContent(NewDialogCard(title, content))
	body := dom.GetWindow().Document().Body()
	WithElement(body, dlg).Element()
	dlgElem := body.LastChild()
	dom.Post(func() {
		dlg.Show()
	})

	dlg.AddCloseListener(func() {
		dom.SetTimeout(250*time.Millisecond, func() {
			dlgElem.Remove()
		})
	})

	return dlg
}

// ShowAlert just shows a message with an acknowledge or close button. See also ShowDialog.
func ShowAlert(title, msg, closeCaption string) *Dialog {
	var dlg *Dialog
	dlg = ShowAlertActions(title, msg, button.NewTextButton(closeCaption, func() {
		if dlg != nil {
			dlg.Hide()
		}
	}))

	return dlg
}

// ShowAlertActions shows a Dialog with the given actions. See also ShowDialog.
func ShowAlertActions(title, msg string, actions ...Renderable) *Dialog {
	return ShowDialog(title,
		Div(
			P(Class("pb-2 text-gray-500"), Text(msg)),
			Div(Class("flex flex-row-reverse"),
				ForEach(
					len(actions), func(i int) Renderable {
						return actions[len(actions)-1-i]
					},
				),
			),
		),
	)
}
