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
		//Field FieldOpts         `url:"field"`
		Arr []FieldOpts `url:"arr,indexed"`
	}
	opt := Options{
		//Field: FieldOpts{
		//	Name: "x",
		//},
		Arr: []FieldOpts{
			{
				Name: "y",
			},
			{
				Name: "z",
			},
		},
	}
	v, _ := query.Values(opt)
	// fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
	ss, _ := url.QueryUnescape(v.Encode())
	fmt.Print(ss) // will output: "q=foo&all=true&page=2"
}