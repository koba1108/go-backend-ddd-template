package model

import "fmt"

var (
	ErrRequiredName = fmt.Errorf("name is required")
)

type Sample struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewSample(name string) (*Sample, error) {
	if name == "" {
		return nil, ErrRequiredName
	}
	return &Sample{
		Name: name,
	}, nil
}
