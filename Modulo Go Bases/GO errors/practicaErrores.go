package main

import (
	"fmt"
	"os"
)

/*func main() {

	err1 := fmt.Errorf("first error :(")

	err2 := fmt.Errorf("second erro: %w", err1)

	fmt.Println(err1)
	fmt.Println(err2)
}*/

type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf("%s (%d)",
		err.Message,
		err.StatusCode,
	)
}
func ObtainError(status int) (code int, err error) {
	if status > 400 {
		code = 500
		err = &MyCustomError{
			StatusCode: code,
		}
	}
	return 200, nil

}
func main() {

	status, err := ObtainError(300)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d funciona! :D", status)
}
