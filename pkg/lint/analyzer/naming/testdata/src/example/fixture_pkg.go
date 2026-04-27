package example

type keywordGuard struct {
	pkg string // want `use "package" instead of "pkg" in pkg`
}

func pkg() {} // want `use "package" instead of "pkg" in pkg`
