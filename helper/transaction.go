package helper

import (
	"fmt"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		fmt.Println(errorRollback)
	} else {
		errorCommit := tx.Commit().Error
		fmt.Println(errorCommit)
	}
}
