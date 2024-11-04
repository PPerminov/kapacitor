package snmptrap

import (
	"errors"
)

type Config struct {
	// Common config
	// Whether Snmptrap is enabled.
	Enabled bool `toml:"enabled" override:"enabled"`
	// SNMP version to use. (The default is 2)
	Version int
	// The host:port address of the SNMP trap server
	Addr string `toml:"addr" override:"addr"`
	// Network
	Protocol string `toml:"proto" override:"proto"`
	// Retries count for traps (The default is 3)
	Retries uint `toml:"retries" override:"retries"`
	// Timeout in seconds (The default is 1)
	Timeout int `toml:"timeout" override:"timeout"`
	// Message max size (The default is `1400`)
	MessageMaxSize int `toml:"msgmaxsize" override:"msgmaxsize"`

	// SNMP v1 and v2 specific config
	// SNMP Community
	Community string `toml:"community" override:"community,redact"`

	// SNMPv3 specific config
	// Username name. Required
	UserName string `toml:"username" override:"username"`
	// Security level (NoAuthNoPriv, AuthNoPriv, AuthPriv). Required.
	SecurityLevel string `toml:"security" override:"security"`
	// Authentication protocol pass phrase. Required if AuthNoPriv or AuthPriv selected as Security level.
	AuthPassword string `toml:"authpassword" override:"authpassword"`
	// Authentication protocol. Required if AuthNoPriv or AuthPriv selected as Security level.
	AuthProtocol string `toml:"authproto" override:"authproto"`
	// Privacy protocol pass phrase. Required if AuthPriv selected as Security level.
	PrivPassword string `toml:"privpassword" override:"privpassword"`
	// Privacy protocol. Required if AuthPriv selected as Security level.
	PrivProtocol string `toml:"privproto" override:"privproto"`

	// Optional SNMPv3 zone
	// Security engine ID.
	SecurityEngineId string `toml:"securityengineid" override:"securityengineid"`
	// Context engine ID (V3 specific)
	ContextEngineId string `toml:"contextengineid" override:"contextengineid"`
	// Context name (V3 specific)
	ContextName string `toml:"contextname" override:"contextname"`
}

func NewConfig() Config {
	return Config{
		Addr:      "localhost:162",
		Community: "kapacitor",
		Retries:   1,
	}
}

func (c Config) Validate() error {
	if c.Enabled && c.Addr == "" {
		return errors.New("must specify addr")
	}
	return nil
}
