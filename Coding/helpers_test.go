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