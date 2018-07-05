package conf

type Config struct {
	BaseUrlConfig BaseUrl
}

type BaseUrl struct {
	BaseDNS    string
	ProductDNS string
}
