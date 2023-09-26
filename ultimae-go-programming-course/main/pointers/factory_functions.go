package main

type user struct {
	name  string
	email string
}

func main() {
	u := createUserV2()
	println("user v2 = ", u.name, u.email)
}

// createUserV1 creates a user value and passed a copy back to the caller. This is value semantics
func createUserV1() user {
	u := user {
		name: "Bill value seamantics",
		email: "abc@gmail.com",
	}
	println("V1", &u)
	return u
}

// createUserV2  creates a user value and passes address of the user. This is pointer semantics
func createUserV2() *user {
	// this u is converted underneath into a pointer to be able to access the heap value
	u := user {
		name: "Bill pointer semantics",
		email: "abc@gmail.com",
	}
	println("V1", &u)
	return &u
}

// createUserV3  creates a user value using pointer semantics construction and this is a complex and clever code
func createUserV3() *user {
	// we are making code much harder to read and we are really also mixing the semantics. We are walking away from
	// readability. The only time we do this is if we are directly returning this or we are doing it inside a function
	// call as shown below
	//return &user{
	//	name:  "Bill pointer semantics",
	//	email: "abc@gmail.com",
	//}
	//foo(&user{
	//	name:  "Bill pointer semantics",
	//	email: "abc@gmail.com",
	//})
	u := &user{
		name:  "Bill pointer semantics",
		email: "abc@gmail.com",
	}

	println("V1", u)
	return u
}