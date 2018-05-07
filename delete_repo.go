
import (
	"database/sql"
	"fmt"
)

// START OMIT
func Delete(db *sql.DB, theStream *Stream) error {
	stmt, err := db.Prepare("DELETE FROM streams WHERE path = ?")
	if err != nil {
		return fmt.Errorf("could not prepare DELETE statement: %v", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(theStream.Path)
	if err != nil {
		return fmt.Errorf("could not execute DELETE statement: %v", err)
	}

	return nil
}

// END OMIT
