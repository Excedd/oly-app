package domain

type PersonalRecord struct {
	Snatch       float64 `json:"snatch"`
	CleanAndJerk float64 `json:"clean_and_jerk"`
	Jerk         float64 `json:"jerk"`
	User         string  `json:"user"`
}
