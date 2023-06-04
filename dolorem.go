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

// TODO ADD TESTS!
// TODO ADD COMMENTS!

//go:embed data/latin.txt
var dictionary string

type Dolorem struct {
	dictionary        []string
	paragraph_starter string
	text              string
	seed              *rand.Rand
}

func Ipsum() Dolorem {
	dict := loadDictionary()
	return Dolorem{
		dictionary:        dict,
		seed:              rand.New(rand.NewSource(time.Now().Unix())),
		paragraph_starter: "dolorem ipsum dolor sit amet, ",
	}
}

func (d *Dolorem) Word() string {
	index := d.seed.Intn(len(d.dictionary) - 1)
	d.text = d.dictionary[index]
	return d.dictionary[index]
}

func (d *Dolorem) Sentence(length ...int) string {
	sen_len := 15
	if len(length) > 0 {
		sen_len = length[0] // Number of Words per Sentence
	}
	sentence := ""
	for i := 0; i < sen_len; i++ {
		sentence = sentence + d.Word()
		if i < sen_len-1 || i == 0 {
			sentence = sentence + " "
		}
	}
	d.text = sentence
	return sentence
}

func (d *Dolorem) Paragraph(length ...int) string {
	num := 1      // Number of Paragraphs
	par_len := 7  // Number of Sentences per Paragraph
	sen_len := 15 // Number of Words per Sentence
	if len(length) > 0 {
		num = length[0] // Number of Paragraphs
	}
	if len(length) > 1 {
		par_len = length[1] // Number of Sentences per Paragraph
	}
	if len(length) > 2 {
		sen_len = length[2] // Number of Words per Sentence
	}
	var paragraph = d.paragraph_starter
	for i := 0; i < num; i++ {
		for j := 0; j < par_len; j++ {
			paragraph = paragraph + d.Sentence(sen_len)
		}
		if i < num-1 {
			paragraph = paragraph + "\n\n"
		}
	}
	d.text = paragraph
	return paragraph
}

func loadDictionary() []string {
	return strings.Split(dictionary, "\n")
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
