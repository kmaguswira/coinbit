package external_services

type IRedis interface {
	GetValue(key string) (string, error)
	SetValue(key string, value string, duration int) error
}
