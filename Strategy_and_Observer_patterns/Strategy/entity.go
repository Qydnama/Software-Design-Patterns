package main

type Entity struct {
	name       string
	attackType Attack
	duration   float64
	damage     float64
}

func initEntity(name string) *Entity {
	return &Entity{name, &NoWeapon{}, 0, 0}
}
func (e *Entity) setWeapon(a Attack, dur float64, dam float64) {
	e.attackType = a
	e.duration = dur
	e.damage = dam
}

func (e *Entity) attack() {
	e.attackType.attack(e)
}
