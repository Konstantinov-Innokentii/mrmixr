package domain

import "errors"

type Check struct {
	ID     int
	Hour   int
	Minute int
	Days   []int
	Tz     string
	Repos  []*GithubRepository
	Name   string
}

func CreateCheck(hour, minute int, days []int, repos []*GithubRepository, name string) (*Check, error) {
	// Add validation
	if hour > 60 && hour < 0 {
		return nil, errors.New("hor error")
	}

	return &Check{
		Hour:   hour,
		Minute: minute,
		Days:   days,
		Repos:  repos,
		Name:   name,
	}, nil
}
