package gobank_test

import (
	"encoding/json"

	"github.com/durmaze/gobank"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"sync"
)

var _ = Describe("Stub Builder Tests", func() {

	Describe("When building a Stub with single Response", func() {
		var (
			actualStubAsMap map[string]interface{}
			once            sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()
				actualStub := gobank.Stub().Responses(expectedResponse).Build()

				jsonBytes, _ := json.Marshal(actualStub)
				actualStubAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualStubAsMap)
			})
		})

		It("should create a Stub that has one Response", func() {
			responses := actualStubAsMap["responses"]

			Expect(responses).To(HaveLen(1))
		})
	})

	Describe("When building a Stub with single Response and multiple, different Predicates", func() {
		var (
			actualStubAsMap map[string]interface{}
			once            sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

				expectedPredicate1 := predicates.Equals().Path("/test-path").Build()
				expectedPredicate2 := predicates.Contains().Method("POST").Build()

				actualStub := gobank.Stub().Responses(expectedResponse).Predicates(expectedPredicate1, expectedPredicate2).Build()

				jsonBytes, _ := json.Marshal(actualStub)
				actualStubAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualStubAsMap)
			})
		})

		It("should create a Stub that has one Response", func() {
			responses := actualStubAsMap["responses"]

			Expect(responses).To(HaveLen(1))
		})

		It("should create a Stub that has \"Is\" Response", func() {
			resps := actualStubAsMap["responses"].([]interface{})

			Expect(resps[0]).To(HaveKey(responses.Is().Build().Type()))
		})

		It("should create a Stub that has two Predicates", func() {
			predicates := actualStubAsMap["predicates"]

			Expect(predicates).To(HaveLen(2))
		})

		It("should create a Stub that has an \"Equals\" Predicate", func() {
			preds := actualStubAsMap["predicates"].([]interface{})

			Expect(preds[0]).To(HaveKey(predicates.Equals().Build().Type()))
		})

		It("should create a Stub that has a \"Contains\" Predicate", func() {
			preds := actualStubAsMap["predicates"].([]interface{})

			Expect(preds[1]).To(HaveKey(predicates.Contains().Build().Type()))
		})

	})

})
