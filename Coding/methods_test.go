package coding

import "testing"

//method to test structs of type B return the bID
func TestGetTypeIdB(t *testing.T) {
	var bStruct B
	if bStruct.getTypeId() != bID {
		t.Errorf("bstruct.getTypeId returns %d, not %d", bStruct.getTypeId(), bID)
	}
}

//method to test strucsts of type C return the cID
func TestGetTypeIdC(t *testing.T) {
	var cStruct C
	if cStruct.getTypeId() != cID {
		t.Errorf("cstruct.getTypeId returns %d, not %d", cStruct.getTypeId(), cID)
	}
}

//method to test structs of type D return the dID
func TestGetTypeIdD(t *testing.T) {
	var dStruct D
	if dStruct.getTypeId() != dID {
		t.Errorf("dstruct.getTypeId returns %d, not %d", dStruct.getTypeId(), dID)
	}
}

//method to test structs of type E return the eID
func TestGetTypeIdE(t *testing.T) {
	var eStruct E
	if eStruct.getTypeId() != eID {
		t.Errorf("estruct.getTypeId returns %d, not %d", eStruct.getTypeId(), eID)
	}
}

func TestEncodeDecodeB_1(t *testing.T) {
	var original B
	original.ByteMember = 8
	encoding := original.Encode()
	var decoded B
	decoded.Decode(encoding)

	if decoded != original {
		t.Errorf("Inequality. Original had ByteMember value of %d, decoded has value of %d",
			original.ByteMember, decoded.ByteMember)
	}
}

func TestEncodeDecodeB_2(t *testing.T) {
	var original B
	original.ByteMember = 127
	encoding := original.Encode()
	var decoded B
	decoded.Decode(encoding)

	if decoded != original {
		t.Errorf("Inequality. Original had ByteMember value of %d, decoded has value of %d",
			original.ByteMember, decoded.ByteMember)
	}
}

func TestEncodeDecodeB_3(t *testing.T) {
	var original B
	original.ByteMember = 255
	encoding := original.Encode()
	var decoded B
	decoded.Decode(encoding)

	if decoded != original {
		t.Errorf("Inequality. Original had ByteMember value of %d, decoded has value of %d",
			original.ByteMember, decoded.ByteMember)
	}
}

func TestEncodeDecodeB_4(t *testing.T) {
	var original B
	original.ByteMember = 0
	encoding := original.Encode()
	var decoded B
	decoded.Decode(encoding)

	if decoded != original {
		t.Errorf("Inequality. Original had ByteMember value of %d, decoded has value of %d",
			original.ByteMember, decoded.ByteMember)
	}
}

func TestInterfaceB(t *testing.T) {
	originalStruct := B{5}
	var original AutoStruct = &originalStruct
	var decodedStruct B
	var decoded AutoStruct = &decodedStruct
	decoded.Decode(original.Encode())

	//compare strings instead of AutoStructs because comparing the AutoStructs would
	//compare the pointers, not the contents
	if decoded.ToString() != original.ToString() {
		t.Errorf("Interface fails with B. %s != %s", original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeC_1(t *testing.T) {
	var original C
	original.Uint64Member = 0xFFFFFFFFFFFFFFFF
	var decoded C
	decoded.Decode(original.Encode())

	if decoded != original {
		t.Errorf("Issue with encoding/decoding C. %X != %X", original.Uint64Member, decoded.Uint64Member)
	}
}

func TestEncodeDecodeC_2(t *testing.T) {
	var original C
	original.Uint64Member = 0x739
	var decoded C
	decoded.Decode(original.Encode())

	if decoded != original {
		t.Errorf("Issue with encoding/decoding C. %X != %X", original.Uint64Member, decoded.Uint64Member)
	}
}

func TestEncodeDecodeC_3(t *testing.T) {
	var original C
	original.Uint64Member = 0xFFFFF45FFABCFFFF
	var decoded C
	decoded.Decode(original.Encode())

	if decoded != original {
		t.Errorf("Issue with encoding/decoding C. %X != %X", original.Uint64Member, decoded.Uint64Member)
	}
}

func TestEncodeDecodeC_4(t *testing.T) {
	var original C
	original.Uint64Member = 0x0
	var decoded C
	decoded.Decode(original.Encode())

	if decoded != original {
		t.Errorf("Issue with encoding/decoding C. %X != %X", original.Uint64Member, decoded.Uint64Member)
	}
}

func TestInterfaceC(t *testing.T) {
	var originalStruct C
	var decodedStruct C
	originalStruct.Uint64Member = 0xFF77
	var original AutoStruct = &originalStruct
	var decoded AutoStruct = &decodedStruct
	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with C interface. %s != %s", original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeD_1(t * testing.T) {
	var original D
	var decoded D
	sliceLength := 7
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	original.SliceMember = slice

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encode/decode D. %s \n is not the same as \n %s", 
				original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeD_2(t * testing.T) {
	var original D
	var decoded D
	sliceLength := 11
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i*i)
	}
	original.SliceMember = slice

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encode/decode D. %s \n is not the same as \n %s", 
				original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeD_3(t * testing.T) {
	var original D
	var decoded D
	sliceLength := 0
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	original.SliceMember = slice

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encode/decode D. %s \n is not the same as \n %s", 
				original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeD_4(t * testing.T) {
	var original D
	var decoded D
	sliceLength := 5
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i*i*i)
	}
	original.SliceMember = slice

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encode/decode D. %s \n is not the same as \n %s", 
				original.ToString(), decoded.ToString())
	}
}

func TestInterfaceD(t *testing.T) {
	var originalStruct D
	var decodedStruct D
	sliceLength := 4
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	originalStruct.SliceMember = slice

	var original AutoStruct = &originalStruct
	var decoded AutoStruct = & decodedStruct
	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with D interface. %s \nis not the same as \n%s", original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeE_1(t *testing.T) {
	var original E
	var decoded E

	sliceLength := 8
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	bMemb := B{9}
	cMemb := C{0xF7734AD}
	dMemb := D{slice}

	original.ByteMember = 8
	original.Uint64Member = 0xFFF8
	original.SliceMember = slice
	original.Bstruct = bMemb
	original.Cstruct = cMemb
	original.Dstruct = dMemb

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encoding/decoding type E.\n%s\nIs not the same as\n%s",
					original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeE_2(t *testing.T) {
	var original E
	var decoded E

	sliceLength := 11
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	bMemb := B{99}
	cMemb := C{0x1F7734AD}
	dMemb := D{slice}

	original.ByteMember = 88
	original.Uint64Member = 0xFFF8FF8
	original.SliceMember = slice
	original.Bstruct = bMemb
	original.Cstruct = cMemb
	original.Dstruct = dMemb

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encoding/decoding type E.\n%s\nIs not the same as\n%s",
					original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeE_3(t *testing.T) {
	var original E
	var decoded E

	sliceLength := 3
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i*i)
	}
	bMemb := B{0xFF}
	cMemb := C{0xFB744434AD}
	dMemb := D{slice}

	original.ByteMember = 0xFF
	original.Uint64Member = 0xFFFFFFFFFFFFFFFF
	original.SliceMember = slice
	original.Bstruct = bMemb
	original.Cstruct = cMemb
	original.Dstruct = dMemb

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encoding/decoding type E.\n%s\nIs not the same as\n%s",
					original.ToString(), decoded.ToString())
	}
}

func TestEncodeDecodeE_4(t *testing.T) {
	var original E
	var decoded E

	sliceLength := 12
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i)
	}
	bMemb := B{38}
	cMemb := C{0xFFFFFFFFFFFFFFFF}
	dMemb := D{slice}

	original.ByteMember = 0
	original.Uint64Member = 0x0
	original.SliceMember = slice
	original.Bstruct = bMemb
	original.Cstruct = cMemb
	original.Dstruct = dMemb

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encoding/decoding type E.\n%s\nIs not the same as\n%s",
					original.ToString(), decoded.ToString())
	}
}

func TestInterfaceE(t *testing.T) {
	var originalStruct E
	var decodedStruct E

	sliceLength := 6
	slice := make([]byte, sliceLength)
	for i := 0; i < sliceLength; i++ {
		slice[i] = byte(i*i*i)
	}
	bMemb := B{138}
	cMemb := C{0xFFFFFFFFFFFFF}
	dMemb := D{slice}

	originalStruct.ByteMember = 48
	originalStruct.Uint64Member = 0x096
	originalStruct.SliceMember = slice
	originalStruct.Bstruct = bMemb
	originalStruct.Cstruct = cMemb
	originalStruct.Dstruct = dMemb

	var original AutoStruct = &originalStruct
	var decoded AutoStruct = &decodedStruct

	decoded.Decode(original.Encode())

	if original.ToString() != decoded.ToString() {
		t.Errorf("Issue with encoding/decoding type E.\n%s\nIs not the same as\n%s",
					original.ToString(), decoded.ToString())
	}
}