package gox

type DeviceType string

func (d DeviceType) IsValid() bool {
	switch d {
	case IOS, Android, MacOS, Windows:
		return true
	default:
		return false
	}
}

const (
	IOS     DeviceType = "ios"
	Android DeviceType = "android"
	MacOS   DeviceType = "mac"
	Windows DeviceType = "windows"
)

type Env string

func (e Env) IsValid() bool {
	switch e {
	case Dev, Testing, Staging, Prod:
		return true
	default:
		return false
	}
}

func (e Env) String() string {
	return string(e)
}

const (
	Dev     Env = "dev"
	Testing Env = "testing"
	Staging Env = "staging"
	Prod    Env = "prod"
)
