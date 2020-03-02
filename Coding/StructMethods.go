package coding

import "fmt"

//Encode method for struct B
func (caller *B) Encode() (encoding []byte) {
	encoding = appendEncodings(encodeTypeId(bID), encodeByte(caller.ByteMember))
	return encoding
}

//Decode method for struct B
func (caller *B) Decode(encoding []byte) {
	caller.ByteMember = decodeByte(encoding[4:]) //first 4 bytes are the TypeId
	//so use what's left to get the byte
}

//getTypeId method for struct B
func (caller *B) getTypeId() TypeId {
	return bID
}

//ToString method of B
func (caller *B) ToString() string {
	return fmt.Sprintf("%d", caller.ByteMember)
}

//Encode method for struct C
func (caller *C) Encode() (encoding []byte) {
	encoding = appendEncodings(encodeTypeId(cID), encodeUint64(caller.Uint64Member))
	return encoding
}

//Decode method for struct C
func (caller *C) Decode(encoding []byte) {
	caller.Uint64Member = decodeUint64(encoding[4:])
}

//getTypeId method for struct C
func (caller *C) getTypeId() TypeId {
	return cID
}

//ToString method of C
func (caller *C) ToString() string {
	return fmt.Sprintf("%X", caller.Uint64Member)
}

//Encode method for struct D
func (caller *D) Encode() (encoding []byte) {
	encoding = appendEncodings(encodeTypeId(dID), encodeByteSlice(caller.SliceMember))
	return encoding
}

//Decode method for struct D
func (caller *D) Decode(encoding []byte) {
	caller.SliceMember = decodeByteSlice(encoding[4:])
}

//getTypeId method for struct D
func (caller *D) getTypeId() TypeId {
	return dID
}

//ToString method of D
func (caller *D) ToString() string {
	return byteSliceToString(caller.SliceMember)
}

//Encode method for struct E
func (caller *E) Encode() (encoding []byte) {
	encoding = appendEncodings(encodeTypeId(eID), encodeByte(caller.ByteMember),
		encodeUint64(caller.Uint64Member), encodeByteSlice(caller.SliceMember),
		caller.Bstruct.Encode(), caller.Cstruct.Encode(), caller.Dstruct.Encode())
	return encoding
}

//Decode method for struct E
func (caller *E) Decode(encoding []byte) {
	//decode the byte and uint64 members as we know these are the indices the locations
	//will be seen for sure
	caller.ByteMember = decodeByte(encoding[4:5])
	caller.Uint64Member = decodeUint64(encoding[5:13])

	//figure out how long the slice member is and decode from there
	sliceEnd := decodeUint64(encoding[13:21]) + 21
	caller.SliceMember = decodeByteSlice(encoding[13:sliceEnd]) //include the uint64 because the decode still expects it

	//using the end of that slice as the new starting point, decode the structs
	caller.Bstruct.Decode(encoding[sliceEnd : sliceEnd+5])    //4 for typeId + 1 for byte member
	caller.Cstruct.Decode(encoding[sliceEnd+5 : sliceEnd+17]) //4 for type, 8 for uint62
	caller.Dstruct.Decode(encoding[sliceEnd+17:])             //end of C struct to the end for D's slice member
}

//getTypeId method for struct E
func (caller *E) getTypeId() TypeId {
	return eID
}

//ToString method of E
func (caller *E) ToString() string {
	return fmt.Sprintf("%d, %x \n%s, \nB: %s \nC: %s, \nD: %s", caller.ByteMember,
		caller.Uint64Member, byteSliceToString(caller.SliceMember),
		caller.Bstruct.ToString(), caller.Cstruct.ToString(),
		caller.Dstruct.ToString())
}

//helper function for converting slices to strings for the purposes of this program
func byteSliceToString(slice []byte) string {
	var sliceString string
	sliceString += fmt.Sprintf("Length: %d\nContents: ", len(slice))
	for _, b := range slice {
		sliceString += fmt.Sprintf("%d ", b)
	}
	return sliceString
}
