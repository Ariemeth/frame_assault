package weapon

import "testing"

type testTarget struct {
	DamageTaken int
}

func (fakeTarget *testTarget) Hit(damage int) {
	fakeTarget.DamageTaken += damage
}

func TestWeapon(t *testing.T) {
	weapon1 := Create(2, 2, "test weapon1", 0.6)

	if weapon1.damage != 2 {
		t.Errorf("weapon1 damage is %d instead of 2", weapon1.damage)
	}
	if weapon1.maxRange != 2 {
		t.Errorf("weapon1 range is %d instead of 2", weapon1.maxRange)
	}
}

func TestWeaponFire(t *testing.T) {
	weapon1 := Create(2, 2, "test weapon1", 0.6)

	target := testTarget{}

	weapon1.Fire(3, &target)
	if target.DamageTaken != 0 {
		t.Errorf("mech destroyed at range 3 by range 2 weapon")
	}

	weapon1.Fire(2, &target)
	if target.DamageTaken != 2 {
		t.Errorf("mech not destroyed at range 2 by range 2, damage 2 weapon")
	}
}
