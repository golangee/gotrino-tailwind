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

package icon

import (
	. "github.com/golangee/gotrino"
	h "github.com/golangee/gotrino-html"
)

// Icon provides a component wrapper around a material-icon.
type Icon struct {
	name string
	View
}

func NewIcon(name string) *Icon {
	return &Icon{name: name}
}

func (c *Icon) Render() Node {
	return h.I(h.Class("material-icons align-sub"), h.Text(c.name))
}
