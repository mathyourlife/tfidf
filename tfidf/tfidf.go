package tfidf

import (
	"bytes"
	"math"
	"sort"
)

type Word struct {
	text string
	Docs []*Document
}

func NewWord(text string) *Word {
	w := &Word{
		text: text,
		Docs: []*Document{},
	}
	return w
}

func (w *Word) String() string {
	return string(w.text)
}

type Document struct {
	Text        []byte
	Words       map[string]*Word
	WordCount   map[*Word]int
	NumberWords int
	corpus      *Corpus
	tf          map[*Word]float64
}

func NewDocument(text []byte, corpus *Corpus) (*Document, error) {
	doc := &Document{
		Text:        text,
		Words:       map[string]*Word{},
		WordCount:   map[*Word]int{},
		NumberWords: 0,
		corpus:      corpus,
	}
	corpus.AddDocument(doc)
	doc.readDoc()
	doc.TF()
	return doc, nil
}

func (doc *Document) TF() map[*Word]float64 {
	if doc.tf != nil {
		return doc.tf
	}
	tf := map[*Word]float64{}
	total := float64(doc.NumberWords)
	for w, c := range doc.WordCount {
		tf[w] = float64(c) / total
	}
	doc.tf = tf
	return doc.tf
}

func (doc *Document) readDoc() {
	words := bytes.Split(doc.Text, []byte(" "))
	for _, wordBytes := range words {
		wordText := string(wordBytes)
		doc.NumberWords++
		w := doc.corpus.Word(wordText)
		doc.WordCount[w]++

		exist := doc.Words[wordText]
		if exist != nil {
			continue
		}
		doc.Words[wordText] = w
		w.Docs = append(w.Docs, doc)
	}
}

type Corpus struct {
	Docs  []*Document
	Words map[string]*Word
	idf   map[*Word]float64
}

func NewCorpus() *Corpus {
	c := &Corpus{
		Docs:  []*Document{},
		Words: map[string]*Word{},
	}
	return c
}

func (c *Corpus) Word(text string) *Word {
	w := c.Words[text]
	if w == nil {
		w = NewWord(text)
		c.Words[text] = w
	}
	return w
}

func (c *Corpus) AddDocument(doc *Document) {
	c.Docs = append(c.Docs, doc)
}

func (c *Corpus) IDF() map[*Word]float64 {
	if c.idf != nil {
		return c.idf
	}
	docCount := float64(len(c.Docs))
	idf := map[*Word]float64{}
	for _, w := range c.Words {
		idf[w] = math.Log(docCount / float64(1+len(w.Docs)))
	}
	c.idf = idf
	return c.idf
}

func RankTFIDF(tfidf map[*Word]float64) SortList {
	pl := make(SortList, len(tfidf))
	i := 0
	for k, v := range tfidf {
		pl[i] = TFIDF{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type TFIDF struct {
	Key   *Word
	Value float64
}

type SortList []TFIDF

func (p SortList) Len() int           { return len(p) }
func (p SortList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p SortList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
