package test

import (
	"encoding/json"
	"net/http"
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
	Context("Fetch", func() {
		It("should perform a get request", func() {
			type Todos []struct {
				UserId    int    `json:"userId"`
				Id        int    `json:"id"`
				Title     string `json:"title"`
				Completed bool   `json:"completed"`
			}
			r, err := http.Get("https://jsonplaceholder.typicode.com/todos")
			Expect(err).To(BeNil())
			defer r.Body.Close()
			var todos Todos
			json.NewDecoder(r.Body).Decode(&todos)
			Expect(todos[0].UserId).To(Equal(1))
		})
	})
})
