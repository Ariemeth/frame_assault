package weapon

// Weapon is weapon with specific characteristics
type Weapon struct {
	maxRange, damage int
	name             string
	hitRate          float64
}

// Target is an interface used by objects that can be hit and take damage
type Target interface {
	Hit(int)
}

// Create creates a new Weapon.
func Create(maxRange int, damage int, name string,
	hitRate float64) Weapon {

	return Weapon{maxRange: maxRange, damage: damage, name: name,
		hitRate: hitRate}
}

//Name returns the name of the weapon
func (weapon Weapon) Name() string {
	return weapon.name
}

//Range returns the range of the weapon
func (weapon Weapon) Range() int {
	return weapon.maxRange
}

//Damage returns the damage of the weapon
func (weapon Weapon) Damage() int {
	return weapon.damage
}

//Accuracy returns the accuracy of the weapon
func (weapon Weapon) Accuracy() float64 {
	return weapon.hitRate
}

// Fire is used by an object to fire at a Target.
// Requires the range to the Target and the Target.
func (weapon Weapon) Fire(rangeToTarget int, target Target) {
	if rangeToTarget <= weapon.maxRange {
		target.Hit(weapon.damage)
	}
}
