module longtest

go 1.18

replace k8s.io/client-go v12.0.0+incompatible => k8s.io/client-go v0.22.1

require (
	github.com/akvlad/flog v0.4.4-0.20220607095327-09ef70e4099e
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4
	github.com/gorilla/websocket v1.5.0
	github.com/influxdata/influxdb-client-go/v2 v2.12.2
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf
	github.com/openzipkin/zipkin-go v0.4.1
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/prometheus v0.42.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.13
	go.opentelemetry.io/proto/otlp v0.19.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/brianvoe/gofakeit v3.18.0+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/deepmap/oapi-codegen v1.12.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.39.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)
