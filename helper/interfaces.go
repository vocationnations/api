package helper

type VNDatabaseConfiguration struct {
	Port string
	Name string
	User string
	Pass string
	Host string
}

type ONETConfiguration struct {
	Base        string
	AuthToken   string
	AuthType    string
	ContentType string
	AcceptType  string
}

type VNConfiguration struct {
	DatabaseConfiguration VNDatabaseConfiguration
	APIPort               string
	ONET                  ONETConfiguration
	Version               string
	Env                   string
}
