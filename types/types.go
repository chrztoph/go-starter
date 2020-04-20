package types

import "github.com/go-openapi/strfmt"

// basic swagger model properties: https://goswagger.io/use/spec/model.html#properties
// special types: https://goswagger.io/use/spec/strfmt.html

// swagger:model
type SampleEntry struct {

	// Min Length: 5
	// Max Length: 10
	Data string `json:"data"`

	ID strfmt.UUID `json:"id"`

	// Min: 1
	// Max: 100
	Num int `json:"num"`

	// Default: false
	// Required: true
	IsActive bool `json:"isActive"`

	Mail strfmt.Email `json:"mail"`

	MoreData string `json:"moreData"`
}

// swagger:model
type HelloWorld struct {

	// Allow: "World"
	// Required: true
	Hello string `json:"hello"`
}