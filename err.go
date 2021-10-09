package err

import (
	"errors"
	"fmt"
)

func Err(str string) error {
	fmt.Printf("error=%d", str)
	return errors.New(str)
}
