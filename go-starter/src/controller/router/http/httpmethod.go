package http

type Method string

const (
	Post 	= Method("Post")
	Put 	= Method("Put")
	Get 	= Method("Get")
	Delete 	= Method("Delete")
	Patch 	= Method("Patch")
)
