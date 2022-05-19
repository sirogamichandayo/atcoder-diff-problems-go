//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package web

type ScrapeHandler interface {
	NewDocument(url string) Document
}

type Document interface {
	Find(selector string) Document
	Text() string
	Attr(string) (string, bool)
}
