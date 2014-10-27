package consul

import "time"

type Event struct {
	ID            string
	Name          string
	Payload       []byte
	NodeFilter    string
	ServiceFilter string
	TagFilter     string
	Version       uint
	LTime         uint
}

type Check struct {
	Node        string
	CheckID     string
	Name        string
	Status      string
	Notes       string
	Output      string
	ServiceID   string
	ServiceName string
}

type ConsulAlertConfig struct {
	Checks    *ChecksConfig
	Events    *EventsConfig
	Notifiers *NotifiersConfig
}

type ChecksConfig struct {
	Enabled         bool
	ChangeThreshold int
}

type EventsConfig struct {
	Enabled  bool
	Handlers []string
}

type NotifiersConfig struct {
	Email    *EmailNotifierConfig
	Log      *LogNotifierConfig
	Influxdb *InfluxdbNotifierConfig
	Custom   []string
}

type EmailNotifierConfig struct {
	ClusterName string
	Enabled     bool
	Url         string
	Port        int
	Username    string
	Password    string
	SenderAlias string
	SenderEmail string
	Receivers   []string
	Template    string
}

type LogNotifierConfig struct {
	Enabled bool
	Path    string
}

type InfluxdbNotifierConfig struct {
	Enabled    bool
	Host       string
	Username   string
	Password   string
	Database   string
	SeriesName string
}

type Status struct {
	Current          string
	CurrentTimestamp time.Time
	Pending          string
	PendingTimestamp time.Time
	HealthCheck      *Check
	ForNotification  bool
}

type Consul interface {
	LoadConfig()

	EventsEnabled() bool
	ChecksEnabled() bool
	EventHandlers(eventName string) []string

	EmailConfig() *EmailNotifierConfig
	LogConfig() *LogNotifierConfig
	InfluxdbConfig() *InfluxdbNotifierConfig

	CheckChangeThreshold() int
	UpdateCheckData()
	NewAlerts() []Check

	CustomNotifiers() []string

	CheckStatus(node, statusId, checkId string) (status, output string)
}

func DefaultAlertConfig() *ConsulAlertConfig {

	checks := &ChecksConfig{
		Enabled:         true,
		ChangeThreshold: 60,
	}

	events := &EventsConfig{
		Enabled:  true,
		Handlers: []string{},
	}

	email := &EmailNotifierConfig{
		ClusterName: "Consul-Alerts",
		Enabled:     false,
		SenderAlias: "Consul Alerts",
		Receivers:   []string{},
	}

	log := &LogNotifierConfig{
		Enabled: true,
		Path:    "/tmp/consul-notifications.log",
	}

	influxdb := &InfluxdbNotifierConfig{
		Enabled:    false,
		SeriesName: "consul-alerts",
	}

	notifiers := &NotifiersConfig{
		Email:    email,
		Log:      log,
		Influxdb: influxdb,
		Custom:   []string{},
	}

	return &ConsulAlertConfig{
		Checks:    checks,
		Events:    events,
		Notifiers: notifiers,
	}
}