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
