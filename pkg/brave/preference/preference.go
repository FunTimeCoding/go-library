package preference

type Preference struct {
	Cookie struct {
		Accounts string `json:"last_list_accounts_binary_data"`
	} `json:"gaia_cookie"`
}
