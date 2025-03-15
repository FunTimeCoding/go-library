package rule_list

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule"
	"github.com/go-test/deep"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
)

func (l *List) Add(i *rule.Rule) {
	for _, element := range l.rules {
		if i.Name != element.Name || i.Group != element.Group {
			continue
		}

		if i.RawAlert != nil && element.RawAlert != nil {
			if !LabelsSame(i.RawAlert, element.RawAlert) {
				if false {
					// This is fine
					fmt.Printf(
						"labels differ: %+v %+v\n",
						i.RawAlert.Labels,
						element.RawAlert.Labels,
					)
				}
			} else if i.RawAlert.Query != element.RawAlert.Query {
				if false {
					// This is fine
					fmt.Printf(
						"queries differ: %s %s\n",
						i.RawAlert.Query,
						element.RawAlert.Query,
					)
				}
			} else {
				if !cmp.Equal(
					i.RawAlert,
					element.RawAlert,
					cmpopts.IgnoreFields(
						v1.AlertingRule{},
						"EvaluationTime",
						"LastEvaluation",
					),
				) {
					log.Panicf(
						"duplicate RawAlert: name=%s comparison=%+v",
						element.Name,
						deep.Equal(i.RawAlert, element.RawAlert),
					)
				}
			}
		} else if i.RawRecord != nil && element.RawRecord != nil {
			if c := deep.Equal(i.RawAlert, element.RawAlert); len(c) != 0 {
				log.Panicf(
					"duplicate RawRecord: name=%s comparison=%+v",
					element.Name,
					c,
				)
			}
		} else {
			log.Panicf(
				"duplicate: name=%s existing=%+v new=%+v",
				element.Name,
				i,
				element,
			)
		}
	}

	l.rules = append(l.rules, i)
}
