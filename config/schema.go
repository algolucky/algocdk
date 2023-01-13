package config

// spf13/viper uses github.com/mitchellh/mapstructure
// under the hood for unmarshaling values which uses
// mapstructure tags by default.

type config struct {
	Stack        StackConfig    `mapstructure:"stack"`
	Network      NetworkConfig  `mapstructure:"network"`
	Relays       InstanceConfig `mapstructure:"relays"`
	NonPartNodes InstanceConfig `mapstructure:"nonPartNodes"`
	PartNodes    InstanceConfig `mapstructure:"partNodes"`
	Groups       GroupsConfig   `mapstructure:"groups"`
}

// StackConfig contains the stack-wide configuration
type StackConfig struct {
	Name string `mapstructure:"name"`
}

// NetworkConfig contains the network-wide configuration
type NetworkConfig struct {
	Name    string `mapstructure:"name"`
	Wallets string `mapstructure:"wallets"`
}

// InstanceConfig represents the configuration for a specific type of instance
// "relay", "partNode", or "nonPartNode"
type InstanceConfig struct {
	Count    int    `mapstructure:"count"`
	Type     string `mapstructure:"type"`
	DiskSize int    `mapstructure:"disk_size"`
}

type GroupsTopologyConfig struct {
	RelaysPercent       int `mapstructure:"relays"`
	PartNodesPercent    int `mapstructure:"partNodes"`
	NonPartNodesPercent int `mapstructure:"nonPartNodes"`
}

type PlatformConfig struct {
	Region string `mapstructure:"region"`
}

type GroupsConfig struct {
	Name                    string
	Topology                GroupsTopologyConfig `mapstructure:"topology"`
	Platform                string               `mapstructure:"platform"`
	PlatformConfigOverrides PlatformConfig       `mapstructure:"platform_config"`
}

//

type ContainerConfig struct {
	ContainerRepo string
	ContainerTag  string
	AlgodPort     float64
	Network       string
	KmdPort       float64
	StartKMD      bool
	KmdToken      string
	FastCatchup   bool
	TelemetryName string
	Token         string
	AdminToken    string
}
