package post

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mattermost/mattermost/server/public/model"
	"log"
)

func Decode(v *model.WebSocketEvent) *model.Post {
	a, anyOkay := v.GetData()[constant.PostField]

	if !anyOkay {
		log.Panicf("no post field")
	}

	post, castOkay := a.(string)

	if !castOkay {
		log.Panicf("post field not string %T", a)
	}

	var result *model.Post
	notation.DecodeStrict(post, &result, false)

	return result
}
