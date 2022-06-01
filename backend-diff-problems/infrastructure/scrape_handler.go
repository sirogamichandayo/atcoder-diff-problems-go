package infrastructure

import (
	"bytes"
	"diff-problems/interfaces/web"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
)

type ScrapeHandler struct{}

func NewScrapeHandler() web.ScrapeHandler {
	return &ScrapeHandler{}
}

func (s ScrapeHandler) NewDocument(url string) (web.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	det := chardet.NewTextDetector()
	detResult, err := det.DetectBest(buf)
	if err != nil {
		return nil, err
	}
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detResult.Charset, bReader)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	return Document{Doc: doc}, nil
}

type Document struct {
	Doc *goquery.Document
}

func (d Document) Find(selector string) web.Selection {
	return Selection{Sel: d.Doc.Find(selector)}
}

func (d Document) Text() string {
	return d.Doc.Text()
}

func (d Document) Attr(attr string) (string, bool) {
	return d.Doc.Attr(attr)
}

type Selection struct {
	Sel *goquery.Selection
}

func (s Selection) Find(selector string) web.Selection {
	return Selection{Sel: s.Sel.Find(selector)}
}

func (s Selection) Text() string {
	return s.Sel.Text()
}

func (s Selection) Attr(attr string) (string, bool) {
	return s.Sel.Attr(attr)
}
