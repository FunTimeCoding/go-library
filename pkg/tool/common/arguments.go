package common

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
)

func Arguments() {
	pflag.String(
		argument.Host,
		environment.Optional(constant.QualifiedName),
		fmt.Sprintf("Host, fallback: %s", constant.QualifiedName),
	)
	pflag.String(
		argument.Token,
		environment.Fallback(
			TokenEnvironment,
			environment.Optional(constant.JobToken),
		),
		fmt.Sprintf(
			"Token, fallbacks: %s, %s",
			TokenEnvironment,
			constant.JobToken,
		),
	)
	pflag.String(
		argument.Owner,
		environment.Fallback(
			OwnerEnvironment,
			environment.Optional(constant.ProjectNamespace),
		),
		fmt.Sprintf(
			"Owner, fallbacks: %s, %s",
			OwnerEnvironment,
			constant.ProjectNamespace,
		),
	)
	pflag.String(
		argument.Repository,
		environment.Optional(RepositoryEnvironment),
		fmt.Sprintf("Repository, fallback: %s", RepositoryEnvironment),
	)
}
