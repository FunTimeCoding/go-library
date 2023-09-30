package git

func Fetch() {
	Run("fetch", "--prune", "--prune-tags")
}
