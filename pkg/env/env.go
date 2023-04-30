package env

type Env string

const (
	PROD Env = "prod"
	DEV  Env = "dev"
)

func (e Env) Is(ee Env) bool {
	return e == ee
}

func (e Env) IsProd() bool {
	return e == PROD
}

func (e Env) IsDev() bool {
	return e == DEV
}
