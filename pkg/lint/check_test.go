package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"testing"
)

func TestCheckImportFix(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.String(
		t,
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
		fixes.Read("pkg/foo/foo.go"),
	)
}

func TestCheckCleanNoFixes(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	v.Write(
		"pkg/foo/foo_test.go",
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n\tt.Parallel()\n}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo_test.go"))
}

func TestCheckSkipsGeneratedFile(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/generated.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Generated() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/generated.go"))
}

func TestCheckVariableFlaggedNoFix(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport \"fmt\"\n\nfunc Foo() {\n\terr := fmt.Errorf(\"oops\")\n\t_ = err\n}\n",
	)
	v.Write(
		"pkg/foo/foo_test.go",
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo.go"))
}

func TestCheckMultipleFilesOnlyBrokenFixed(t *testing.T) {
	v := virtual_file_system.New()
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
	fixes := Check(v, option.New("", false), false)
	assert.Boolean(t, true, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/bar.go"))
}

func TestCheckStubCreated(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nfunc Foo() {}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.String(
		t,
		"package foo\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestFoo(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		fixes.Read("pkg/foo/foo_test.go"),
	)
}

func TestCheckStubMainPackage(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"cmd/foo/main.go",
		"package main\n\nfunc main() {}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.String(
		t,
		"package main\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestStub(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		fixes.Read("cmd/foo/main_test.go"),
	)
}

func TestCheckStubTestdata(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/lint/analyzer/naming/testdata/src/example/example.go",
		"package example\n\ntype Example struct{}\n",
	)
	fixes := Check(v, option.New("", false), false)
	assert.String(
		t,
		"package example\n\nimport \"testing\"\n\nfunc TestExample(t *testing.T) {\n\tt.Helper()\n}\n",
		fixes.Read("pkg/lint/analyzer/naming/testdata/src/example/example_test.go"),
	)
}
