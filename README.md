# CS485_1

name: Matthew Wiewel
uin: see email
assignment: 1

As a note, I have not tested this project on Ethos. It has been tested on Windows.

In the coding folder is the code of the coding package.

AutoStruct.go contains the definition of the AutoStruct interface and TypeId
As a note, I added a ToString() string method to the AutoStruct interface. 
This was done to allow ease of printing for evaluation and allowed me to compare
two AutoStructs for equality as my initial tests ended up testing the pointers
to the underlying structs instead of the actual values of those structs.

CodingUtilities.go contains the Read, Write, SetType, and typeFactory functions.
It also contains various helper functions for decoding/encoding certain datatypes.
Those function have names of encodeType/decodeType. 
The AutoFillTypeMap function was created by me for ease of use when it came to 
putting in values for the TypeIds of the structs I defined.
appendEncoding was another helper function I defined. It takes a variadic parameter
of byte slices and combines them all into one byte slice. This was implemented for ease 
of combining the various byte slices created by my helper encode methods.

Structs.go defines the constants of TypeIds associated with each struct. It also 
defines make functions for each type of struct that return an AutoStruct with that
kind of struct as the underlying type. This file also contains the defintions of each struct.

StructMethods.go contains the implementation of each struct's methods. Each struct implements
getTypeId(), Encode(), Decode(), and ToString(). 

coding_test.go have some methods that helped me in testing Read and Write. There are 
not any tests of equality inside these methods. Instead testing was done via looking at
the outputs printed to the terminal from each method.

helper_test.go have unit tests for each helper method of the CodingUtilities.go folder

methods_test.go tests the required methods of each struct class. It does not test ToString()

In the main folder is a sample program to test the package and its usage. In main.go a few structs
are created, printed, written to a file, read from a file, and the decoded structs are also
printed to the terminal. The makefile is simplistic due to the lack of ethos specificity.