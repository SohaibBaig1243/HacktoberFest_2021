package database

import (
	"context"
	"database/sql"
)

type Database struct {
	SqlDb *sql.DB
}
/*here is the comment*/
var dbContext = context.Background()
