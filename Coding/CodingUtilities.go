package coding

import (
	"encoding/binary"
)

//typeMap is a private package member used to correlate typeIds
//to function that generate that type
var typeMap = make(map[TypeId]func() AutoStruct)

//SetType is a method to allow entries to be added to the typeMap
func SetType(id TypeId, f func() AutoStruct) {
	typeMap[id] = f
}

//Read is a function to read in a byte stream and create the appropriate AutoStruct
//TODO: see if erro os.Error will work on the linux
func Read(path string) (t AutoStruct, err error) {
	//TODO
	return t, err
}

//Write is a function to take an AutoStruct and write it to a byte steam
//TODO: see if erro os.Error will work on the linux
func Write(path string, t AutoStruct) (err error) {
	//TODO
	return err
}

//AutofillTypeMap is a method to be called that can fill the typeMap with
//the appropriate TypeId to function mapping at runtime
func AutofillTypeMap() {
	typeMap[bID] = makeB
	typeMap[cID] = makeC
	typeMap[dID] = makeD
	typeMap[eID] = makeE
}

//typeFactor takes a TypeId, looks up the function set in the typeMap
//and uses that function to create the appropriate result, assigning it
//to AutoStruct 'a' and then returning 'a'
func typeFactory(id TypeId) (a AutoStruct) {
	function := typeMap[id]
	a = function()
	return a
}

//helper function for encoding TypeIds
func encodeTypeId(id TypeId) []byte {
	idCast := uint32(id) //cast the TypeId to uint32 to allow use of standard library
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, idCast)
	return buffer
}

//helper function for encoding bytes
func encodeByte(b byte) []byte {
	buffer := make([]byte, 1)
	buffer[0] = b
	return buffer
}

//helper function for encoding uint64s
func encodeUint64(u uint64) []byte {
	buffer := make([]byte, 8) //uint64 has 8 bytes, hence a byte slice of size 8
	binary.BigEndian.PutUint64(buffer, u) //write u as bytes
	return buffer
}

//helper function for encoding byte slices
func encodeByteSlice(s []byte) []byte {
	buffer := make([]byte, len(s))
	//TODO
	return buffer
}

//helper function to decode a TypeId from a byte slice
func decodeTypeId(slice []byte) (id TypeId) {
	var idAsUint32 = binary.BigEndian.Uint32(slice)
	id = TypeId(idAsUint32)
	return id
}

//helper function to decode a byte from a byte slice
func decodeByte(slice []byte) (b byte) {
	b = slice[0]
	return b
}

//helepr function to decode a uint64 from a byte slice
func decodeUint64(slice []byte) (u uint64) {
	u = binary.BigEndian.Uint64(slice)
	return u
}

//helper function to decode a byte slice from a byte slice
func decodeByteSlice(slice []byte) (s []byte) {
	s = slice //TODO: check if this gets intended result
	return s
}

//helper to append a slice to another
func appendEncodings(oldBuffer, sliceToAppend []byte) (newBuffer []byte) {
	//TODO
	return newBuffer
}
