package tracer

import (
	"flag"
	"fmt"

	"github.com/200Lab-Education/go-sdk/logger"
	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

type jaeger struct {
	processName string
	tracingRate float64
	agentURI    string
	port        int
	stopChan    chan bool
	logger      logger.Logger
}

func NewJaeger(processName string) *jaeger {
	return &jaeger{
		processName: processName,
		stopChan:    make(chan bool),
	}
}

func (j *jaeger) Name() string {
	return "jaeger"
}

func (j *jaeger) GetPrefix() string {
	return j.Name()
}

func (j *jaeger) Get() interface{} {
	return nil
}

func (j *jaeger) InitFlags() {
	flag.Float64Var(
		&j.tracingRate,
		"jaeger-tracing-rate",
		1.0,
		"Sample tracing rate from OpenSensus: 0.0 -> 1.0 (default is 1.0)",
	)

	flag.StringVar(
		&j.agentURI,
		"jaeger-agent-uri",
		"",
		"Jaeger agent URI to receive tracing data directly",
	)

	flag.IntVar(
		&j.port,
		"jaeger-agent-port",
		6831,
		"Jaeger agent port",
	)
}

func (j *jaeger) isEnabled() bool {
	return j.agentURI != ""
}

func (j *jaeger) traceConfig() trace.Config {
	if j.tracingRate >= 1 {
		return trace.Config{
			DefaultSampler: trace.AlwaysSample(),
		}
	}

	return trace.Config{
		DefaultSampler: trace.ProbabilitySampler(j.tracingRate),
	}
}

func (j *jaeger) Configure() error {
	if !j.isEnabled() {
		return nil
	}

	url := fmt.Sprintf("%s:%d", j.agentURI, j.port)

	je, err := jg.NewExporter(jg.Options{
		AgentEndpoint: "localhost:6831",
		Process: jg.Process{
			ServiceName: j.processName,
		},
	})

	if err != nil {
		return err
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(j.traceConfig())

	j.logger = logger.GetCurrent().GetLogger(j.Name())
	j.logger.Infof("Connecting tracer (%s) on %s", j.Name(), url)

	return nil
}

func (j *jaeger) Run() error {
	return j.Configure()
}

func (j *jaeger) Stop() <-chan bool {
	go func() {
		if !j.isEnabled() {
			j.stopChan <- true
			return
		}

		j.stopChan <- true
	}()

	return j.stopChan
}
