package type_assertions

import "fmt"

func Run() {
	var s interface{} = "test"

	if s2, ok := s.(string); ok {
		fmt.Printf("To string was ok. Val: %s \n", s2)
	} else {
		fmt.Println("To string wasn't ok.")
	}

	if f1, ok := s.(float32); ok {
		fmt.Printf("To float was ok. Val: %f \n", f1)
	} else {
		fmt.Println("To float wasn't ok.")
	}

	// -------

	switch s.(type) {
	case int:
		fmt.Println("Type is int")
		break
	case string:
		fmt.Println("Type is int")
		break
	case float32:
		fmt.Println("Type is float")
		break
	case bool:
		fmt.Println("Type is bool")
		break
	default:
		fmt.Println("Type is unknown")
	}

}
