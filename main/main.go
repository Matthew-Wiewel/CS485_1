package main

import "coding"
import "fmt"

func main() {
	
	//fills the underlying typeMap
	coding.AutofillTypeMap()

	//create structs that will be encoded and written to a file
	var bStruct coding.B = coding.B{42}
	var bStruct2 coding.B = coding.B{163}
	var cStruct coding.C = coding.C{0xFFEEDD}
	slice1 := []byte("This is an example slice value")
	var dStruct coding.D = coding.D{slice1}
	slice2 := []byte("This is an second example of a slice value")
	var eStruct coding.E = coding.E{8, 0x1234, slice2, bStruct, cStruct, dStruct}

	//print out what those structs we're writing to a file look like
	fmt.Println("Here are the structs we wrote \n\nB struct: ", bStruct.ToString(),
				"\n\nB struct #2: ", bStruct2.ToString(),
				"\n\nC struct: ", cStruct.ToString(),
				"\n\nD struct: ", dStruct.ToString(),
				"\n\nE struct: ", eStruct.ToString())

	//write structs to a file
	coding.Write("bData.txt", &bStruct)
	coding.Write("bData2.txt", &bStruct2)
	coding.Write("cData.txt", &cStruct)
	coding.Write("dData.txt", &dStruct)
	coding.Write("eData.txt", &eStruct)
	
	//read from the respective files and decode the results
	//these are now AutoStructs we have
	bAuto, _ := coding.Read("bData.txt")
	bAuto2, _ := coding.Read("bData2.txt")
	cAuto, _ := coding.Read("cData.txt")
	dAuto, _ := coding.Read("dData.txt")
	eAuto, _ := coding.Read("eData.txt")

	//print out what structs were returned
	fmt.Println("Here are the structs we read in. \n\nB struct: ", bAuto.ToString(),
				"\n\nB struct #2: ", bAuto2.ToString(),
				"\n\nC struct: ", cAuto.ToString(),
				"\n\nD struct: ", dAuto.ToString(),
				"\n\nE struct: ", eAuto.ToString())

}