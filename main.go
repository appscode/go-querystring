package main

import (
	"github.com/google/go-querystring/query"
	"fmt"
)

func main() {
	type FieldOpts struct {
		FStrArr  []string  `url:"fsa,comma,numbered"`
		M2    map[string]string `url:"m2"`
	}
	type Options struct {
		Fields  []FieldOpts         `url:"fields"`
		Maps    map[string]string `url:"maps"`
	}
	opt := Options{
		Maps: map[string]string{
			"f_1": "p_1",
		},
	}
	opt.Fields = []FieldOpts{
		{
			FStrArr: []string{"b100", "c100"},
			M2: map[string]string{
				"f_2": "p_2",
			},
		},
	}
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
}
