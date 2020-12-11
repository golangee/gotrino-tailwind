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

package appbar

import (
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	. "github.com/golangee/gotrino-html/svg"
	"github.com/golangee/property"
)

// The AppBar provides a drawer area (side menu) at the left and a toolbar area at the top. It has only limited
// capabilities for customization. IfCond you are sure, that you really need a custom AppBar, feel free to
// copy-paste to start a new component for your specific project.
type AppBar struct {
	isOpen       property.Bool
	toolbarArea  Renderable
	icon         Renderable
	title        Renderable
	drawerHeader Renderable
	drawerMain   Renderable
	drawerFooter Renderable
	View
}

// NewAppBar allocates a new AppBar instance.
func NewAppBar() *AppBar {
	c := &AppBar{}

	return c
}

// SetToolbarArea updates the node for right side of the AppBar. Consider mobile devices and only offer a context
// menu for small screens.
func (c *AppBar) SetToolbarArea(node Renderable) *AppBar {
	c.toolbarArea = node
	c.Invalidate()

	return c
}

// SetIcon sets a Node as the first entry right of the hamburger menu.
func (c *AppBar) SetIcon(node Renderable) *AppBar {
	c.icon = node
	c.Invalidate()

	return c
}

// Self assigns the receiver to the given reference.
func (c *AppBar) Self(ref **AppBar) *AppBar {
	*ref = c

	return c
}

// SetTitle sets a Node as the entry right of the Icon.
func (c *AppBar) SetTitle(node Renderable) *AppBar {
	c.title = node
	c.Invalidate()

	return c
}

// SetDrawerHeader sets a Node into the header section of the drawer. At least this should be the app icon.
func (c *AppBar) SetDrawerHeader(node Renderable) *AppBar {
	c.drawerHeader = node
	c.Invalidate()

	return c
}

// SetDrawerMain sets a Node as the drawers main content.
func (c *AppBar) SetDrawerMain(node Renderable) *AppBar {
	c.drawerMain = node
	c.Invalidate()

	return c
}

// SetDrawerFooter sets a Node into the bottom of the drawer.
func (c *AppBar) SetDrawerFooter(node Renderable) *AppBar {
	c.drawerFooter = node
	c.Invalidate()

	return c
}

// Close closes the side menu (also known as drawer).
func (c *AppBar) Close() *AppBar {
	c.isOpen.Set(false)

	return c
}

func (c *AppBar) Render() Node {
	return Div(
		Nav(Class("flex fixed w-full items-center justify-between px-6 h-12 bg-primary text-on-primary shadow z-10"),

			// menu and logo
			Div(Class("flex items-center"),

				// burger menu button
				Button(Class("focus:outline-none"), AriaLabel("Open Menu"),
					AddClickListener(c.isOpen.Toggle),
					Svg(
						Class("w-8 h-8"),
						Fill("none"),
						Stroke("currentColor"),
						StrokeLinecap("round"),
						StrokeLinejoin("round"),
						StrokeWidth("2"),
						ViewBox("0 0 24 24"),
						Path(D("M4 6h16M4 12h16M4 18h16")),
					),
				),

				// app logo in app bar
				c.icon,

				// app title
				c.title,
			),

			// button section in app bar
			Div(Class("flex items-center"),
				c.toolbarArea,
			),

			// semi-transparent content blocking layer
			Div(
				Class(" z-10 fixed ease-in-out inset-0 bg-black opacity-0 transition-all duration-500"),

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
					AddClickListener(c.isOpen.Toggle),
				),
			),

			// Side menu
			Aside(
				Class("transform top-0 left-0 w-64 bg-white fixed h-full overflow-auto ease-in-out transition-all duration-500 z-30"),

				IfCond(&c.isOpen,
					Modifiers(
						AddClass("translate-x-0"),
						RemoveClass("-translate-x-full"),
					),
					Modifiers(
						RemoveClass("translate-x-0"),
						AddClass("-translate-x-full"),
					),
				),

				// keep the logo in the menu
				Span(
					Class("flex w-full items-center p-4 border-b"),
					c.drawerHeader,
				),

				Div(
					c.drawerMain,
				),

				// button at the bottom in the side menu

				Div(
					c.drawerFooter,
				),
			),
		),
	)
}
