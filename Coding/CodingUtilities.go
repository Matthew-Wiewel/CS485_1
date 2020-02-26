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
func encodeTypeId(id TypeId) (buffer []byte) {
	//TODO
	//idCast := uint32(id)
	//_ = binary.Write(buffer, binary.BigEndian, idCast)
	return buffer
}

//helper function for encoding bytes
func encodeByte(b byte) (slice []byte) {
	//TODO
	return slice
}

//helper function for encoding uint64s
func encodeUint64(u uint64) (slice []byte) {
	//TODO
	return slice
}

//helper function for encoding byte slices
func encodeByteSlice(s []byte) (slice []byte) {
	//TODO
	return slice
}

//helper function to decode a TypeId from a byte slice
func decodeType(slice []byte) (id TypeId) {
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
