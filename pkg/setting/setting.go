package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

// App ...
type App struct {
	JwtSecret string
	PageSize  int
	PrefixURL string

	RuntimeRootPath string
	TemplatePath    string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// AppSetting ...
var AppSetting = &App{}

var cfg *ini.File

// Server ...
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting ...
var ServerSetting = &Server{}

// Database ...
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting ...
var DatabaseSetting = &Database{}

// Instagram ...
type Instagram struct {
	APIURL string
}

// InstagramSetting ...
var InstagramSetting = &Instagram{}

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
