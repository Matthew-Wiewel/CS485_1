package coding

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

//Encode method for struct E
func (caller *E) Encode() (encoding []byte) {
	encoding = appendEncodings(encodeTypeId(eID), encodeByte(caller.ByteMember),
				encodeUint64(caller.Uint64Member), encodeByteSlice(caller.SliceMember),
				caller.Bstruct.Encode(), caller.Cstruct.Encode(), caller.Dstruct.Encode())
	return encoding
}

//Decode method for struct E
func (caller *E) Decode(encoding []byte) {
	//TODO
	caller.ByteMember = decodeByte(encoding[4:8])
	caller.Uint64Member = decodeUint64(encoding[8:16])
	caller.SliceMember = decodeByteSlice(encoding[16:]) //TODO, figure out cut-offs for the rest
	//TODO with the cutoff
	caller.Bstruct.Decode(encoding[:])
	caller.Cstruct.Decode(encoding[:])
	caller.Dstruct.Decode(encoding[:])
}

//getTypeId method for struct E
func (caller *E) getTypeId() TypeId {
	return eID
}
