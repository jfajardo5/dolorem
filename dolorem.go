package main

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed dictionary.txt
var f embed.FS

type Dolorem struct {
	dictionary        []string
	paragraph_starter string
	text              string
	seed              *rand.Rand
}

func main() {
	dol, err := New()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dol)
}

func New() (Dolorem, error) {
	dict, err := loadDictionary()
	if err != nil {
		return Dolorem{}, err
	}
	return Dolorem{
		dictionary:        dict,
		seed:              rand.New(rand.NewSource(time.Now().Unix())),
		paragraph_starter: "Dolorem ipsum dolor sit amet, ",
	}, nil
}

func loadDictionary() ([]string, error) {
	content, err := f.ReadFile("dictionary.txt")
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func (d *Dolorem) SingleWord() string {
	index := d.seed.Intn(len(d.dictionary) - 1)
	return d.dictionary[index]
}

func (d *Dolorem) Sentence() string {
	var sentence string
	for i := 0; i < 15; i++ {
		index := d.seed.Intn(len(d.dictionary) - 1)
		if i > 0 {
			sentence = sentence + " " + d.dictionary[index]
			continue
		}
		sentence = strings.Title(d.dictionary[index])
	}
	return sentence
}

func (d *Dolorem) Paragraph(param ...int) string {
	num := 1
	if param != nil {
		num = param[0]
	}
	var paragraph = d.paragraph_starter
	for i := 0; i < num; i++ {
		for j := 0; j < 7; j++ {
			paragraph = paragraph + " " + d.Sentence()
		}
		paragraph = paragraph + "."
	}
	return paragraph
}
