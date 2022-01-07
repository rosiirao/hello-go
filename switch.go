package main

import "fmt"

func init() {
	fmt.Println(" select init ")
}

// PrintType print the type of param
func PrintType(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	case float32, float64:
		fmt.Printf("pointer to float %f\n", t)
	}
}
