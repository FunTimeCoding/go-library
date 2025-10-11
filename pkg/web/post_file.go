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

	if false {
		// Not sure why, but file is already closed at the end
		defer errors.PanicClose(f)
	}

	r := NewPostBytes(url, f)
	r.SetBasicAuth(user, password)
	r.Header.Set(constant.ContentType, constant.Form)
	r.ContentLength = system.FileStat(f).Size()
	s := Send(Client(true), r)
	defer errors.PanicClose(s.Body)

	if !ResponseOkay(s) {
		log.Panicf("upload %d: %s", s.StatusCode, ReadString(s))
	}
}
