package constant

type EnvironmentType string

var (
	DEV  = EnvironmentType("DEV")
	STAG = EnvironmentType("STAG")
	PROD = EnvironmentType("PROD")
	TEST = EnvironmentType("TEST")
)
