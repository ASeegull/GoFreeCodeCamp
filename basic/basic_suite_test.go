package basic

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBasic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basic Suite")
}

var _ = Describe("Example tests", func() {
	Context("ReverseString Function", func() {
		It("should test that the function returns reversed string", func() {
			Expect(ReverseString("hello")).To(Equal("olleh"))
			Expect(ReverseString("Howdy")).To(Equal("ydwoH"))
			Expect(ReverseString("Greetings from Earth")).To(Equal("htraE morf sgniteerG"))
		})
	})

	Context("Factorialize Function", func() {
		It("should test that the function returns correct value of factorial", func() {
			Expect(Factorialize(5)).To(Equal(120))
			Expect(Factorialize(10)).To(Equal(3628800))
			Expect(Factorialize(20)).To(Equal(2432902008176640000))
			Expect(Factorialize(0)).To(Equal(1))
		})
	})

	Context("Is Palindrome Function", func() {
		It("should check if a string is palinndrome", func() {
			Expect(IsPalindrome("_eye)")).To(Equal(true))
			Expect(IsPalindrome("race car")).To(Equal(true))
			Expect(IsPalindrome("A man, a plan, a canal. Panama")).To(Equal(true))
			Expect(IsPalindrome("not a palindrome")).To(Equal(false))
			Expect(IsPalindrome("nope")).To(Equal(false))
			Expect(IsPalindrome("My age is 0, 0 si ega ym.")).To(Equal(true))
			Expect(IsPalindrome("1 eye for of 1 eye.")).To(Equal(false))
		})
	})
})
