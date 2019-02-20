package cat

type Cat struct {
	Name string
	Age  int
}

// Go does not have constructors. Simply use a function instead
func New() Cat {
	return Cat{
		Name: "Diana",
		Age:  12,
	}
}
