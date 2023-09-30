package main

func main() {
	sword := &Sword{}
	entity1 := initEntity("zombie")
	entity1.setWeapon(sword, 4, 4)
	entity1.attack()
	//Attacking with sword!!

	bow := &Bow{}
	entity2 := initEntity("skeleton")
	entity2.setWeapon(bow, 6, 10)
	entity2.attack()
	//Attacking with bow!!!

	entity3 := initEntity("Steve")
	entity3.attack()
	//Attacking without weapon, just ARMS!!!

}
