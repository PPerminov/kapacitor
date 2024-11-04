package snmptrap

import (
	"strings"

	"github.com/k-sone/snmpgo"
)

func SecurityLevels(s string) snmpgo.SecurityLevel {
	switch strings.ToLower(s) {
	case "noauthnopriv":
		return snmpgo.NoAuthNoPriv
	case "authnopriv":
		return snmpgo.AuthNoPriv
	case "authpriv":
		return snmpgo.AuthPriv
	}
	return 0
}

func PrivProtocols(s string) snmpgo.PrivProtocol {
	switch strings.ToLower(s) {
	case "des":
		return snmpgo.Des
	case "aes":
		return snmpgo.Aes
	}
	return ""
}

func AuthProtocols(s string) snmpgo.AuthProtocol {
	switch strings.ToLower(s) {
	case "md5":
		return snmpgo.Md5
	case "sha":
		return snmpgo.Sha
	}
	return ""
}
