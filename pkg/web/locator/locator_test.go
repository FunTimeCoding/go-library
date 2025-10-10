package locator

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"net/url"
	"path"
	"testing"
)

func TestLocator(t *testing.T) {
	assert.NotNil(t, New())
}

func TestParse(t *testing.T) {
	u, _ := url.Parse("https://example.org")
	u.Path = path.Join(u.Path, "/u", "1", "p")
	v := url.Values{}
	v.Add("a", "1")
	v.Add("b", "2")
	u.RawQuery = v.Encode()
	assert.String(t, "https://example.org/u/1/p?a=1&b=2", u.String())
}

func TestNew(t *testing.T) {
	l := NewHost("example.org")
	assert.String(t, "https://example.org", l.String())
}

func TestPath(t *testing.T) {
	l := NewHost("example.org").Path("/path")
	assert.String(t, "https://example.org/path", l.String())
}

func TestBasePath(t *testing.T) {
	l := NewHost("example.org").Base("/base").Path("/path")
	assert.String(t, "https://example.org/base/path", l.String())

	l = NewHost("example.org").Base("base").Path("/path")
	assert.String(t, "https://example.org/base/path", l.String())

	l = NewHost("example.org").Base("/base/").Path("/path")
	assert.String(t, "https://example.org/base/path", l.String())

	l = NewHost("example.org").Base("/base").Path("path")
	assert.String(t, "https://example.org/base/path", l.String())
}

func TestValue(t *testing.T) {
	l := NewHost("example.org").Add("a", "1").Add("b", "2")
	assert.String(t, "https://example.org?a=1&b=2", l.String())
}

func TestValuePath(t *testing.T) {
	l := NewHost("example.org").Path("/path").Add(
		"a",
		"1",
	).Add("b", "2")
	assert.String(t, "https://example.org/path?a=1&b=2", l.String())
}

func TestPort(t *testing.T) {
	l := NewHost("example.org").Port(8443)
	assert.String(t, "https://example.org:8443", l.String())
}

func TestPortPath(t *testing.T) {
	l := NewHost("example.org").Port(8443).Path("/path")
	assert.String(t, "https://example.org:8443/path", l.String())
}

func TestInsecure(t *testing.T) {
	l := New().Insecure().Host("example.org")
	assert.String(t, "http://example.org", l.String())
}

func TestFragment(t *testing.T) {
	l := NewHost("example.org").Fragment("section1")
	assert.String(t, "https://example.org#section1", l.String())
}

func TestAdd(t *testing.T) {
	l := NewHost("example.org").Add("a", "10").Add("a", "20")
	assert.String(t, "https://example.org?a=10&a=20", l.String())
}

func TestSet(t *testing.T) {
	l := NewHost("example.org").Set("a", "1").Set("a", "2")
	assert.String(t, "https://example.org?a=2", l.String())
}
