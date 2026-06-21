package face

import "gitlab.com/gitlab-org/api/client-go/v2"

type Forge interface {
	File(
		project int64,
		branch string,
		name string,
	) (*gitlab.File, error)
	MustFile(
		project int64,
		branch string,
		name string,
	) *gitlab.File
	CommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) (*gitlab.Commit, error)
	MustCommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) *gitlab.Commit
	RegistryRepositories(
		project int64,
		panicOnForbidden bool,
	) ([]*gitlab.RegistryRepository, error)
	Images(
		project int64,
		repository int64,
	) ([]*gitlab.RegistryRepositoryTag, error)
}
