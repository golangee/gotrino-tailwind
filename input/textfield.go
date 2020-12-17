package input

import (
	"github.com/golangee/dom"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/property"
)

// TextField provides a material design style for text fields and related html5 variants. However, note
// that a native date picker is not available at least in Safari desktop and looks broken on iOS.
type TextField struct {
	text                   property.String
	label                  property.String
	floatTop               property.Bool
	focus                  property.Bool
	inputType              property.String
	name                   property.String
	placeholderDatePattern string
	View
}

func NewTextField() *TextField {
	c := &TextField{}
	c.label.Attach(c.Invalidate)
	c.inputType.Set("text")
	c.inputType.Attach(c.Invalidate)
	c.name.Set(dom.GenerateID())
	c.name.Attach(c.Invalidate)
	c.focus.Observe(func(old, new bool) {
		c.floatTop.Set(new || c.text.Get() != "")
	})
	c.text.Observe(func(old, new string) {
		// date: non-empty text to fix non-empty placeholders in chrome and ff - Safari is different
		if c.inputType.Get() == "date" {
			c.floatTop.Set(true)
			return
		}

		c.floatTop.Set(new != "")
	})
	c.inputType.Observe(func(old, new string) {
		// date: non-empty text to fix non-empty placeholders in chrome and ff - Safari is different
		if new == "date" {
			c.placeholderDatePattern = "yyyy-mm-dd"
			c.floatTop.Set(true)
		} else {
			c.placeholderDatePattern = " "
			c.Invalidate()
		}
	})
	return c
}

func (c *TextField) TextProperty() *property.String {
	return &c.text
}

func (c *TextField) NameProperty() *property.String {
	return &c.name
}

func (c *TextField) SetName(t string) *TextField {
	c.NameProperty().Set(t)
	return c
}

func (c *TextField) SetText(t string) *TextField {
	c.TextProperty().Set(t)
	return c
}

func (c *TextField) BindText(dst *string) *TextField {
	c.TextProperty().Bind(dst)
	return c
}

func (c *TextField) LabelProperty() *property.String {
	return &c.label
}

func (c *TextField) SetLabel(t string) *TextField {
	c.LabelProperty().Set(t)
	return c
}

func (c *TextField) TypeProperty() *property.String {
	return &c.inputType
}

func (c *TextField) SetType(t string) *TextField {
	c.TypeProperty().Set(t)
	return c
}

func (c *TextField) Render() Node {
	labelFocus := "scale-75 -translate-y-4 z-0 ml-3 px-1 py-0 bg-surface"

	return Div(Class("box-border rounded outline relative focus-within:border-primary"),
		IfCond(&c.focus,
			Modifiers(
				RemoveClass("hover:border-black border"),
				AddClass("border-2"),
			),
			Modifiers(
				AddClass("hover:border-black border"),
				RemoveClass("border-2"),
			)),
		Input(Class("block p-4 w-full text-lg appearance-none focus:outline-none bg-transparent"),
			AddEventListener("focus", func() {
				c.focus.Set(true)
			}),
			AddEventListener("blur", func() {
				c.focus.Set(false)
			}),
			Observe(&c.text.Property, func(e dom.Element) Modifier {
				e.Set("value", c.text.Get())
				return nil
			}),
			InsideDom(func(e dom.Element) {
				e.AddEventListener("input", false, func() {
					c.text.Set(e.Get("value").(string))
				})
			}),
			Type(c.inputType.Get()), Name(c.name.Get()), Placeholder(c.placeholderDatePattern),
		),
		Label(Class("absolute top-0 p-4 text-lg -z-1 duration-300 origin-0 transform"),
			Style("pointer-events", "none"),
			IfCond(&c.floatTop,
				Modifiers(
					AddClass(labelFocus),
					Style("left", "-0.75rem"),
				),
				Modifiers(
					RemoveClass(labelFocus),
					Style("left", "inherit"),
				)),
			For(c.name.Get()), Text(c.label.Get()),
		),
	)
}
