package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"testing"
)

func TestCheckImportFix(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		false,
		&r,
	)
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
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		false,
		&r,
	)
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/foo_test.go"))
}

func TestCheckSkipsGeneratedFile(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/generated.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Generated() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		false,
		&r,
	)
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
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		false,
		&r,
	)
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
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		false,
		&r,
	)
	assert.Boolean(t, true, fixes.Has("pkg/foo/foo.go"))
	assert.Boolean(t, false, fixes.Has("pkg/foo/bar.go"))
}

func TestCheckStubCreated(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"go.mod",
		"module example\n\nrequire github.com/funtimecoding/go-library v0.0.1\n",
	)
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nfunc Foo() {}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		true,
		&r,
	)
	assert.String(
		t,
		"package foo\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestFoo(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		fixes.Read("pkg/foo/foo_test.go"),
	)
}

func TestCheckStubMainPackage(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"go.mod",
		"module example\n\nrequire github.com/funtimecoding/go-library v0.0.1\n",
	)
	v.Write(
		"cmd/foo/main.go",
		"package main\n\nfunc main() {}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		true,
		&r,
	)
	assert.String(
		t,
		"package main\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestStub(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		fixes.Read("cmd/foo/main_test.go"),
	)
}

func TestCheckStubToolPackage(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"go.mod",
		"module example\n\nrequire github.com/funtimecoding/go-library v0.0.1\n",
	)
	v.Write(
		"pkg/tool/gofoo/main.go",
		"package gofoo\n\nfunc Main() {}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		true,
		&r,
	)
	assert.String(
		t,
		"package gofoo\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/assert\"\n\t\"testing\"\n)\n\nfunc TestStub(t *testing.T) {\n\tassert.Stub(t)\n}\n",
		fixes.Read("pkg/tool/gofoo/main_test.go"),
	)
}

func TestCheckStubWithoutGoLibrary(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("go.mod", "module example\n")
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nfunc Foo() {}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		true,
		&r,
	)
	assert.String(
		t,
		"package foo\n\nimport \"testing\"\n\nfunc TestFoo(t *testing.T) {\n\tt.Helper()\n}\n",
		fixes.Read("pkg/foo/foo_test.go"),
	)
}

func TestCheckStubTestdata(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/lint/analyzer/naming/testdata/src/example/example.go",
		"package example\n\ntype Example struct{}\n",
	)
	var r output.Results
	fixes := Check(
		v,
		option.New("", false),
		false,
		false,
		true,
		&r,
	)
	assert.String(
		t,
		"package example\n\nimport \"testing\"\n\nfunc TestExample(t *testing.T) {\n\tt.Helper()\n}\n",
		fixes.Read("pkg/lint/analyzer/naming/testdata/src/example/example_test.go"),
	)
}

func TestCheckResultImportCollapsed(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nimport (\n\t\"fmt\"\n)\n\nfunc Foo() {\n\tfmt.Println(\"hello\")\n}\n",
	)
	var r output.Results
	Check(
		v,
		option.New("", false),
		true,
		false,
		false,
		&r,
	)
	assertApplied(
		t,
		r.Entries,
		"pkg/foo/foo.go",
		"collapsed single multi-import",
	)
}

func TestCheckResultBlankLineRemoved(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nfunc Foo() {\n\ta := 1\n\n\t_ = a\n}\n",
	)
	var r output.Results
	Check(
		v,
		option.New("", false),
		true,
		false,
		false,
		&r,
	)
	assertApplied(
		t,
		r.Entries,
		"pkg/foo/foo.go",
		"removed blank line (line 5)",
	)
}

func TestCheckResultBlankLineInserted(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"pkg/foo/foo.go",
		"package foo\n\nfunc Foo() {\n\ta := 1\n\tif a > 0 {\n\t\t_ = a\n\t}\n}\n",
	)
	var r output.Results
	Check(
		v,
		option.New("", false),
		true,
		false,
		false,
		&r,
	)
	assertApplied(
		t,
		r.Entries,
		"pkg/foo/foo.go",
		"inserted blank line (line 5)",
	)
}

func TestCheckResultMissingSentry(t *testing.T) {
	v := virtual_file_system.New()
	v.Write(
		"cmd/gofoo/main.go",
		"package main\n\nfunc main() {}\n",
	)
	var r output.Results
	Check(
		v,
		option.New("", false),
		true,
		false,
		false,
		&r,
	)
	assertBlocked(t, r.Entries, "cmd/gofoo", "missing sentry reporter")
}

func assertApplied(
	t *testing.T,
	entries []output.Result,
	path string,
	message string,
) {
	t.Helper()

	for _, e := range entries {
		if e.Path == path && e.Message == message && !e.Blocked {
			return
		}
	}

	t.Errorf(
		"expected applied result {path: %q, message: %q} not found",
		path,
		message,
	)
}

func assertBlocked(
	t *testing.T,
	entries []output.Result,
	path string,
	message string,
) {
	t.Helper()

	for _, e := range entries {
		if e.Path == path && e.Message == message && e.Blocked {
			return
		}
	}

	t.Errorf(
		"expected blocked result {path: %q, message: %q} not found",
		path,
		message,
	)
}
