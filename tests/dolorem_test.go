package dolorem

import (
	"strings"
	"testing"

	"github.com/jfajardo5/dolorem"
)

func TestDefaultInitializer(t *testing.T) {
	lorem := dolorem.Ipsum()
	if lorem.ParagraphStarter != "dolorem ipsum dolor sit amet, " {
		t.Error("@TestDefaultInitializer: Ipsum() did not return default ParagraphStarter.")
	}
	if len(lorem.Dictionary) < 3086 {
		t.Error("@TestDefaultInitializer: Ipsum() did not return full Dictionary.")
	}
}

func TestWord(t *testing.T) {
	lorem := dolorem.Ipsum()
	lorem.Dictionary = []string{"test", "go"}
	result := lorem.Word()
	if result != "test" && result != "go" {
		t.Error("@TestWord: Word() did not return a value from it's dictionary")
	}
}

func TestSentence(t *testing.T) {
	lorem := dolorem.Ipsum()
	lorem.Dictionary = []string{"test", "go"}
	result := strings.Split(lorem.Sentence(), " ")

	if len(result) != 15 {
		t.Error("@TestSentence: Sentence() did not return Sentence() of expected default length")
	}

	result = strings.Split(lorem.Sentence(5), " ")
	if len(result) != 5 {
		t.Error("@TestSentence: Sentence() did not return Sentence() of expected custom length (5)")
	}
}

func TestParagraph(t *testing.T) {
	lorem := dolorem.Ipsum()
	lorem.Dictionary = []string{"test", "go"}
	result := strings.Split(lorem.Paragraph(), "\n\n")
	if len(result) != 1 {
		t.Error("@TestParagraph: Paragraph() did not return Paragraph of expected default number")
	}
	result = strings.Split(lorem.Paragraph(5), "\n\n")
	if len(result) != 5 {
		t.Error("@TestParagraph: Paragraph(5) did not return Paragraph of expected custom number (5)")
	}

	lorem.ParagraphStarter = ""
	result = strings.Split(lorem.Paragraph(2, 5), "\n\n")
	if len(result) != 2 {
		t.Error("@TestParagraph: Paragraph(2, 5) did not return Paragraph of expected custom number (2)")
	}
	for i := range result {
		words := strings.Split(result[i], " ")
		if len(words)/5 != 15 {
			t.Error("@TestParagraph: Paragraph(2, 5) did not return Paragraph Sentences of expected custom word count(5)")
		}
	}

}
