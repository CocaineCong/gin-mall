package track

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func GetDefaultConfig() *config.Configuration {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	return cfg
}

func InitJaeger() (opentracing.Tracer, io.Closer) {
	cfg := GetDefaultConfig()
	service := "mall"
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	// 设置顶级span
	span := tracer.StartSpan(name)
	return span
}

func WithSpan(ctx context.Context, name string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	return span, ctx
}

func GetCarrier(span opentracing.Span) (opentracing.HTTPHeadersCarrier, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext))
	if err != nil {
		return nil, err
	}
	return parentSpan, err
}
