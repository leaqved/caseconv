package caseconv

type Rule func(index int, words []string) bool

type CharRule func(index int, word string) bool

func WordIsFirst() Rule {
	return func(index int, words []string) bool {
		return index == 0
	}
}

func WordInMiddle() Rule {
	return func(index int, words []string) bool {
		return index > 0 && index < len(words)-1
	}
}

func WordIsLast() Rule {
	return func(index int, words []string) bool {
		return index == len(words)-1
	}
}

func CharIsFirst() CharRule {
	return func(index int, word string) bool {
		return index == 0
	}
}

func CharInMiddle() CharRule {
	return func(index int, word string) bool {
		return index > 0 && index < len(word)-1
	}
}

func CharIsLast() CharRule {
	return func(index int, word string) bool {
		return index == len(word)-1
	}
}
