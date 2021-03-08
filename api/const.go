package api

const (
	V0Prefix = "/api/v0/"

	// Connections
	GetForwardedPortsPath     = V0Prefix + "connections/forwarded_ports"
	GetInboundConnectionsPath = V0Prefix + "connections/inbound"

	// Peers
	GetKnownPeersPath        = V0Prefix + "peers/get_known"
	GetKnownPeerSettingsPath = V0Prefix + "peers/get_known_peer_settings"

	SendFriendRequestPath    = V0Prefix + "peers/invite_peer"
	AcceptPeerInvitationPath = V0Prefix + "peers/accept_peer"
	UpdatePeerSettingsPath   = V0Prefix + "peers/update_settings"
	GetAuthRequestsPath      = V0Prefix + "peers/auth_requests"

	// Settings
	GetMyPeerInfoPath      = V0Prefix + "settings/peer_info"
	UpdateMyInfoPath       = V0Prefix + "settings/update"
	ExportServerConfigPath = V0Prefix + "settings/export_server_config"

	// Debug
	GetP2pDebugInfoPath = V0Prefix + "debug/p2p_info"
	GetDebugLogPath     = V0Prefix + "debug/log"
)
