package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/vfs"
	"testing"
)

func TestLintVFSImportFix(t *testing.T) {
	v := vfs.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	fixes := LintVFS(v, option.New("", false), false)
	assert.String(
		t,
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
		fixes.Read("pkg/foo/foo.go"),
	)
}

func TestLintVFSCleanNoFixes(t *testing.T) {
	v := vfs.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	v.Write(
		"pkg/foo/foo_test.go",
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n\tt.Parallel()\n}\n",
	)
	fixes := LintVFS(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo_test.go"))
}

func TestLintVFSSkipsGeneratedFile(t *testing.T) {
	v := vfs.New()
	v.Write(
		"pkg/foo/generated.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Generated() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	fixes := LintVFS(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/generated.go"))
}

func TestLintVFSVariableFlaggedNoFix(t *testing.T) {
	v := vfs.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\terr := fmt.Errorf(\"oops\")\n\t_ = err\n}\n",
	)
	v.Write(
		"pkg/foo/foo_test.go",
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n}\n",
	)
	fixes := LintVFS(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo.go"))
}

func TestLintVFSMultipleFilesOnlyBrokenFixed(t *testing.T) {
	v := vfs.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	v.Write(
		"pkg/foo/bar.go",
		"package foo\n\nimport \"fmt\"\n\nfunc Bar() {\n\tfmt.Println(\"world\")\n}\n",
	)
	v.Write(
		"pkg/foo/foo_test.go",
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n}\n",
	)
	fixes := LintVFS(v, option.New("", false), false)
	assert.Boolean(t, true, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/bar.go"))
}
