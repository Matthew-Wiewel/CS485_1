package coding

import "testing"

func TestByteEncoding_1(t *testing.T) {
	var byteToEncode byte = 0
	encoding := encodeByte(byteToEncode)
	decodedByte := decodeByte(encoding)

	if byteToEncode != decodedByte {
		t.Errorf("Encoded byte %b was decoded as %b", int(byteToEncode), int(decodedByte))
	}
}

func TestByteEncoding_2(t *testing.T) {
	var byteToEncode byte = 127
	encoding := encodeByte(byteToEncode)
	decodedByte := decodeByte(encoding)

	if byteToEncode != decodedByte {
		t.Errorf("Encoded byte %b was decoded as %b", int(byteToEncode), int(decodedByte))
	}
}

func TestByteEncoding_3(t *testing.T) {
	var byteToEncode byte = 75
	encoding := encodeByte(byteToEncode)
	decodedByte := decodeByte(encoding)

	if byteToEncode != decodedByte {
		t.Errorf("Encoded byte %b was decoded as %b", int(byteToEncode), int(decodedByte))
	}
}

func TestByteEncoding_4(t *testing.T) {
	var byteToEncode byte = 92
	encoding := encodeByte(byteToEncode)
	decodedByte := decodeByte(encoding)

	if byteToEncode != decodedByte {
		t.Errorf("Encoded byte %b was decoded as %b", int(byteToEncode), int(decodedByte))
	}
}

func TestUint64Encoding_1(t *testing.T) {
	var uint64ToEncode uint64 = 0
	encoding := encodeUint64(uint64ToEncode)
	decoded := decodeUint64(encoding)

	if uint64ToEncode != decoded {
		t.Errorf("Encoded uint64 %d was decoded as %d", uint64ToEncode, decoded)
	}
}

func TestUint64Encoding_2(t *testing.T) {
	var uint64ToEncode uint64 = 0xFFFFFFFFFFFFFFFF
	encoding := encodeUint64(uint64ToEncode)
	decoded := decodeUint64(encoding)

	if uint64ToEncode != decoded {
		t.Errorf("Encoded uint64 %d was decoded as %d", uint64ToEncode, decoded)
	}
}

func TestUint64Encoding_3(t *testing.T) {
	var uint64ToEncode uint64 = 98468
	encoding := encodeUint64(uint64ToEncode)
	decoded := decodeUint64(encoding)

	if uint64ToEncode != decoded {
		t.Errorf("Encoded uint64 %d was decoded as %d", uint64ToEncode, decoded)
	}
}

func TestUint64Encoding_4(t *testing.T) {
	var uint64ToEncode uint64 = 36375
	encoding := encodeUint64(uint64ToEncode)
	decoded := decodeUint64(encoding)

	if uint64ToEncode != decoded {
		t.Errorf("Encoded uint64 %d was decoded as %d", uint64ToEncode, decoded)
	}
}

func TestEncodeTypeId_1(t *testing.T) {
	encode := bID
	encoding := encodeTypeId(encode)
	decode := decodeTypeId(encoding)

	if encode != decode {
		t.Errorf("TypeId %d was decoded as %d", encode, decode)
	}
}

func TestEncodeTypeId_2(t *testing.T) {
	encode := cID
	encoding := encodeTypeId(encode)
	decode := decodeTypeId(encoding)

	if encode != decode {
		t.Errorf("TypeId %d was decoded as %d", encode, decode)
	}
}

func TestEncodeTypeId_3(t *testing.T) {
	encode := dID
	encoding := encodeTypeId(encode)
	decode := decodeTypeId(encoding)

	if encode != decode {
		t.Errorf("TypeId %d was decoded as %d", encode, decode)
	}
}

func TestEncodeTypeId_4(t *testing.T) {
	encode := eID
	encoding := encodeTypeId(encode)
	decode := decodeTypeId(encoding)

	if encode != decode {
		t.Errorf("TypeId %d was decoded as %d", encode, decode)
	}
}

func TestEncodeByteSlice_1(t *testing.T) {
	sliceLength := 8
	encode := make([]byte, sliceLength)
	
	for i := 0; i < sliceLength; i++ {
		encode[i] = byte(i*i)
	}

	encoding := encodeByteSlice(encode)
	decode := decodeByteSlice(encoding)

	if len(encode) != len(decode) {
		t.Errorf("Lenght of byte slice {%d} was changed to {%d} when decoded", len(encode), len(decode))
	}

	for i := 0; i < sliceLength; i++ {
		if encode[i] != decode[i] {
			t.Errorf("Index %d of encoded and decoded byte slices is not equal\n" +
		"encode[%d] = %d and decode[%d] = %d", i, i, int(encode[i]), i, int(decode[i]))
		}
	}
}

func TestEncodeByteSlice_2(t *testing.T) {
	sliceLength := 13
	encode := make([]byte, sliceLength)
	
	for i := 0; i < sliceLength; i++ {
		encode[i] = 0xFF
	}

	encoding := encodeByteSlice(encode)
	decode := decodeByteSlice(encoding)

	if len(encode) != len(decode) {
		t.Errorf("Lenght of byte slice {%d} was changed to {%d} when decoded", len(encode), len(decode))
	}

	for i := 0; i < sliceLength; i++ {
		if encode[i] != decode[i] {
			t.Errorf("Index %d of encoded and decoded byte slices is not equal\n" +
		"encode[%d] = %d and decode[%d] = %d", i, i, int(encode[i]), i, int(decode[i]))
		}
	}
}

func TestEncodeByteSlice_3(t *testing.T) {
	sliceLength := 5
	encode := make([]byte, sliceLength)
	
	for i := 0; i < sliceLength; i++ {
		encode[i] = 0
	}

	encoding := encodeByteSlice(encode)
	decode := decodeByteSlice(encoding)

	if len(encode) != len(decode) {
		t.Errorf("Lenght of byte slice {%d} was changed to {%d} when decoded", len(encode), len(decode))
	}

	for i := 0; i < sliceLength; i++ {
		if encode[i] != decode[i] {
			t.Errorf("Index %d of encoded and decoded byte slices is not equal\n" +
		"encode[%d] = %d and decode[%d] = %d", i, i, int(encode[i]), i, int(decode[i]))
		}
	}
}

func TestEncodeByteSlice_4(t *testing.T) {
	sliceLength := 8
	encode := make([]byte, sliceLength)
	
	for i := 0; i < sliceLength; i++ {
		encode[i] = byte(i+i)
	}

	encoding := encodeByteSlice(encode)
	decode := decodeByteSlice(encoding)

	if len(encode) != len(decode) {
		t.Errorf("Lenght of byte slice {%d} was changed to {%d} when decoded", len(encode), len(decode))
	}

	for i := 0; i < sliceLength; i++ {
		if encode[i] != decode[i] {
			t.Errorf("Index %d of encoded and decoded byte slices is not equal\n" +
		"encode[%d] = %d and decode[%d] = %d", i, i, int(encode[i]), i, int(decode[i]))
		}
	}
}

func TestAppendEncodings_1(t *testing.T) {
	slice1 := make([]byte, 4)
	slice2 := make([]byte, 8)

	for i := 0; i < 8; i++ {
		if i < 4 {
			slice1[i] = byte(i)
		}
		slice2[i] = byte(i+4)
	}

	expectedSlice := make([]byte, 12)
	for i := 0; i < 12; i++ {
		expectedSlice[i] = byte(i)
	}

	returnedSlice := appendEncodings(slice1, slice2)

	if len(expectedSlice) != len(returnedSlice) {
		t.Errorf("\nexpected size: %d, actual size: %d", len(expectedSlice), len(returnedSlice))
	}

	for i := 0; i < len(expectedSlice); i++ {
		if expectedSlice[i] != returnedSlice[i] {
			t.Errorf("Index %d mismatch. %d != %d.", i, expectedSlice[i], returnedSlice[i])
		}
	}
}