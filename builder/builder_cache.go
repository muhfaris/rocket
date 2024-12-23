package builder

import "github.com/muhfaris/rocket/shared/constanta"

func availableCache() []string {
	return []string{
		constanta.CacheRedis,
		constanta.CacheInMemory,
	}
}

func findCache(c string) bool {
	for _, cache := range availableCache() {
		if cache == c {
			return true
		}
	}

	return false
}
