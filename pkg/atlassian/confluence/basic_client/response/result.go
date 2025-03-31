package response

type Result struct {
	Id                  string              `json:"id"`
	Type                string              `json:"type"`
	Status              string              `json:"status"`
	Title               string              `json:"title"`
	ChildTypes          ChildTypes          `json:"childTypes"`
	MacroRenderedOutput MacroRenderedOutput `json:"macroRenderedOutput"`
	Restrictions        Restrictions        `json:"restrictions"`
	Expandable          Expandable          `json:"_expandable"`
	Links               Links               `json:"_links"`
}
