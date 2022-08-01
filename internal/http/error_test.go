package http_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	httpService "microservices-boilerplate/internal/http"
	assertionErrors "microservices-boilerplate/internal/test/assertion/errors"
)

func TestError(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Error", func() {
	var (
		httpError httpService.Error
	)

	BeforeEach(func() {
		httpError = httpService.NewHttpError()
	})

	Context("Getting status code from error", func() {
		When("A record is not found", func() {
			It("Should return status not found", func() {
				err := httpError.GetHttpResponseError(gorm.ErrRecordNotFound)

				Expect(err.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
		When("User sent request with missing required value", func() {
			It("Should return status bad request", func() {
				err := httpError.GetHttpResponseError(gorm.ErrPrimaryKeyRequired)

				Expect(err.StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
		When("Error is not mapped", func() {
			It("Should return status internal server error", func() {
				err := httpError.GetHttpResponseError(assertionErrors.ErrGeneric)

				Expect(err.StatusCode).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
