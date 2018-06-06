package gobuilder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

// Builder is the type that your ToBuilder method should return. A Builder
// knows how to take the data you give it, pass it into the builder DSL, and
// output the resulting JSON that your DSL produces
type Builder interface {
	// Execute passes the data into your DSL and returns the resulting JSON
	Execute() (string, error)
}

type builder struct {
	scriptName string
	dataName   string
	data       interface{}
}

// NewBuilder creates a new Builder that will pass data into the script
// with the name dataName. It will execute the
// 'templates/${dataName}/show.json.gobuilder' script by default
func NewBuilder(dataName string, data interface{}) Builder {
	return &builder{
		scriptName: fmt.Sprintf("templates/%s/show.json.gobuilder", dataName),
		dataName:   dataName,
		data:       data,
	}
}

func (b *builder) Execute() (string, error) {
	// get script file
	path, err := filepath.Abs(b.scriptName)
	if err != nil {
		return "", err
	}
	script, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// create env and add funcs to it
	ctx := newContext()
	env := vm.NewEnv()
	if err := env.Define(b.dataName, b.data); err != nil {
		return "", err
	}
	if err := env.Define("add", ctx.add); err != nil {
		return "", err
	}

	// execute the script
	_, err = env.Execute(string(script))
	if err != nil {
		switch t := err.(type) {
		case *vm.Error:
			return "", fmt.Errorf("%s on line %d, col %d", t.Message, t.Pos.Line, t.Pos.Column)
		case *parser.Error:
			return "", fmt.Errorf("invalid script on line %d, col %d: %s", t.Pos.Line, t.Pos.Column, t.Message)
		default:
			return "", err
		}
	}

	jsonBytes, err := json.Marshal(ctx)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
