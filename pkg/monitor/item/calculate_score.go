package item

func (i *Item) CalculateScore() {
	severity := i.severityScore()
	age := i.ageScore()
	stale := i.staleScore()
	triage := i.triageScore()

	i.Score = (severity + age + stale + triage) * i.collector.Source.BaseWeight
}
