package test

import (
	"net/http"
	"testing"

	"github.com/judegiordano/startup/pkg/gravatar"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGravatarTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gravatar Test Suite")
}

var _ = Describe("Gravatar", func() {
	Context("Hash", func() {
		It("should create an email hash", func() {
			hash1 := gravatar.GenerateHash("mail@mail.com")
			hash2 := gravatar.GenerateHash("another+email@mail.com")
			Expect(hash1).To(Not(Equal(hash2)))
		})
		It("should build gravatar image", func() {
			hash := gravatar.GenerateHash("mail@mail.com")
			url := gravatar.BuildUrl(hash)
			res, err := http.Get(url)
			Expect(err).To(BeNil())
			contentType := res.Header.Get("Content-Type")
			Expect(contentType).To(Equal("image/png"))
		})
	})
})
