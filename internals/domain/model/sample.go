package model

import "fmt"

var (
	ErrRequiredName = fmt.Errorf("name is required")
)

type Sample struct {
	ID   int    `json:"id"`
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type SampleFindParams struct {
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
