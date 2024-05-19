package myGoProjects

import (
	"fmt"
	"myProject.io/pkg/db"
)

func getFromDBAndHandleIt(myDB db.DB, input string) (string, error) {
	result, err := myDB.Get(input)
	if err != nil {
		return "", err
	}

	// Start to handle the result from DB
	fmt.Println(result)

	return result, nil
}
