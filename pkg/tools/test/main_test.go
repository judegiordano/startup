package test

import (
	"testing"

	"github.com/judegiordano/startup/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tools Test Suite")
}

var _ = Describe("Tools", func() {
	Context("Slug", func() {
		It("should create correct slug", func() {
			name := "jOhN     dOe"
			slug := tools.Slugify(&name)
			Expect(slug).To(Equal("john-doe"))
		})
	})
	Context("CUID", func() {
		It("should create a unique cuid", func() {
			cuid1, err1 := tools.Cuid()
			cuid2, err2 := tools.Cuid()
			Expect(err1).To(BeNil())
			Expect(err2).To(BeNil())
			Expect(cuid1).To(Not(Equal(cuid2)))
		})
	})
})
