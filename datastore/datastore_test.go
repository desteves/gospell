package datastore_test

import (
	. "github.com/desteves/gospell/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datastore", func() {
	var ds Datastore
	var source string
	var err error
	Context("When the connection string is not set\n", func() {
		BeforeEach(func() {
			ds = &Local{}
			err = ds.Open("")
		})
		It("can generate an error\n", func() {
			Expect(err).Should(HaveOccurred())
		})
	})
	Context("When the connection string is valid\n", func() {
		BeforeEach(func() {
			source = "internal"
			ds = &Local{}
			err = ds.Open(source)
		})

		Context("When the datastore is opened\n", func() {
			It("can connect\n", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			Context("When the phonetic alphabet is defined\n", func() {
				It("can find a codeword\n", func() {
					_, err := ds.Read("z")
					Expect(err).ShouldNot(HaveOccurred())
				})
				It("can find codewords\n", func() {
					_, err := ds.Read("albert")
					Expect(err).ShouldNot(HaveOccurred())
				})

			})
		})
	})
})
