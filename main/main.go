package main

import "coding"
import "fmt"

func main() {
	
	coding.AutofillTypeMap()

	var bStruct coding.B = coding.B{42}
	var cStruct coding.C = coding.C{0xFFEEDD}
	slice1 := []byte("This is an example slice value")
	var dStruct coding.D = coding.D{slice1}
	slice2 := []byte("This is an second example of a slice value")
	var eStruct coding.E = coding.E{8, 0x1234, slice2, bStruct, cStruct, dStruct}

	coding.Write("bData.txt", &bStruct)
	coding.Write("cData.txt", &cStruct)
	coding.Write("dData.txt", &dStruct)
	coding.Write("eData.txt", &eStruct)

	fmt.Println("Here are the structs we wrote \nB struct: ", bStruct.ToString(),
				"\nC struct: ", cStruct.ToString(),
				"\nD struct: ", dStruct.ToString(),
				"\nE struct: ", eStruct.ToString())
	
	bAuto, _ := coding.Read("bData.txt")
	cAuto, _ := coding.Read("cData.txt")
	dAuto, _ := coding.Read("dData.txt")
	eAuto, _ := coding.Read("eData.txt")

	fmt.Println("Here are the structs we read in. \nB struct: ", bAuto.ToString(),
				"\nC struct: ", cAuto.ToString(),
				"\nD struct: ", dAuto.ToString(),
				"\nE struct: ", eAuto.ToString())

}