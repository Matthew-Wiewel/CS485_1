package coding

//Encode method for struct B
func (caller *B) Encode() (encoding []byte) {
	//TODO
	//encoding = append(encoding, encodeTypeId(bID))
	//encoding = append(encoding, encodeByte(b.B))
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
	//TODO
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
	//TODO
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
	//TODO
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