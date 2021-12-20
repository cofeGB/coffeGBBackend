module github.com/cofeGB/coffeGBBackend

go 1.16

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/kelseyhightower/envconfig v1.4.0
)

// heroku deploy specific
// +heroku goVersion go1.16
// +heroku install ./cmd/...

