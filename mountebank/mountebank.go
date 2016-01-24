package mountebank

import (
	. "github.com/durmaze/gobank/builders"
	"github.com/parnurzeal/gorequest"
)

func CreateImposter(imposter Imposter) {
	gorequest.New().Post("http://localhost:2525/imposters").Send(imposter).End()
}