package coding

//TypeId is used to differentiate difference kinds of structs that implement
//the AutoStruct interface
type TypeId uint32

//AutoStruct is an interface that will be used to allow
//easy encoding and decoding of structs that implement it
type AutoStruct interface {
	Encode() []byte
	Decode([]byte)
	getTypeId() TypeId
}