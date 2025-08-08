package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"log"
)

func PostFile(
	name string,
	url string,
	user string,
	password string,
) {
	f := system.Open(name)
	defer errors.PanicClose(f)
	i := system.FileStat(f)
	r := NewPostBytes(url, f)
	r.SetBasicAuth(user, password)
	r.Header.Set(constant.ContentTypeHeader, constant.FormDataContentType)
	r.ContentLength = i.Size()
	s := Send(Client(true), r)
	defer errors.PanicClose(s.Body)

	if !ResponseOkay(s) {
		log.Panicf("upload %d: %s", s.StatusCode, ReadString(s))
	}
}
