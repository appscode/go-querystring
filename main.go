package main

import (
	"github.com/google/go-querystring/query"
	"fmt"
)

func main() {
	type FieldOpts struct {
		FQuery   string `url:"fq"`
		FShowAll bool   `url:"fall"`
		FPage    int    `url:"fpage"`
		FStrArr  []string  `url:"fsa"`
		M2    map[string]string `url:"m2"`
	}
	type Options struct {
		Query   string `url:"q"`
		ShowAll bool   `url:"all"`
		Page    int    `url:"page"`
		Fields  FieldOpts         `url:"fields"`
		Maps    map[string]string `url:"maps"`
	}
	opt := Options{
		Query:   "foo",
		ShowAll: true,
		Page:    2,
		Maps: map[string]string{
			"f_1": "p_1",
		},
	}
	opt.Fields = FieldOpts{
		FQuery:   "a",
		FShowAll: true,
		FPage:    10,
		FStrArr:  []string{"b100", "c100"},
		M2: map[string]string{
			"f_2": "p_2",
		},
	}
	v, _ := query.Values(opt)
	fmt.Print(v.Encode()) // will output: "q=foo&all=true&page=2"
}
