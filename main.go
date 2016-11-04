package main

import (
	"bufio"
	"fmt"
	"log"
	"io"
	"os"

	"github.com/mathyourlife/tfidf/tfidf"
)

func main() {
	corpus := tfidf.NewCorpus()

	bio := bufio.NewReader(os.Stdin)

	for {
		fullLine := []byte{}
		line, hasMoreInLine, err := bio.ReadLine()
		fullLine = append(fullLine, line...)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for {
			if !hasMoreInLine {
				break
			}
			line, hasMoreInLine, err = bio.ReadLine()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fullLine = append(fullLine, line...)
		}
		tfidf.NewDocument(fullLine, corpus)

	}

	idf := corpus.IDF()

	for _, doc := range corpus.Docs {
		tf := doc.TF()
		scores := map[*tfidf.Word]float64{}
		for _, w := range doc.Words {
			scores[w] = idf[w] * tf[w]
		}
		sorted := tfidf.RankTFIDF(scores)
		for i := 0; i < len(sorted); i++ {
			fmt.Printf("%f\t%s\t", sorted[i].Value, sorted[i].Key)
		}
		fmt.Println()
	}
}
