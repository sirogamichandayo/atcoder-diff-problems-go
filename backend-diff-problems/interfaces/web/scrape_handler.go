//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package web

type ScrapeHandler interface {
	NewDocument(url string) (Document, error)
}

type Document interface {
	Find(selector string) Selection
	Text() string
	Attr(string) (string, bool)
}

type Selection interface {
	Find(selector string) Selection
	Text() string
	Attr(string) (string, bool)
}
