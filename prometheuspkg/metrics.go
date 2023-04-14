package prometheuspkg

import (
	"github.com/prometheus/client_golang/prometheus"
)

type (
	// metricsContainer 容器结构体，用于定义指标的名称、帮助文本和标签
	metricsContainer struct {
		name   string
		help   string
		labels []string
	}

	// PrometheusGaugeData 结构体，用于包装 Prometheus Gauge 对象
	PrometheusGaugeData struct {
		prometheus.Gauge
	}
	// PrometheusGaugeVecData 结构体，用于包装 Prometheus GaugeVec 对象
	PrometheusGaugeVecData struct {
		*prometheus.GaugeVec
	}
)

func NewMetrics(name string) *metricsContainer {
	return &metricsContainer{
		name: name,
	}
}

func (p *metricsContainer) SetHelp(help string) *metricsContainer {
	p.help = help
	return p
}

func (p *metricsContainer) SetLabels(labels ...string) *metricsContainer {
	p.labels = labels
	return p
}

func (p *metricsContainer) BuildGauge() *PrometheusGaugeData {
	return &PrometheusGaugeData{prometheus.NewGauge(prometheus.GaugeOpts{Name: p.name, Help: p.help})}
}

func (p *metricsContainer) BuildGaugeVec() *PrometheusGaugeVecData {
	return &PrometheusGaugeVecData{prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: p.name, Help: p.help}, p.labels)}
}

func (p *PrometheusGaugeData) Data() prometheus.Gauge {
	return p.Gauge
}

func (p *PrometheusGaugeVecData) Data() *prometheus.GaugeVec {
	return p.GaugeVec
}

func (p *PrometheusGaugeData) Set(data float64) {
	p.Gauge.Set(data)
}

func (p *PrometheusGaugeVecData) Set(data float64, labels map[string]string) {
	p.GaugeVec.With(labels).Set(data)
}
