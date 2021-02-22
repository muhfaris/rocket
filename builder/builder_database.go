package builder

func availableDatabase() []string {
	return []string{"postgresql", "mysql", "mongodb"}
}

func findDatabase(db string) bool {
	for _, database := range availableDatabase() {
		if database == db {
			return true
		}
	}

	return false
}
