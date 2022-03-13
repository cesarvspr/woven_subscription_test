package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWovenChallengeTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WovenChallengeTest Suite")
}
