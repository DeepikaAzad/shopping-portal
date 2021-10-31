package users_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	gomega "github.com/onsi/gomega"
)

func TestUsers(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "Users Suite")
}
