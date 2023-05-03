package kit

import (
	"docman/pkg/database"
)

func ExecuteInTransaction(fn func(args ...any) error) error {
	tx := database.Inst.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := fn()
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
