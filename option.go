package caseconv

import "github.com/leaqved/splitter"

type Option func(*Converter)

func WithSplitter(splitter *splitter.Splitter) Option {
	return func(c *Converter) {
		c.Splitter = splitter
	}
}

func WithDelimiter(delimeter rune) Option {
	return func(c *Converter) {
		c.Delimiter = delimeter
	}
}
