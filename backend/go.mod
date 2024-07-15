module backend

go 1.21.5

replace api => ./api

replace pokerLogic => ./poker

require github.com/dgryski/go-pcgr v0.0.0-20211101192959-4b34ab9ccb8c

require (
	github.com/enki/fastprng v0.0.0-20190912035746-b46f877a1a50 // indirect
	github.com/lazybeaver/xorshift v0.0.0-20170702203709-ce511d4823dd // indirect
	github.com/rs/cors v1.11.0 // indirect
)
