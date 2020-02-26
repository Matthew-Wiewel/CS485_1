package coding

import "testing"

func TestGetTypeIdB(t *testing.T) {
	var bStruct B
	if bStruct.getTypeId() != bID {
		t.Errorf("bstruct.getTypeId returns %d, not %d", bStruct.getTypeId(), bID)
	}
}

func TestGetTypeIdC(t *testing.T) {
	var cStruct C
	if cStruct.getTypeId() != cID {
		t.Errorf("cstruct.getTypeId returns %d, not %d", cStruct.getTypeId(), cID)
	}
}

func TestGetTypeIdD(t *testing.T) {
	var dStruct D
	if dStruct.getTypeId() != dID {
		t.Errorf("dstruct.getTypeId returns %d, not %d", dStruct.getTypeId(), dID)
	}
}

func TestGetTypeIdE(t *testing.T) {
	var eStruct E
	if eStruct.getTypeId() != eID {
		t.Errorf("estruct.getTypeId returns %d, not %d", eStruct.getTypeId(), eID)
	}
}