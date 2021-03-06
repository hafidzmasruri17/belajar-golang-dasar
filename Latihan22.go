package main

import "fmt"

// TODO: Pointer

// TODO: Struct
type Address struct {
	Village, City, Country string
}

// TODO: Pointer Function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	/** pass by value
	& = pointer
	* = reference all
	*/
	address1 := Address{"jambekumbu", "lumajang", "indonesia"}
	address2 := &address1
	address3 := &address1

	// pass by reference
	address2.Village = "Isekai"

	*address2 = Address{"Konoha", "Malang", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//	TODO: New Function
	getAddress := new(Address)
	getAddress.Country = "Japan"
	fmt.Println(getAddress)

	var address = Address{
		Village: "Konoha",
		City:    "Hi No Ku Ni",
		Country: "",
	}

	var addressPointer *Address = &address
	fmt.Println(addressPointer)

}
