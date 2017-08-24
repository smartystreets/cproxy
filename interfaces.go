package cproxy

import (
	"io"
	"net/http"
)

type (
	Filter interface {
		IsAuthorized(*http.Request) bool
	}

	ClientConnector interface {
		Connect(w http.ResponseWriter) Socket
	}

	ServerConnector interface {
		Connect(Socket, string) Proxy
	}

	Proxy interface {
		Proxy()
	}
)


type (
	Socket interface {
		io.ReadWriteCloser
	}
	TCPSocket interface {
		Socket
		CloseRead() error
		CloseWrite() error
	}
)