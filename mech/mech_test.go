package mech

import (
	"testing"

	"github.com/Ariemeth/frame_assault/mech/weapon"
	tl "github.com/Ariemeth/termloop"
)

func TestNewMech(t *testing.T) {

	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created",
			mechName)
	}

	if mech1.StructureLeft() != structure {
		t.Errorf("%s does not have %d structure",
			mechName,
			structure)
	}

	if mech1.name != mechName {
		t.Errorf("%s does not have the correct name",
			mechName)
	}
}

func TestHit(t *testing.T) {

	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created",
			mechName)
	}

	mech1.Hit(0)
	if mech1.structure != structure {
		t.Errorf("%s took damage when it was hit with 0",
			mechName)
	}

	mech1.Hit(structure)
	if mech1.structure != 0 {
		t.Errorf("%s was not destroyed by taking %d damage",
			mechName,
			structure)
	}
}

func TestStructureLeft(t *testing.T) {
	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created",
			mechName)
	}

	if mech1.StructureLeft() != structure {
		t.Errorf("%s has %d structure and StructureLeft indicates %d",
			mechName,
			mech1.structure,
			mech1.StructureLeft())
	}
}

func TestAddWeapon(t *testing.T) {
	const mechName string = "testMech"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created",
			mechName)
	}
	weapon1 := weapon.Create(5, 1, "rifle", .75)

	mech1.AddWeapon(weapon1)

	if len(mech1.weapons) != 1 {
		t.Errorf("%s has %d weapons instead of 1 weapon",
			mech1.name,
			len(mech1.weapons))
	}
}

func TestMechFireInRange(t *testing.T) {
	const mechName string = "testMech"
	const mechName2 string = "testMech2"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created", mechName)
	}

	mech2 := NewMech(mechName2, structure, 0, 0, tl.ColorRed, 'T')
	if mech2 == nil {
		t.Errorf("%s was unable to be created", mechName2)
	}

	weapon1 := weapon.Create(5, 1, "rifle", .75)

	mech1.AddWeapon(weapon1)

	mech1.Fire(3, mech2)
	if mech2.structure != 1 {
		t.Errorf("%s fired at %s at range %d and %s still has %d structure",
			mech1.name,
			mech2.name,
			3,
			mech2.name,
			mech2.structure)
	}
}

func TestMechFireOutOfRange(t *testing.T) {
	const mechName string = "testMech"
	const mechName2 string = "testMech2"
	const structure int = 2

	mech1 := NewMech(mechName, structure, 0, 0, tl.ColorRed, 'T')
	if mech1 == nil {
		t.Errorf("%s was unable to be created", mechName)
	}

	mech2 := NewMech(mechName2, structure, 0, 0, tl.ColorRed, 'T')
	if mech2 == nil {
		t.Errorf("%s was unable to be created", mechName2)
	}

	weapon1 := weapon.Create(5, 1, "rifle", .75)

	mech1.AddWeapon(weapon1)

	mech1.Fire(9, mech2)
	if mech2.structure != structure {
		t.Errorf("%s fired at %s at range %d and %s still has %d structure",
			mech1.name,
			mech2.name,
			9,
			mech2.name,
			mech2.structure)
	}
}

func TestMechPosition(t *testing.T) {

}

func TestMechSize(t *testing.T) {

}

func TestMechDraw(t *testing.T) {

}

func TestMechCollide(t *testing.T) {

}
