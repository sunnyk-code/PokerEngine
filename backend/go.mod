module backend

go 1.21.5

replace api => ./api

replace pokerLogic => ./pokerLogic

require api v0.0.0-00010101000000-000000000000

require (
	github.com/dgryski/go-pcgr v0.0.0-20211101192959-4b34ab9ccb8c // indirect
	github.com/labstack/echo/v4 v4.11.4 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	pokerLogic v0.0.0-00010101000000-000000000000 // indirect
)
