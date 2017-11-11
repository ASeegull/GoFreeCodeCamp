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
		It("should check if a string is palindrome", func() {
			Expect(IsPalindrome("_eye)")).To(Equal(true))
			Expect(IsPalindrome("race car")).To(Equal(true))
			Expect(IsPalindrome("A man, a plan, a canal. Panama")).To(Equal(true))
			Expect(IsPalindrome("not a palindrome")).To(Equal(false))
			Expect(IsPalindrome("nope")).To(Equal(false))
			Expect(IsPalindrome("My age is 0, 0 si ega ym.")).To(Equal(true))
			Expect(IsPalindrome("1 eye for of 1 eye.")).To(Equal(false))
		})
	})

	Context("FindLongestWord", func() {
		It("should return length of the longest word", func() {
			Expect(FindLongestWord("The quick brown fox jumped over the lazy dog")).To(Equal(6))
			Expect(FindLongestWord("May the force be with you")).To(Equal(5))
			Expect(FindLongestWord("Google do a barrel roll")).To(Equal(6))
			Expect(FindLongestWord("What is the average airspeed velocity of an unladen swallow")).To(Equal(8))
			Expect(FindLongestWord("What if we try a super-long word such as otorhinolaryngology")).To(Equal(19))
		})
	})

	Context("TitleCase", func() {
		It("should type every word like this 'Word' length of the longest word", func() {
			Expect(TitleCase("I'm a little tea pot")).To(Equal("I'm A Little Tea Pot"))
			Expect(TitleCase("sHoRt AnD sToUt")).To(Equal("Short And Stout"))
			Expect(TitleCase("HERE IS MY HANDLE HERE IS MY SPOUT")).To(Equal("Here Is My Handle Here Is My Spout"))
		})
	})

	Context("LargestOfFour", func() {
		It("should return array of 4 max integers", func() {
			Expect(TitleCase("I'm a little tea pot")).To(Equal("I'm A Little Tea Pot"))
			Expect(TitleCase("sHoRt AnD sToUt")).To(Equal("Short And Stout"))
			Expect(TitleCase("HERE IS MY HANDLE HERE IS MY SPOUT")).To(Equal("Here Is My Handle Here Is My Spout"))
		})
	})

	Context("ConfirmEnding", func() {
		It("should return true if string ends with passed substr", func() {
			Expect(ConfirmEnding("Bastian", "n")).To(Equal(true))
			Expect(ConfirmEnding("Connor", "n")).To(Equal(false))
			Expect(ConfirmEnding("Walking on water and developing software from a specification are easy if both are frozen", "specification")).To(Equal(false))
			Expect(ConfirmEnding("He has to give me a new name", "name")).To(Equal(true))
			Expect(ConfirmEnding("Open sesame", "same")).To(Equal(true))
			Expect(ConfirmEnding("Open sesame", "pen")).To(Equal(false))
			Expect(ConfirmEnding("If you want to save our world, you must hurry. We dont know how much longer we can withstand the nothing", "mountain")).To(Equal(false))
		})
	})

	Context("RepeatStringNumTimes", func() {
		It("should repeat string n times without spaces", func() {
			Expect(RepeatStringNumTimes("*", 3)).To(Equal("***"))
			Expect(RepeatStringNumTimes("abc", 3)).To(Equal("abcabcabc"))
			Expect(RepeatStringNumTimes("abc", 3)).To(Equal("abc"))
			Expect(RepeatStringNumTimes("abc", -2)).To(Equal(""))
		})
	})

	Context("TruncateString", func() {
		It("should return string of n chars, that ends with ... if it was truncated", func() {
			Expect(TruncateString("Peter Piper picked a peck of pickled peppers", 14)).To(Equal("Peter Piper..."))
			Expect(TruncateString("A-", 1)).To(Equal("A..."))
			Expect(TruncateString("Absolutely Longer", 2)).To(Equal("Ab..."))
			Expect(TruncateString("abc", -2)).To(Equal("..."))
			Expect(TruncateString("abc", 9)).To(Equal("abc"))
		})
	})

	Context("ChunkArrayInGroups", func() {
		ex1 := []int{1, 3, 5, 6, 2, 4, 7, 8}
		out1 := [][]int{
			{1, 3},
			{5, 6},
			{2, 4},
			{7, 8},
		}
		ex2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		out2 := [][]int{
			{0, 1, 2, 3},
			{4, 5, 6, 7},
			{8},
		}
		It("should return string of n chars, that ends with ... if it was truncated", func() {
			Expect(ChunkArrayInGroups(ex1, 2)).To(Equal(out1))
			Expect(ChunkArrayInGroups(ex2, 4)).To(Equal(out2))
		})
	})
})
