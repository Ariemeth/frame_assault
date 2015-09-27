package weapon

// CreateShotgun creates a new shotgun weapon
func CreateShotgun() Weapon {
	return Create(3, 2, "Shotgun", .50)
}

// CreateRifle creates a new rifle weapon
func CreateRifle() Weapon {
	return Create(5, 1, "Rifle", .75)
}

// CreateFist creates a new fist weapon
func CreateFist() Weapon {
	return Create(1, 1, "Fist", .60)
}

// CreateSword creates a new sword weapon
func CreateSword() Weapon {
	return Create(1, 2, "Sword", .80)
}
