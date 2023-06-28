package enforcer

import (
	"github.com/headingy/trireme/policy"
)

type nfLogger interface {
	start()
	stop()
}

type puInfoFunc func(string) (string, *policy.TagStore)
