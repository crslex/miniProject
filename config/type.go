package config

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	RepoName string
	AppName  string
	HostName string
	IP       string
	GRPC     ServerHost `yaml:"grpc"`
	NSQ      NSQ        `yaml:"nsq"`
}

type ServerHost struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type NSQ struct {
	NSQLookupd string `yaml:"nsqlookupd_address"`
}
