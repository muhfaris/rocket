package builder

func availableConfig() []string {
	return []string{"toml", "yaml", "env"}
}

func findConfig(c string) bool {
	for _, config := range availableConfig() {
		if config == c {
			return true
		}
	}

	return false
}
