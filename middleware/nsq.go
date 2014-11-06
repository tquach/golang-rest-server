package middleware

import (
	// "github.com/bitly/go-nsq"

	"github.com/go-martini/martini"
)

func Nsq() martini.Handler {
	return func(c martini.Context) {
		c.Next()
	}
}
