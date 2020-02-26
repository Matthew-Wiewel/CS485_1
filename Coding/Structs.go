package coding

//constants for the types of each struct
const (
	bID TypeId = iota
	cID TypeId = iota
	dID TypeId = iota
	eID TypeId = iota
)

//methods to create each type of struct

func makeB() (b B) {
	return b
}
func makeC() (c C) {
	return c
}
func makeD() (d D) {
	return d
}
func makeE() (e E) {
	return e
}

//note: there is no struct A due to some annoyance on my part at the linter

//B is a struct of AutoStruct
type B struct {
	ByteMember byte
}

//C is a struct of AutoStruct
type C struct {
	Uint64Member uint64
}

//D is a struct of Austrocut
type D struct {
	SliceMember []byte
}

//E is a struct of AutoStruct
type E struct {
	ByteMember byte
	Uint64Member uint64
	SliceMember []byte
	Bstruct B
	Cstruct C
	Dstruct D
}