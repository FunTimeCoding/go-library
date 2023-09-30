package git

func Push() {
	Run("push", "origin", "--tags")
}
