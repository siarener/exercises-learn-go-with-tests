package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {

		stackOfInts := new(Stack[int])
		AssertTrue(t, stackOfInts.IsEmpty())

		firstItem := 123
		stackOfInts.Push(firstItem)
		AssertFalse(t, stackOfInts.IsEmpty())

		secondItem := 654
		stackOfInts.Push(secondItem)

		firstPopItem, successful := stackOfInts.Pop()
		AssertTrue(t, successful)
		AssertEqual(t, firstPopItem, secondItem)

		secondPopItem, successful := stackOfInts.Pop()
		AssertTrue(t, successful)
		AssertEqual(t, secondPopItem, firstItem)

		AssertTrue(t, stackOfInts.IsEmpty())

		thirdPopItem, successful := stackOfInts.Pop()
		AssertFalse(t, successful)
		AssertEqual(t, thirdPopItem, 0)

		AssertEqual(t, firstItem+secondItem, 777)
	})

	/*
		 	t.Run("string stack", func(t *testing.T) {

				testStack := new(Stack[string])
				AssertTrue(t, testStack.IsEmpty())

				firstItem := "one-two-three"
				testStack.Push(firstItem)
				AssertFalse(t, testStack.IsEmpty())

				secondItem := "six-five-four"
				testStack.Push(secondItem)

				value, successful := testStack.Pop()
				AssertTrue(t, successful)
				AssertEqual(t, value, secondItem)

				value, successful = testStack.Pop()
				AssertTrue(t, successful)
				AssertEqual(t, value, firstItem)

				AssertTrue(t, testStack.IsEmpty())

				value, successful = testStack.Pop()
				AssertFalse(t, successful)
				AssertEqual(t, value, "")

				AssertEqual(t, firstItem+secondItem, "one-two-threesix-five-four")
			})
	*/
}
