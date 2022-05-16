package transaction

import (
	"go-restful-api-lamp/helper"

	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		helper.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit().Error
		helper.PanicIfError(errorCommit)

		panic(err)
	}
}
