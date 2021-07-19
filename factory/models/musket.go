package models

type musket struct {
	gun
}

func NewMusket() IGun {
	return &musket{
		gun: gun{
			Name:  "musket",
			Power: 1,
		},
	}
}
