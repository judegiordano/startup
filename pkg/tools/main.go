package tools

import (
	"crypto/rand"

	"github.com/golang-module/carbon"
	"github.com/gosimple/slug"
	"github.com/lucsky/cuid"
)

func Slugify(s *string) string {
	return slug.MakeLang(*s, "en")
}

func Cuid() (string, error) {
	id, err := cuid.NewCrypto(rand.Reader)
	return id, err
}

func Now() carbon.Carbon {
	return carbon.Now()
}
