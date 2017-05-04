package main

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func main() {
	type FieldOpts struct {
		Name string `url:"name"`
	}
	type Options struct {
		SSM map[string]string    `url:"ssm,indexed"`
		SrM map[string]FieldOpts `url:"srm,indexed"`
	}
	opt := Options{
		SSM: map[string]string{
			"key_a": "v_a",
			"key_b": "v_b",
		},
		SrM: map[string]FieldOpts{
			"key_x": {
				Name: "y",
			},
			"key_y": {
				Name: "z",
			},
		},
	}
	v, _ := query.Values(opt)
	// fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	ss, _ := url.QueryUnescape(v.Encode())
	fmt.Print(ss) // will output: "q=foo&all=true&page=2"
}
