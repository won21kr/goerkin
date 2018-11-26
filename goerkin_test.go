package goerkin_test

import (
	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("Inline Definitions", func() {
	var total int

	steps := NewSteps()

	Scenario("Everything inline", func() {
		steps.Given("the current total is cleared", func() {
			total = 0
		})

		steps.When("3 is added", func() {
			total = total + 3
		})

		steps.When("2 is subtracted", func() {
			total = total - 2
		})

		steps.Then("the total is 1", func() {
			Expect(total).To(Equal(1))
		})

	})

})

var _ = Describe("Reuse Definitions", func() {
	var total int

	steps := Define(func(define Definitions) {
		define.Given("^the current total is cleared$", func() {
			total = 0
		})

		define.When("^3 is added$", func() {
			total = total + 3
		})

		define.When("^2 is subtracted$", func() {
			total = total - 2
		})

		define.Then("^the total is 1$", func() {
			Expect(total).To(Equal(1))
		})
	})

	Scenario("No definitions", func() {
		steps.Given("the current total is cleared")

		steps.When("3 is added")
		steps.And("2 is subtracted")

		steps.Then("the total is 1")
	})
})

var _ = Describe("Scenario first", func() {
	steps := NewSteps()

	Scenario("Scenario before steps", func() {
		steps.Then("it works")
	})

	steps.Define(func(define Definitions) {
		define.Then("^it works$", func() {
			Succeed()
		})
	})

})

var _ = Describe("Groups as params", func() {
	var total int
	steps := NewSteps()

	Scenario("Use RE groups as params", func() {
		steps.Given("4 and 6")
		steps.Then("the answer is 10")
	})

	steps.Define(func(define Definitions) {
		// TODO, reflect and report so the func can expect ints
		define.Given(`^(\d+) and (\d+)$`, func(first, second string) {
			firstInt, _ := strconv.Atoi(first)
			secondInt, _ := strconv.Atoi(second)
			total = firstInt + secondInt
		})
		define.Then(`^the answer is (\d+)$`, func(answer string) {
			answerInt, _ := strconv.Atoi(answer)
			Expect(total).To(Equal(answerInt))
		})
	})

})

