package models

type musket struct {
	gun
}

func NewMusket() iGun {
	return &musket{
		gun: gun{
			Name:  "musket",
			Power: 1,
		},
	}
}
