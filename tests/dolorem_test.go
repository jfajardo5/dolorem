// MIT License
// Copyright (c) 2023 Julio "jfajardo5" Fajardo

// See end of file for extended license information

package dolorem

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/jfajardo5/dolorem"
)

// dolorem.Ipsum() (Default Initializer)
func TestDefaultInitializer(t *testing.T) {
	lorem := dolorem.Ipsum()
	t.Error("Testing workflow")
	if lorem.ParagraphStarter != "Dolorem ipsum dolor sit amet," {
		t.Error("@TestDefaultInitializer: Ipsum() did not return default ParagraphStarter.")
	}

	if len(lorem.Dictionary) < 3086 {
		t.Error("@TestDefaultInitializer: Ipsum() did not return full Dictionary.")
	}
}

// dolorem.Word() - Returns a single, random word from the dictionary
func TestWord(t *testing.T) {
	lorem := dolorem.Ipsum()
	lorem.Dictionary = []string{"test", "go"}
	result := lorem.Word()

	if result != "Test" && result != "Go" {
		t.Error("@TestWord: Word() did not return a value from it's dictionary")
	}
}

// dolorem.Sentence() - Returns a single Sentence with 7 random Words from the dictionary
func TestDefaultSentence(t *testing.T) {
	lorem := dolorem.Ipsum()
	result := strings.Split(lorem.Sentence(), " ")

	if len(result) != 7 {
		t.Error("@TestDefaultSentence: Sentence() did not return Sentence() of expected default length")
	}

	if !unicode.IsUpper(rune(result[0][0])) {
		t.Error("@TestDefaultSentence: Sentence() did not return Sentence() with expected uppercase first char")
	}
}

// dolorem.Sentence(5) - Returns a single Sentence with 5 (custom length) random Words from the dictionary
func TestCustomSentence(t *testing.T) {
	lorem := dolorem.Ipsum()
	result := strings.Split(lorem.Sentence(5), " ")

	if len(result) != 5 {
		t.Error("@TestCustomSentence: Sentence() did not return Sentence() of expected custom length (5)")
	}

	if !unicode.IsUpper(rune(result[0][0])) {
		t.Error("@TestCustomSentence: Custom Sentence() did not return Sentence() with expected uppercase first char")
	}
}

// dolorem.Paragraph() - Returns a single Paragraph with 7 Sentences populated from 7 random Words
func TestDefaultParagraph(t *testing.T) {
	lorem := dolorem.Ipsum()
	result := lorem.Paragraph()
	if !strings.Contains(result, lorem.ParagraphStarter) {
		t.Error("@TestDefaultParagraph: Paragraph() did not contain default ParagraphStarter")
	}

	paragraphs := strings.Split(result, "\n\n")
	if len(paragraphs) != 1 {
		t.Error("@TestDefaultParagraph: Paragraph() did not return Paragraph of expected default number")
	}

	if !unicode.IsUpper(rune(paragraphs[0][0])) {
		t.Error("@TestDefaultParagraph: Paragraph() did not return Paragraph with expected uppercase first char")
	}

	paragraphs[0] = strings.Replace(paragraphs[0], "Dolorem ipsum dolor sit amet,", "", 1)
	words := strings.Split(paragraphs[0], " ")
	if len(words)/7 != 7 {
		t.Error("@TestDefaultParagraph: Paragraph() did not return Paragraph with expected default Sentence and Word count")
	}
}

// dolorem.Paragraph(5) - Returns 5 Paragraphs with 7 Sentences populated from 7 random Words per Paragraph
// dolorem.Paragraph(2, 5) - Returns 2 Paragraphs with 5 Sentences populated from 7 random Words per Paragraph
// dolorem.Paragraph(3, 5, 10) - Returns 3 Paragraphs with 5 Sentences populated from 10 random Words per Paragraph
func TestCustomParagraphs(t *testing.T) {
	lorem := dolorem.Ipsum()
	lorem.ParagraphStarter = ""
	result := strings.Split(lorem.Paragraph(5), "\n\n")

	if len(result) != 5 {
		t.Error("@TestCustomParagraphs: Paragraph(5) did not return Paragraph of expected custom number (5)")
	}

	for i := range result {
		if !unicode.IsUpper(rune(result[i][0])) {
			fmt.Println(result[i])
			t.Error("@TestCustomParagraphs: Paragraph(5) did not return Paragraph with expected uppercase first char")
		}

		words := strings.Split(result[i], " ")
		if len(words)/7 != 7 {
			t.Error("@TestCustomParagraphs: Paragrah(5) did not return expected default number of Sentences per Paragraph")
		}
	}

	result = strings.Split(lorem.Paragraph(2, 5), "\n\n")

	if len(result) != 2 {
		t.Error("@TestCustomParagraphs: Paragraph(2, 5) did not return Paragraph of expected custom number (2)")
	}

	for i := range result {
		if !unicode.IsUpper(rune(result[i][0])) {
			t.Error("@TestCustomParagraphs: Paragraph(2, 5) did not return Paragraph with expected uppercase first char")
		}

		words := strings.Split(result[i], " ")
		if len(words)/5 != 7 {
			t.Error("@TestCustomParagraphs: Paragraph(2, 5) did not return Paragraph with expect custom number of Sentences per Paragraph (5)")
		}
	}

	result = strings.Split(lorem.Paragraph(3, 5, 10), "\n\n")

	if len(result) != 3 {
		t.Error("@TestCustomParagraphs: Paragraph(3, 5, 10) did not return Paragraph of expected custom number (3)")
	}

	for i := range result {
		if !unicode.IsUpper(rune(result[i][0])) {
			t.Error("@TestCustomParagraphs: Paragraph(2, 5) did not return Paragraph with expected uppercase first char")
		}

		words := strings.Split(result[i], " ")
		if len(words)/5 != 10 {
			t.Error("@TestCustomParagraphs: Paragraph(3, 5, 10) did not return Paragraph with expected custom number of Sentences and Word count (5, 10)")
		}
	}

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
