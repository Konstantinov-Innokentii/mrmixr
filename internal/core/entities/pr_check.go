package entities

type Weekday int64

const (
	MON Weekday = iota + 1
	TUE
	WED
	THU
	FRI
	SAT
	SUN
)

type PRCheck struct {
	Id     int
	Hour   int
	Minute int
	Days   []Weekday
	Tz     string
	Repos  *[]GithubRepository
}

func CreateCheck(hour, minute int, days []Weekday, repos *[]GithubRepository) (*PRCheck, error) {
	// Add validation
	return &PRCheck{
		Hour:   hour,
		Minute: minute,
		Days:   days,
		Repos:  repos,
	}, nil
}
