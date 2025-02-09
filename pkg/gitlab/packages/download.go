package packages

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	library "github.com/funtimecoding/go-library/pkg/gitlab/request"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
)

func Download(
	link string,
	token string,
	outputFile string,
) {
	request := web.NewGet(link)
	library.PrivateToken(request, token)
	response := web.Send(web.Client(true), request)
	defer errors.PanicClose(response.Body)
	errors.PanicStatus(response)
	f := system.Create(outputFile)
	defer errors.PanicClose(f)
	system.Copy(response.Body, f)
}
