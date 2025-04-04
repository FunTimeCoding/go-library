package loader

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestLoader(t *testing.T) {
	l := New()
	l.Load(
		system.Join(
			system.FindDirectoryUp(
				system.WorkingDirectory(),
				constant.Directory,
			),
			systemConstant.FixturePath,
			systemConstant.NotationPath,
		),
	)
	assert.Count(t, 2, l.contents)
	assert.Any(
		t,
		`{
    "Classified": "AnAlertName",
    "Reason": "A reason why the answer was chosen",
    "Answer": "not-yet-broken"
}
`,
		l.contents["response-1.json"],
	)
	actualMap := l.ToMap()
	assert.Any(
		t,
		map[string]string{
			"Classified": "AnAlertName",
			"Reason":     "A reason why the answer was chosen",
			"Answer":     "not-yet-broken",
		},
		actualMap["response-1.json"],
	)
	assert.Any(
		t,
		map[string]string{
			"Classified": "AnotherAlertName",
			"Reason":     "Another reason why the answer was chosen",
			"Answer":     "already-broken",
		},
		actualMap["response-2.json"],
	)
	reducedMap := actualMap
	delete(reducedMap, "response-2.json")
	assert.String(
		t,
		`1
Answer: not-yet-broken
Classified: AnAlertName
Reason: A reason why the answer was chosen`,
		l.ToText(reducedMap),
	)
}
