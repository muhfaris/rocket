package builder

func availableCache() []string {
	return []string{"redis"}
}

func findCache(c string) bool {
	for _, cache := range availableCache() {
		if cache == c {
			return true
		}
	}

	return false
}
