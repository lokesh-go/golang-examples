package solid

import "fmt"

// Subtypes should be substitutable for their base types without altering the correctness of the program.
// A derived class (subtype) must be able to replace its parent class (supertype) without breaking the program.
// The behavior of the child class should be consistent with the behavior expected from the parent class

// Problem -------------------
// SuperType structure
type user struct {
	id   string
	name string
}

func (u *user) getData() {
	fmt.Println(u.id, u.name)
}

// SubType
type vipUser struct {
	id    string
	name  string
	level string
}

func (v *vipUser) getData() {
	fmt.Println(v.id, v.name, v.level)
}

func printUserDetails(u *user) {
	u.getData()
}

func calls() {
	// Calls for SuperType struct
	u := &user{id: "1", name: "superUser"}
	printUserDetails(u)

	// Calls for SubType struct
	//vu := &vipUser{id: "2", name: "vipUser"}
	// Below statement will through error
	// A derived class (SubType) will not able to replace its parent class - Not following the LiskovSubstitution principle
	//printUserDetails(vu)
}

// Solution ----------------------

// SuperType
type userInterface interface {
	getData()
}

func printUserDetail(u userInterface) {
	u.getData()
}

func callss() {
	// Calls for user
	u := &user{id: "1", name: "user"}
	printUserDetails(u)

	// Calls for vip user
	vu := &vipUser{id: "1", name: "vipUser", level: "DOG"}
	printUserDetail(vu)
}

// Both user & vipUser able to call the printUserDetail function and functionality is not breaking
