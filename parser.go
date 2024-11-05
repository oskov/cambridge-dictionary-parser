package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Definition holds a word definition and its associated examples
type Definition struct {
	Definition string
	Examples   []string
}

// WordData holds the word and its first definition and examples
type WordData struct {
	Word        string
	Definitions []Definition
}

// DictionaryParser holds the settings for the parsing process
type DictionaryParser struct {
	Collector *colly.Collector
	BaseURL   string
}

// Option is a type for functional options used in configuring DictionaryParser
type Option func(*DictionaryParser)

// NewDictionaryParser is the constructor that applies functional options for custom configuration
func NewDictionaryParser(opts ...Option) *DictionaryParser {
	// Create a default parser with default values
	parser := &DictionaryParser{
		Collector: colly.NewCollector(
			colly.AllowedDomains("dictionary.cambridge.org"),
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
		),
		BaseURL: "https://dictionary.cambridge.org/dictionary/english/",
	}

	// Apply all options passed
	for _, opt := range opts {
		opt(parser)
	}

	return parser
}

// WithCustomCollector is an option to provide a custom colly.Collector
func WithCustomCollector(collector *colly.Collector) Option {
	return func(dp *DictionaryParser) {
		dp.Collector = collector
	}
}

// WithBaseURL is an option to provide a custom base URL
func WithBaseURL(baseURL string) Option {
	return func(dp *DictionaryParser) {
		dp.BaseURL = baseURL
	}
}

// ParseWord scrapes the first .pr block and returns the WordData and an error if any
func (dp *DictionaryParser) ParseWord(word string) (WordData, error) {
	var wordData WordData
	wordData.Word = word
	// foundFirstDefinition := false

	url := dp.BaseURL + word

	// Process only the first .pr block and its .def-blocks
	dp.Collector.OnHTML(".pr .dictionary", func(e *colly.HTMLElement) {
		// foundFirstDefinition = true
		fmt.Println("Found .pr block")
		fmt.Println(e.Text)
	})
	// dp.Collector.OnHTML(".pr .dictionary", func(e *colly.HTMLElement) {
	// 	if foundFirstDefinition {
	// 		return // Stop after processing the first .pr block
	// 	}

	// 	// Process all .def-block elements within the first .pr block
	// 	e.ForEach(".def-block", func(_ int, defBlock *colly.HTMLElement) {
	// 		definition := strings.TrimSpace(defBlock.ChildText(".def"))
	// 		if definition != "" {
	// 			def := Definition{
	// 				Definition: definition,
	// 			}

	// 			// Collect all example sentences under this .def-block
	// 			defBlock.ForEach(".examp", func(_ int, ex *colly.HTMLElement) {
	// 				example := strings.TrimSpace(ex.Text)
	// 				if example != "" {
	// 					def.Examples = append(def.Examples, example)
	// 				}
	// 			})

	// 			wordData.Definitions = append(wordData.Definitions, def)
	// 		}
	// 	})

	// 	foundFirstDefinition = true // Stop further processing after the first .pr block
	// })

	var resultErr error

	dp.Collector.OnError(func(_ *colly.Response, err error) {
		// log.Println("Something went wrong:", err)
		if resultErr == nil {
			resultErr = err
		} else {
			resultErr = fmt.Errorf("%v; %w", err, resultErr)
		}
	})

	dp.Collector.OnRequest(func(r *colly.Request) {
		// log.Println("Visiting", r.URL)
	})

	// Visit the word page
	err := dp.Collector.Visit(url)
	if err != nil {
		return WordData{}, fmt.Errorf("failed to scrape: %w", err)
	}

	// Return the WordData struct with no error
	return wordData, resultErr
}
