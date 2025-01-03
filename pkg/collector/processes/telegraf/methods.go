package telegraf

import (
	"log"
	"time"

	"github.com/influxdata/telegraf"
)

// SentinelAccumulator è un accumulatore per testare la raccolta di metriche.
type SentinelAccumulator struct {
	Metrics []telegraf.Metric
	Errors  []error
}

// AddFields aggiunge un set di campi al SentinelAccumulator.
func (m *SentinelAccumulator) AddFields(_ string, fields map[string]interface{}, tags map[string]string, _ ...time.Time) {
	var metric telegraf.Metric

	for k, v := range fields {
		metric.AddField(k, v)
	}
	for k, v := range tags {
		metric.AddTag(k, v)
	}
	m.AddMetric(metric)
}

// AddGauge implementa l'aggiunta di una metrica Gauge.
func (m *SentinelAccumulator) AddGauge(name string, fields map[string]interface{}, tags map[string]string, ts ...time.Time) {
	m.AddFields(name, fields, tags, ts...)
}

// AddCounter implementa l'aggiunta di una metrica Counter.
func (m *SentinelAccumulator) AddCounter(name string, fields map[string]interface{}, tags map[string]string, ts ...time.Time) {
	m.AddFields(name, fields, tags, ts...)
}

// AddMetric aggiunge una metrica generica.
func (m *SentinelAccumulator) AddMetric(metric telegraf.Metric) {
	m.Metrics = append(m.Metrics, metric)
}

// Debug restituisce false per disabilitare i messaggi di debug.
func (m *SentinelAccumulator) Debug() bool {
	return false
}

// Error stampa un errore.
func (m *SentinelAccumulator) Error(err error) {
	log.Println("Accumulator error:", err)
}

// Altri metodi richiesti dall'interfaccia Accumulator
func (m *SentinelAccumulator) SetPrecision(time.Duration) {}
func (m *SentinelAccumulator) StartFlushTimer()           {}
func (m *SentinelAccumulator) StopFlushTimer()            {}

// contains verifica se un valore esiste in una lista.
func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// safeString restituisce una stringa sicura da una metrica.
func safeString(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}
	return "unknown"
}

// safeFloat restituisce un float sicuro da una metrica.
func safeFloat(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	}
	return 0.0
}

// AddError adds an error to the accumulator.
func (m *SentinelAccumulator) AddError(err error) {
	m.Errors = append(m.Errors, err)
}

func (m *SentinelAccumulator) AddSummary(name string, fields map[string]interface{}, tags map[string]string, ts ...time.Time) {
	m.AddFields(name, fields, tags, ts...)
}

func (m *SentinelAccumulator) AddHistogram(name string, fields map[string]interface{}, tags map[string]string, ts ...time.Time) {
	m.AddFields(name, fields, tags, ts...)
}

func (m *SentinelAccumulator) AddMetricWithTimestamp(ts time.Time, metric telegraf.Metric, tags map[string]string) {
	metric.AddTag("timestamp", ts.Format(time.RFC3339))
	m.AddMetric(metric)
}

func (m *SentinelAccumulator) WithTracking(maxTracked int) telegraf.TrackingAccumulator {
	return m.WithTracking(maxTracked)
}
