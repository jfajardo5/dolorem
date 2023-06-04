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
	return paragraph
}

func loadDictionary() []string {
	return strings.Split(dictionary, "\n")
}
