package database

type Database interface {
	GetUser(string) string
}
