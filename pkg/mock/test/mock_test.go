package test

import (
	"testing"

	"github.com/judegiordano/startup/pkg/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMockTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mock Test Suite")
}

var _ = Describe("Mock", func() {
	Context("User", func() {
		It("should create random user", func() {
			user := mock.CreateUser()
			Expect(len(user.Addresses)).To(Equal(1))
			Expect(len(user.PaymentMethods)).To(Equal(2))
		})
	})
})
