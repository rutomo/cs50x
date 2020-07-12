package main

import "testing"

func TestSum(t *testing.T) {
	index := Readability("One fish. Two fish. Red fish. Blue fish.\n")
	if index > 0 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 0)
	}

	index = Readability("Would you like them here or there? I would not like them here or there. I would not like them anywhere.\n")

	if index != 2 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 2)
	}

	index = Readability("Congratulations! Today is your day. You're off to Great Places! You're off and away!\n")
	if index != 3 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 3)
	}

	index = Readability("Harry Potter was a highly unusual boy in many ways. For one thing, he hated the summer holidays more than any other time of year. For another, he really wanted to do his homework, but was forced to do it in secret, in the dead of the night. And he also happened to be a wizard.\n")
	if index != 5 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 5)
	}

	index = Readability("In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.\n")
	if index != 7 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 7)
	}

	index = Readability("Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, \"and what is the use of a book,\" thought Alice \"without pictures or conversation?\"\n")
	if index != 8 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 8)
	}

	index = Readability("When he was nearly thirteen, my brother Jem got his arm badly broken at the elbow. When it healed, and Jem's fears of never being able to play football were assuaged, he was seldom self-conscious about his injury. His left arm was somewhat shorter than his right; when he stood or walked, the back of his hand was at right angles to his body, his thumb parallel to his thigh.\n")
	if index != 8 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 8)
	}

	index = Readability("There are more things in Heaven and Earth, Horatio, than are dreamt of in your philosophy.\n")
	if index != 9 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 9)
	}

	index = Readability("it was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him.\n")
	if index != 10 {
		t.Errorf("Grade is incorrect, got: %d, want: %d.", index, 10)
	}

	index = Readability("A large class of computational problems involve the determination of properties of graphs, digraphs, integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other countable domains.\n")
	if index < 15 {
		t.Errorf("Grade is incorrect, got: %d, want: %d or higher.", index, 16)
	}
}
