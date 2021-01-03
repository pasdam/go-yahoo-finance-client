module github.com/pasdam/go-yahoo-finance-client

replace github.com/pasdam/go-yahoo-finance-client/pkg => ./pkg

go 1.14

require (
	github.com/pasdam/go-rest-util v0.0.0-20201231095731-b8c8cda9c311
	github.com/pasdam/mockit v0.0.0-20201218205457-c54f12cbcca0
	github.com/stretchr/testify v1.6.1
)
