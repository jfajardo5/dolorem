// MIT License
// Copyright (c) 2023 Julio "jfajardo5" Fajardo

// See end of file for extended license information

package dolorem

import (
	_ "embed"
	"math/rand"
	"strings"
	"time"
)

//go:embed data/latin.txt
var latinDictionary string

type Dolorem struct {
	// Dictionary is an array of words used for generating text.
	// It can be manually set to use external dictionaries.
	Dictionary []string

	// ParagraphStarter is the sentence that starts off the first paragraph in a set.
	// It can be manually set to customize your paragraph starter.
	ParagraphStarter string

	// Seed is used to generate random numbers.
	Seed *rand.Rand
}

// Default initializer for Dolorem
func Ipsum() Dolorem {
	return Dolorem{
		Dictionary:       loadLatinDictionary(),
		ParagraphStarter: "Dolorem ipsum dolor sit amet,",
		Seed:             rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Pull a random word from Dictionary
func (d *Dolorem) Word() string {
	index := d.Seed.Intn(len(d.Dictionary))
	return firstCharToUpper(d.Dictionary[index])
}

// Build a random sentence out of random Words
// Takes 1 optional in param to override default options
//
// @param length[0]: Override number of Words per Sentence
func (d *Dolorem) Sentence(arg ...int) string {
	numberOfWordsPerSentence := 7
	if len(arg) > 0 {
		numberOfWordsPerSentence = arg[0]
	}

	var words []string
	for i := 0; i < numberOfWordsPerSentence; i++ {
		words = append(words, strings.ToLower(d.Word()))
	}

	return firstCharToUpper(strings.Join(words, " ")) + "."
}

// Build random Paragraphs out of random Sentences
// Takes 3 optional int params to override default options
//
// @param args[0]: Override number of Paragraphs
// @param args[1]: Override number of Sentences per Paragraph
// @param args[2]: Override number of Words per Sentence
func (d *Dolorem) Paragraph(args ...int) string {
	numberOfParagraphs := 1
	numberOfSentences := 7
	numberOfWordsPerSentence := 7

	if len(args) > 0 {
		numberOfParagraphs = args[0]
	}
	if len(args) > 1 {
		numberOfSentences = args[1]
	}
	if len(args) > 2 {
		numberOfWordsPerSentence = args[2]
	}

	var paragraphs []string

	for i := 0; i < numberOfParagraphs; i++ {
		sentences := d.buildParagraphBody(numberOfSentences, numberOfWordsPerSentence, i)
		paragraphs = append(paragraphs, strings.Join(sentences, " "))
	}

	return strings.Join(paragraphs, "\n\n")
}

// Logic for building a Paragraph's body
func (d *Dolorem) buildParagraphBody(numberOfSentences int, numberOfWordsPerSentence int, currentIteration int) []string {
	var sentences []string
	if currentIteration == 0 && d.ParagraphStarter != "" {
		starter := strings.ToLower(d.ParagraphStarter)
		starter = firstCharToUpper(starter)
		sentences = append(sentences, starter)
	}
	for i := 0; i < numberOfSentences; i++ {
		sentences = append(sentences, d.buildSentenceBody(sentences, numberOfSentences, numberOfWordsPerSentence, i))
	}
	return sentences
}

// Logic for building Sentence body within a Paragraph
func (d *Dolorem) buildSentenceBody(sentences []string, numberOfSentences int, numberOfWordsPerSentence int, currentIteration int) string {
	sentence := strings.ReplaceAll(d.Sentence(numberOfWordsPerSentence), ".", "")
	sentence = strings.ToLower(sentence)
	sentence = d.addRandomCommaOrPeriod(sentence, numberOfSentences, currentIteration)
	if len(sentences) == 0 || endsWithPeriod(sentences[len(sentences)-1]) {
		sentence = firstCharToUpper(sentence)
	}
	return sentence
}

// Adds a period (.) or a comma(,) at the end of the parameter string
func (d *Dolorem) addRandomCommaOrPeriod(sentence string, numberOfSentences int, currentIndex int) string {
	if numberOfSentences-currentIndex <= 1 {
		return sentence + "."
	}
	rng := d.Seed.Intn(10)
	if rng < 3 {
		return sentence + "."
	}
	return sentence + ","
}

// Checks if the parameter string ends with a period (.)
func endsWithPeriod(sentence string) bool {
	if strings.LastIndex(sentence, ".") == len(sentence)-1 {
		return true
	}
	return false
}

// Returns the parameter string with its first char capitalized
func firstCharToUpper(text string) string {
	if len(text) > 1 {
		return strings.Replace(text, string(text[0]), strings.ToUpper(string(text[0])), 1)
	}
	return strings.ToUpper(text)
}

// Packs and returns data/latin.txt as an array
//
// This file contains the full Latin dictionary,
// with each word separated by a space (\n)
func loadLatinDictionary() []string {
	return strings.Split(latinDictionary, "\n")
}

// Dolorem is a highly inclusive (with a near complete Latin dictionary) Lorem Ipsum generator
//
// Copyright (c) 2023 Julio "jfajardo5" Fajardo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
