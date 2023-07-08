package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/desteves/gospell/api"
	"github.com/desteves/gospell/speller"
)

var _ = Describe("Api", func() {
	Context("When the GET spell route is set\n", func() {
		var mockAPI API
		var r *gin.Engine
		var w *httptest.ResponseRecorder

		BeforeEach(func() {
			mockAPI = API{
				SpellerProvider: &speller.Mock{},
			}
			w = httptest.NewRecorder()
			_, r = gin.CreateTestContext(w)
			r.GET(SpellEndpoint, mockAPI.SpellEndpointHandler)
		})

		Context("when the input query param is missing\n", func() {

			It("can return the appropiate error msg", func() {

				req, _ := http.NewRequest("GET", "/spell", nil)
				req.ParseForm()
				r.ServeHTTP(w, req)
				Expect(w.Code).Should(Equal(http.StatusBadRequest))

				req, _ = http.NewRequest("GET", "/spell?badparamt=error", nil)
				req.ParseForm()
				r.ServeHTTP(w, req)
				Expect(w.Code).Should(Equal(http.StatusBadRequest))

			})
		})

		Context("when the input query param is provided", func() {
			var expected []string
			BeforeEach(func() {
				expected = []string{"Apple", "Banana", "Coconut"}

			})
			It("can return the corresponding codewords", func() {

				editable, _ := mockAPI.SpellerProvider.(*speller.Mock)
				editable.CodeWords = map[rune]string{
					rune('A'): "Apple",
					rune('B'): "Banana",
					rune('C'): "Coconut",
				}
				mockAPI.SpellerProvider = editable

				req, _ := http.NewRequest("GET", "/spell?input=abc", nil)
				req.ParseForm()
				r.ServeHTTP(w, req)
				Expect(w.Code).Should(Equal(http.StatusOK))
				var actual []string
				err := json.Unmarshal([]byte(w.Body.Bytes()), &actual)

				Expect(err).NotTo(HaveOccurred())
				Expect(actual).Should(Equal(expected))
			})
		})

	})
})
