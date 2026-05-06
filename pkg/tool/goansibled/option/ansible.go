package option

type Ansible struct {
	Port            int
	Version         string
	Repository      string
	ClonePath       string
	AnsiblePath     string
	Playbook        []string
	PostgresLocator string
}
