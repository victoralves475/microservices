module github.com/victoralves475/microservices/payment

go 1.24

require (
    // stub do Payment
    github.com/victoralves475/microservices-proto/golang/payment v0.0.0

    // demais dependências...
    github.com/sirupsen/logrus                         v1.9.0
    github.com/uptrace/opentelemetry-go-extra/otelgorm v0.3.2
    go.opentelemetry.io/otel                           v1.36.0
    go.opentelemetry.io/otel/exporters/jaeger          v1.17.0
    go.opentelemetry.io/otel/sdk                       v1.36.0
    go.opentelemetry.io/otel/trace                     v1.36.0
    google.golang.org/grpc                             v1.74.2
    gorm.io/driver/mysql                               v1.6.0
    gorm.io/gorm                                       v1.30.1
)

replace github.com/victoralves475/microservices-proto/golang/payment => ../../microservices-proto/golang/payment
