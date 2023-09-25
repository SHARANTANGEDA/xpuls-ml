package types

import (
	v2 "k8s.io/api/autoscaling/v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type DeploymentConfiguration struct {
	Server  ServerConfig     `yaml:"server"`
	Logging LoggingConfig    `yaml:"logging"`
	K8s     KubernetesConfig `yaml:"kubernetes"`
}
type ServerConfig struct {
	Name                  string      `yaml:"name"`
	Version               string      `yaml:"version,omitempty"`
	ApplicationConnectors []Connector `yaml:"applicationConnectors"`
	AdminConnectors       []Connector `yaml:"adminConnectors"`
}

type Connector struct {
	Type string `yaml:"type"`
	Port int    `yaml:"port"`
}

type LoggingConfig struct {
	Level     string            `yaml:"level"`
	Loggers   map[string]string `yaml:"loggers"`
	Appenders []Appender        `yaml:"appenders"`
}

type Appender struct {
	Type      string `yaml:"type"`
	Threshold string `yaml:"threshold"`
	LogFormat string `yaml:"logFormat"`
	TimeZone  string `yaml:"timeZone"`
}

type KubernetesConfig struct {
	Requests       map[v1.ResourceName]string            `yaml:"requests,omitempty"`
	Limits         map[v1.ResourceName]resource.Quantity `yaml:"limits,omitempty"`
	Tolerations    []v1.Toleration                       `yaml:"tolerations"`
	NodeSelectors  []map[string]string                   `yaml:"nodeSelectors"`
	Image          string                                `yaml:"image"`
	Args           string                                `yaml:"args"`
	HealthCheck    string                                `yaml:"healthcheck"`
	LivenessProbe  v1.Probe                              `yaml:"livenessProbe,omitempty" `
	ReadinessProbe v1.Probe                              `yaml:"readinessProbe,omitempty"`
	StartupProbe   v1.Probe                              `yaml:"startupProbe,omitempty"`

	Env           []DeploymentEnvs `yaml:"env"`
	ContainerSpec ContainerSpec    `yaml:"containerSpec"`
	Volumes       []v1.Volume      `yaml:"volumes"`
}

type DeploymentEnvs struct {
	Namespace string          `yaml:"namespace"`
	instances int             `yaml:"instances"`
	Cluster   string          `yaml:"cluster,omitempty"`
	Requests  v1.ResourceList `yaml:"requests,omitempty"`
	Limits    v1.ResourceList `yaml:"limits,omitempty"`
	AutoScale HpaConfig       `yaml:"autoscale"`
}
type HpaConfig struct {
	MinReplicas              int32           `yaml:"minReplicas,omitempty"`
	MaxReplicas              int32           `yaml:"maxReplicas"`
	CpuAverageUtilization    int             `yaml:"cpuAverageUtilization,omitempty"`
	MemoryAverageUtilization int             `yaml:"memoryAverageUtilization,omitempty"`
	AdvancedMetrics          []v2.MetricSpec `yaml:"metrics,omitempty"`
}

type ContainerSpec struct {
	Env          []v1.EnvVar      `yaml:"env"`
	VolumeMounts []v1.VolumeMount `yaml:"volumeMounts"`
}
