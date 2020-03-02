package coding

import (
	"fmt"
	"log"
	"testing"
)

func TestWrite_1(t *testing.T) {
	obj := C{0xFFEEDDCCBBAA9988}
	err := Write("data1.txt", &obj)

	if err != nil {
		log.Fatal(err)
	}
}

func TestRead_1(t *testing.T) {
	obj, err := Read("data1.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The resulting struct is: ", obj)
}
