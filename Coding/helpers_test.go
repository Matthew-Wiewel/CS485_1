package coding

import "testing"

func TestTypeMapMakeB(t *testing.T) {
	AutofillTypeMap()
	var result AutoStruct = typeFactory(bID)
	var bStruct B
	var expected AutoStruct = &bStruct
	if result != expected {
		t.Errorf("Creating a struct of type B does not work")
	}
}
