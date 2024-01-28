package config

type ServerConfig struct {
	Server    ServerInfo
	System    SystemInfo
	Authority AuthorityInfo
}

type ServerInfo struct {
	IP   string
	Port int
}

type SystemInfo struct {
	Chroot string
}

type AuthorityInfo struct {
	Password string
}
