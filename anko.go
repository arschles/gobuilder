package gobuilder

import (
	"encoding/json"
)

type context struct {
	obj map[string]interface{}
}

func newContext() *context {
	return &context{obj: map[string]interface{}{}}
}

func (c *context) add(name string, data interface{}) {
	c.obj[name] = data
}

func (c *context) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.obj)
}
