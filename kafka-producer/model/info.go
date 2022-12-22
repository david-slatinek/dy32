package model

type Info interface {
	Json() ([]byte, error)
	GetID() string
}
