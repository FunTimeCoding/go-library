package log_manager

import "time"

type Log struct {
	FileName string `json:"FileName"`
	Players  []struct {
		Name                string `json:"Name"`
		AccountName         string `json:"AccountName"`
		Subgroup            int    `json:"Subgroup"`
		Profession          int    `json:"Profession"`
		EliteSpecialization int    `json:"EliteSpecialization"`
		GuildGUID           string `json:"GuildGuid"`
		Tag                 int    `json:"Tag"`
	} `json:"Players"`
	EncounterResult int `json:"EncounterResult"`
	EncounterMode   int `json:"EncounterMode"`
	Encounter       int `json:"Encounter"`
	MapID           int `json:"MapId"`
	PointOfView     struct {
		CharacterName string `json:"CharacterName"`
		AccountName   string `json:"AccountName"`
	} `json:"PointOfView"`
	GameBuild          int       `json:"GameBuild"`
	EvtcVersion        string    `json:"EvtcVersion"`
	GameLanguage       int       `json:"GameLanguage"`
	MainTargetName     string    `json:"MainTargetName"`
	HealthPercentage   any       `json:"HealthPercentage"`
	EncounterStartTime time.Time `json:"EncounterStartTime"`
	EncounterDuration  string    `json:"EncounterDuration"`
	DpsReportEIUpload  struct {
		URL             any `json:"Url"`
		ProcessingError any `json:"ProcessingError"`
		UploadTime      any `json:"UploadTime"`
	} `json:"DpsReportEIUpload"`
	ParsingStatus               int       `json:"ParsingStatus"`
	ParseTime                   time.Time `json:"ParseTime"`
	ParseMilliseconds           int       `json:"ParseMilliseconds"`
	ParsingException            any       `json:"ParsingException"`
	ParsingVersion              string    `json:"ParsingVersion"`
	IsEncounterStartTimePrecise bool      `json:"IsEncounterStartTimePrecise"`
	Tags                        []any     `json:"Tags"`
	IsFavorite                  bool      `json:"IsFavorite"`
	LogExtras                   struct {
		FractalExtras any `json:"FractalExtras"`
	} `json:"LogExtras"`
	MissingEncounterStart bool `json:"MissingEncounterStart"`
}
