//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package controllers

type Context interface {
	Param(string) string
	DefaultQuery(string, string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
