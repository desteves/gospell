package speller_test

import (
	"github.com/desteves/gospell/datastore"
	. "github.com/desteves/gospell/speller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NATO Speller tests\n", func() {
	Context("When the nato codewords are set up\n", func() {
		var source string

		BeforeEach(func() {
			source = "internal"
		})
		It("should create a new nato provider", func() {
			n, err := NewNATOProvider(source)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(n.DS).ShouldNot(BeNil())
		})

		Context("When getting a valid word", func() {
			var input string
			var nato NATO
			var expected []string
			BeforeEach(func() {
				expected = []string{"Apple", "Banana", "Coconut"}
				input = "abc"
				nato = NATO{
					DS: &datastore.Mock{
						Client: map[rune]string{
							rune('A'): "Apple",
							rune('B'): "Banana",
							rune('C'): "Coconut",
						},
					},
				}
			})
			It("can retrieve the codewords", func() {
				output, err := nato.GetCodeWords(input)
				Expect(output).Should(Equal(expected))
				Expect(err).ShouldNot(HaveOccurred())
			})

		})

	})

})
