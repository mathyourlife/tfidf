use 'godoc cmd/github.com/mathyourlife/tfidf/tfidf' for documentation on the github.com/mathyourlife/tfidf/tfidf command 

PACKAGE DOCUMENTATION

package tfidf
    import "github.com/mathyourlife/tfidf/tfidf"


TYPES

type Corpus struct {
    Docs  []*Document
    Words map[string]*Word
    // contains filtered or unexported fields
}

func NewCorpus() *Corpus

func (c *Corpus) AddDocument(doc *Document)

func (c *Corpus) IDF() map[*Word]float64

func (c *Corpus) Word(text string) *Word

type Document struct {
    Text        []byte
    Words       map[string]*Word
    WordCount   map[*Word]int
    NumberWords int
    // contains filtered or unexported fields
}

func NewDocument(text []byte, corpus *Corpus) (*Document, error)

func (doc *Document) TF() map[*Word]float64

type SortList []TFIDF

func RankTFIDF(tfidf map[*Word]float64) SortList

func (p SortList) Len() int

func (p SortList) Less(i, j int) bool

func (p SortList) Swap(i, j int)

type TFIDF struct {
    Key   *Word
    Value float64
}

type Word struct {
    Docs []*Document
    // contains filtered or unexported fields
}

func NewWord(text string) *Word

func (w *Word) String() string


