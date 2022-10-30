package constant

type EnvironmentType string

var (
	DEV  = EnvironmentType("dev")
	STAG = EnvironmentType("stag")
	PROD = EnvironmentType("prod")
	TEST = EnvironmentType("test")
)
