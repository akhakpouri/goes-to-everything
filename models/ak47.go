package models

type ak47 struct {
	gun
}

func NewAk47() iGun {
	return &ak47{
		gun: gun{
			Name:  "AK47 gun",
			Power: 4,
		},
	}
}
