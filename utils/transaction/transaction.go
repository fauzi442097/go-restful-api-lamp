package transaction

import (
	"fmt"
	"go-restful-api-lamp/helper"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		fmt.Println("rollback")
		errorRollback := tx.Rollback().Error
		helper.PanicIfError(errorRollback)
		panic(err)
	} else {
		fmt.Println("commit")
		errorCommit := tx.Commit().Error
		helper.PanicIfError(errorCommit)
	}
}
