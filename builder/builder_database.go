package builder

import "github.com/muhfaris/rocket/shared/constanta"

func availableDatabase() []string {
	return []string{
		constanta.DBMySQL,
		constanta.DBPostgres,
		constanta.DBSQLite,
		constanta.DBMongo,
	}
}

func findDatabase(db string) bool {
	for _, database := range availableDatabase() {
		if database == db {
			return true
		}
	}

	return false
}
