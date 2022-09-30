package tools

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/golang-module/carbon"
	"github.com/gosimple/slug"
	"github.com/judegiordano/startup/pkg/logger"
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

func PrintJson(i any) {
	var pretty bytes.Buffer
	bytes, e := json.MarshalIndent(i, "", "    ")
	if e != nil {
		logger.Error(e)
	}
	json.Indent(&pretty, bytes, "", "\t")
	fmt.Println(pretty.String())
}
