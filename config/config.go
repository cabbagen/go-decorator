package config

var ApplicationConfig map[string]string = map[string]string {
	"static": "./public",
	"templateDir": "./views",
}

var DatabaseConfig map[string]string = map[string]string {
	"dbname": "cb_cms_server",
	"username": "root",
	"password": "artART5201314??",
}

var CacheConfig map[string]string = map[string]string {
	"addr": "localhost:6379",
	"password": "",
	"db": "0",
}
