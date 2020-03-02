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

	fmt.Println("The struct we wrote was: ", obj.ToString())
}

func TestRead_1(t *testing.T) {
	AutofillTypeMap()
	obj, err := Read("data1.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The resulting struct is: ", obj.ToString())
}

func TestWrite_2(t *testing.T) {
	obj := B{124}
	err := Write("data2.txt", &obj)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The struct we wrote was: ", obj.ToString())
}

func TestRead_2(t *testing.T) {
	AutofillTypeMap()
	obj, err := Read("data2.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The resulting struct is: ", obj.ToString())
}

func TestWrite_3(t *testing.T) {
	slice := make([]byte, 12)
	for i := 0; i < 12; i++ {
		slice[i] = byte(4*i)
	}
	obj := D{slice}
	err := Write("data3.txt", &obj)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The struct we wrote was: ", obj.ToString())
}

func TestRead_3(t *testing.T) {
	AutofillTypeMap()
	obj, err := Read("data3.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The resulting struct is: ", obj.ToString())
}

func TestWrite_4(t *testing.T) {
	slice1 := make([]byte, 5)
	slice2 := make([]byte, 7)
	for i := 0; i < 5; i++ {
		slice1[i] = byte(i)
	}
	for i := 0; i < 7; i++ {
		slice2[i] = byte(i*i)
	}
	obj := E{51, 0xFFEECBBA9988, slice1, B{32}, C{0x456}, D{slice2}}
	err := Write("data4.txt", &obj)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The struct we wrote was: ", obj.ToString())
}

func TestRead_4(t *testing.T) {
	AutofillTypeMap()
	obj, err := Read("data4.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The resulting struct is: ", obj.ToString())
}