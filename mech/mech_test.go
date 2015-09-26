package mech

import "testing"

func TestNewMech(t *testing.T) {

	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure)
	if mech1 == nil {
		t.Errorf("%s was unable to be created", mechName)
	}

	if mech1.StructureLeft() != structure {
		t.Errorf("%s does not have %d structure", mechName, structure)
	}

	if mech1.name != mechName {
		t.Errorf("%s does not have the correct name", mechName)
	}
}

func TestHit(t *testing.T) {

	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure)
	if mech1 == nil {
		t.Errorf("%s was unable to be created", mechName)
	}

	mech1.Hit(0)
	if mech1.structure != structure {
		t.Errorf("%s took damage when it was hit with 0", mechName)
	}

	mech1.Hit(structure)
	if mech1.structure != 0 {
		t.Errorf("%s was not destroyed by taking %d damage", mechName, structure)
	}
}

func TestStructureLeft(t *testing.T) {

}

func TestAddWeapon(t *testing.T) {

}

func TestMechFire(t *testing.T) {

}
