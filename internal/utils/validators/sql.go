package validators

import (
	"database/sql"
	"errors"

	"github.com/nuttchai/go-rest/internal/constant"
)

func CheckRowsAffected(sqlResult sql.Result) error {
	rowsAffected, err := (sqlResult).RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New(constant.SampleNotFound)
	}

	return nil
}
