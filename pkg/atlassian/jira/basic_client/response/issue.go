package response

import "github.com/trivago/tgo/tcontainer"

type Issue struct {
	Expand string `json:"expand"`
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		StatusCategoryChangeDate string `json:"statuscategorychangedate"`
		FixVersions              []any  `json:"fixVersions"`
		StatusCategory           struct {
			Self      string `json:"self"`
			Id        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
		Resolution any `json:"resolution"`
		LastViewed any `json:"lastViewed"`
		Priority   struct {
			Self    string `json:"self"`
			IconUrl string `json:"iconUrl"`
			Name    string `json:"name"`
			Id      string `json:"id"`
		} `json:"priority"`
		Labels                        []any `json:"labels"`
		TimeEstimate                  any   `json:"timeestimate"`
		AggregateTimeOriginalEstimate any   `json:"aggregatetimeoriginalestimate"`
		Versions                      []any `json:"versions"`
		IssueLinks                    []struct {
			Id   string `json:"id"`
			Self string `json:"self"`
			Type struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Inward  string `json:"inward"`
				Outward string `json:"outward"`
				Self    string `json:"self"`
			} `json:"type"`
			OutwardIssue struct {
				Id     string `json:"id"`
				Key    string `json:"key"`
				Self   string `json:"self"`
				Fields struct {
					Summary string `json:"summary"`
					Status  struct {
						Self           string `json:"self"`
						Description    string `json:"description"`
						IconUrl        string `json:"iconUrl"`
						Name           string `json:"name"`
						Id             string `json:"id"`
						StatusCategory struct {
							Self      string `json:"self"`
							Id        int    `json:"id"`
							Key       string `json:"key"`
							ColorName string `json:"colorName"`
							Name      string `json:"name"`
						} `json:"statusCategory"`
					} `json:"status"`
					Priority struct {
						Self    string `json:"self"`
						IconUrl string `json:"iconUrl"`
						Name    string `json:"name"`
						Id      string `json:"id"`
					} `json:"priority"`
					IssueType struct {
						Self           string `json:"self"`
						Id             string `json:"id"`
						Description    string `json:"description"`
						IconUrl        string `json:"iconUrl"`
						Name           string `json:"name"`
						Subtask        bool   `json:"subtask"`
						AvatarId       int    `json:"avatarId"`
						HierarchyLevel int    `json:"hierarchyLevel"`
					} `json:"issuetype"`
				} `json:"fields"`
			} `json:"outwardIssue"`
		} `json:"issuelinks"`
		Assignee any `json:"assignee"`
		Status   struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconUrl        string `json:"iconUrl"`
			Name           string `json:"name"`
			Id             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				Id        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Components            []any `json:"components"`
		AggregateTimeEstimate any   `json:"aggregatetimeestimate"`
		Creator               struct {
			Self         string `json:"self"`
			AccountId    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				X48 string `json:"48x48"`
				X24 string `json:"24x24"`
				X16 string `json:"16x16"`
				X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
			AccountType string `json:"accountType"`
		} `json:"creator"`
		Subtasks []any `json:"subtasks"`
		Reporter struct {
			Self         string `json:"self"`
			AccountId    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				X48 string `json:"48x48"`
				X24 string `json:"24x24"`
				X16 string `json:"16x16"`
				X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
			AccountType string `json:"accountType"`
		} `json:"reporter"`
		AggregateProgress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Progress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"progress"`
		Votes struct {
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
			HasVoted bool   `json:"hasVoted"`
		} `json:"votes"`
		IssueType struct {
			Self           string `json:"self"`
			Id             string `json:"id"`
			Description    string `json:"description"`
			IconUrl        string `json:"iconUrl"`
			Name           string `json:"name"`
			Subtask        bool   `json:"subtask"`
			AvatarId       int    `json:"avatarId"`
			HierarchyLevel int    `json:"hierarchyLevel"`
		} `json:"issuetype"`
		TimeSpent any `json:"timespent"`
		Project   struct {
			Self           string `json:"self"`
			Id             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
			Simplified     bool   `json:"simplified"`
			AvatarUrls     struct {
				X48 string `json:"48x48"`
				X24 string `json:"24x24"`
				X16 string `json:"16x16"`
				X32 string `json:"32x32"`
			} `json:"avatarUrls"`
		} `json:"project"`
		AggregateTimeSpent any `json:"aggregatetimespent"`
		ResolutionDate     any `json:"resolutiondate"`
		WorkRatio          int `json:"workratio"`
		Watches            struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
			IsWatching bool   `json:"isWatching"`
		} `json:"watches"`
		IssueRestriction struct {
			IssueRestrictions struct{} `json:"issuerestrictions"`
			ShouldDisplay     bool     `json:"shouldDisplay"`
		} `json:"issuerestriction"`
		Created              string   `json:"created"`
		Updated              string   `json:"updated"`
		TimeOriginalEstimate any      `json:"timeoriginalestimate"`
		Description          any      `json:"description"`
		TimeTracking         struct{} `json:"timetracking"`
		Security             any      `json:"security"`
		Attachment           []any    `json:"attachment"`
		Summary              string   `json:"summary"`
		Environment          any      `json:"environment"`
		DueDate              any      `json:"duedate"`
		Comment              struct {
			Comments   []any  `json:"comments"`
			Self       string `json:"self"`
			MaxResults int    `json:"maxResults"`
			Total      int    `json:"total"`
			StartAt    int    `json:"startAt"`
		} `json:"comment"`
		Unknowns tcontainer.MarshalMap
	} `json:"fields"`
}
