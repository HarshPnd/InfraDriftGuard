package build

var env = "dev"

var enableUsageReporting = "true"

type BuildInterface interface {
	IsRelease() bool
	IsUsageReportingEnabled() bool
}

type Build struct{}

func (b Build) IsRelease() bool {
	return env == "release"
}

func (b Build) IsUsageReportingEnabled() bool {
	return b.IsRelease() && enableUsageReporting == "true"
}
