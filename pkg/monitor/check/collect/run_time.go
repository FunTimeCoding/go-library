package collect

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/monitor/collector"
	library "github.com/funtimecoding/go-library/pkg/time"
	"k8s.io/apimachinery/pkg/util/duration"
	"time"
)

func runTime(
	s *collector.Collector,
	t time.Time,
) {
	fmt.Printf(
		"%s %s %s\n",
		t.Format(library.DateSecond),
		s.Name,
		duration.HumanDuration(s.Interval),
	)
	fetch.Run(s.Name)
}
