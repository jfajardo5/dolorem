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

	// Text holds the last set of generated text, whether a word, sentence, or paragraph.
	Text string

	// seed is used to generate random numbers.
	seed *rand.Rand
}

// Default initializer for Dolorem
func Ipsum() Dolorem {
	dict := loadLatinDictionary()
	return Dolorem{
		Dictionary:       dict,
		seed:             rand.New(rand.NewSource(time.Now().Unix())),
		ParagraphStarter: "dolorem ipsum dolor sit amet, ",
	}
}

// Pull a random word from Dictionary
func (d *Dolorem) Word() string {
	index := d.seed.Intn(len(d.Dictionary) - 1)
	d.Text = d.Dictionary[index]
	return d.Dictionary[index]
}

// Build a random sentence out of random Words
// Takes 1 optional in param to override default options
//
// @param length[0]: Override number of Words per Sentence
func (d *Dolorem) Sentence(length ...int) string {
	senLen := 15 // Number of Words per Sentence
	if len(length) > 0 {
		senLen = length[0]
	}
	sentence := ""
	for i := 0; i < senLen; i++ {
		sentence = sentence + d.Word()
		if i < senLen-1 || i == 0 {
			sentence = sentence + " "
		}
	}
	d.Text = sentence
	return sentence
}

// Build random Paragraphs out of random Sentences
// Takes 3 optional int params to override default options
//
// @param length[0]: Override number of Paragraphs
// @param length[1]: Override number of Sentences per Paragraph
// @param length[2]: Override number of Words per Sentence
func (d *Dolorem) Paragraph(length ...int) string {
	num := 1     // Number of Paragraphs
	parLen := 7  // Number of Sentences per Paragraph
	senLen := 15 // Number of Words per Sentence
	if len(length) > 0 {
		num = length[0]
	}
	if len(length) > 1 {
		parLen = length[1]
	}
	if len(length) > 2 {
		senLen = length[2]
	}
	var paragraph = d.ParagraphStarter
	for i := 0; i < num; i++ {
		for j := 0; j < parLen; j++ {
			paragraph = paragraph + d.Sentence(senLen)
		}
		if i < num-1 {
			paragraph = paragraph + "\n\n"
		}
	}
	d.Text = paragraph
	return paragraph
}

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
