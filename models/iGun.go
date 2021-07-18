package models

type iGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
