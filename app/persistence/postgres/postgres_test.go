package postgres_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPostgresRepository(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres repository test.")
}