package fern91

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FERN_PLUGIN",
	MagicCookieValue: "fern_plugin_interface_v1",
}