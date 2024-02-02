package buffer

import "testing"

func TestNewPieceTable(t *testing.T) {
	origin := []rune("Hello, world!")
	pt := NewPieceTable(origin)
	if pt == nil {
		t.Error("NewPieceTable returned nil")
	}

	if string(pt.origin) != string(origin) {
		t.Errorf("NewPieceTable did not set origin correctly: %v", pt.origin)
	}

	if len(pt.add) != 0 {
		t.Errorf("NewPieceTable did set 'add' propert on init. Expected '' got '%s'", string(pt.add))
	}

	if len(pt.pieces) != 0 {
		t.Errorf("NewPieceTable did set 'pieces' propert on init. Expected 0 got %d", len(pt.pieces))
	}
}

func TestInsertFront(t *testing.T) {
	origin := []rune("there, world!")
	pt := NewPieceTable(origin)
	expectedAdd := "Hello"
	pt.Insert(expectedAdd, 0)

	if string(pt.origin) != string(origin) {
		t.Errorf("Insert changed origin: %v", pt.origin)
	}
	if string(pt.add) != expectedAdd {
		t.Errorf("Insert did not set added correctly. Expected '%s' got '%s' ", expectedAdd, string(pt.add))
	}
	if len(pt.pieces) != 2 {
		t.Errorf("Insert did not set pieces correctly. Expected 2 got %d", len(pt.pieces))
	}
}

func TestInsertEnd(t *testing.T) {
	origin := []rune("Hello there")
	pt := NewPieceTable(origin)
	expectedAdd := ", World"
	pt.Insert(expectedAdd, 11)

	if string(pt.origin) != string(origin) {
		t.Errorf("Insert changed origin: %v", pt.origin)
	}
	if string(pt.add) != expectedAdd {
		t.Errorf("Insert did not set added correctly. Expected '%s' got '%s' ", expectedAdd, string(pt.add))
	}
	if len(pt.pieces) != 2 {
		t.Errorf("Insert did not set pieces correctly. Expected 2 got %d", len(pt.pieces))
	}
}

func TestInsertMid(t *testing.T) {
	origin := []rune("Hello, world!")
	pt := NewPieceTable(origin)
	expectedAdd := " there"
	pt.Insert(expectedAdd, 5)

	if string(pt.origin) != string(origin) {
		t.Errorf("Insert changed origin: %v", pt.origin)
	}
	if string(pt.add) != expectedAdd {
		t.Errorf("Insert did not set added correctly. Expected '%s' got '%s' ", expectedAdd, string(pt.add))
	}
	if len(pt.pieces) != 3 {
		t.Errorf("Insert did not set pieces correctly. Expected 3 got %d", len(pt.pieces))
	}
}
