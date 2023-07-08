package speller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpeller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Speller Suite")
}
