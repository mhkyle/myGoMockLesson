package mysql

import "fmt"

type MySQL struct {
}

func (m *MySQL) Get(input string) (string, error) {
	return fmt.Sprintf("result: %s", input), nil
}
