package models

type ak47 struct {
	gun
}

func NewAk47() IGun {
	return &ak47{
		gun: gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}
