package coding

//Encode method for struct B
func (caller *B) Encode() (encoding []byte) {
	encoding = appendEncodings(encoding, encodeTypeId(bID))
	encoding = appendEncodings(encoding, encodeByte(caller.ByteMember))
	return encoding
}

//Decode method for struct B
func (caller *B) Decode(encoding []byte) {
	//TODO
}

//getTypeId method for struct B
func (caller *B) getTypeId() TypeId {
	return bID
}

//Encode method for struct C
func (caller *C) Encode() (encoding []byte) {
	encoding = appendEncodings(encoding, encodeTypeId(cID))
	encoding = appendEncodings(encoding, encodeUint64(caller.Uint64Member))
	return encoding
}

//Decode method for struct C
func (caller *C) Decode(encoding []byte) {
	//TODO
}

//getTypeId method for struct C
func (caller *C) getTypeId() TypeId {
	return cID
}

//Encode method for struct D
func (caller *D) Encode() (encoding []byte) {
	encoding = appendEncodings(encoding, encodeTypeId(dID))
	encoding = appendEncodings(encoding, encodeByteSlice(caller.SliceMember))
	return encoding
}

//Decode method for struct D
func (caller *D) Decode(encoding []byte) {
	//TODO
}

//getTypeId method for struct D
func (caller *D) getTypeId() TypeId {
	return dID
}

//Encode method for struct E
func (caller *E) Encode() (encoding []byte) {
	encoding = appendEncodings(encoding, encodeByte(caller.ByteMember))
	encoding = appendEncodings(encoding, encodeUint64(caller.Uint64Member))
	encoding = appendEncodings(encoding, encodeByteSlice(caller.SliceMember))
	encoding = appendEncodings(encoding, caller.Bstruct.Encode())
	encoding = appendEncodings(encoding, caller.Cstruct.Encode())
	encoding = appendEncodings(encoding, caller.Dstruct.Encode())
	return encoding
}

//Decode method for struct E
func (caller *E) Decode(encoding []byte) {
	//TODO
}

//getTypeId method for struct E
func (caller *E) getTypeId() TypeId {
	return eID
}