module goapp

go 1.12

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.38.0
	goapp/internal => ./internal
	goapp/pkg => ./pkg
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190429183610-475c5042d3f1
	golang.org/x/image => github.com/golang/image v0.0.0-20190501045829-6d32002ffd75
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190415191353-3e0bab5405d6
	golang.org/x/net => github.com/golang/net v0.0.0-20190415214537-1da14a5a36f2
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190416152802-12500544f89f
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190503185657-3b6f9c0030f7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.4.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190502173448-54afdca5d873
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)

require (
	github.com/jinzhu/gorm v1.9.8
	goapp/internal v0.0.0-00010101000000-000000000000
	goapp/pkg v0.0.0-00010101000000-000000000000
)
