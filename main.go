package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Create a new document from the HTML string

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		fmt.Printf("Error creating document: %v\n", err)
		return
	}

	// Modify the text of the paragraph
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, found := s.Attr("href")
		if !found {
			return
		}

		if before, cut := strings.CutSuffix(href, ".md"); cut {
			// special handling of README
			if readmeBefore, readmeCut := strings.CutSuffix(before, "README"); readmeCut {
				before = readmeBefore + "index"
			}

			s.SetAttr("href", before+".html")
		}
	})

	// Print the modified HTML
	modifiedHtml, err := doc.Html()
	if err != nil {
		fmt.Printf("Error generating HTML: %v\n", err)
		return
	}

	fmt.Println(modifiedHtml)
}
