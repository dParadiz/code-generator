package main

import "io/ioutil"

type Context struct {
	Content string
}
type encoderContext string

func (ec encoderContext) Encode(input string) (interface{}, error) {

	dat, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}

	return Context{Content: string(dat)}, nil
}

// Encoder exported symbol
var Encoder encoderContext
