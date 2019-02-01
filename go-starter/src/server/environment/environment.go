package environment

type Environment struct {
	Identifier int
	Port       int
	PortForGin string
	RootPath string
}

const (
	Test = iota
	Development
	Production
)

const (
	TestIdentifierName        = "Test"
	DevelopmentIdentifierName = "Development"
	ProductionIdentifierName  = "Production"
	IdentifierVariableKey     = "environment-identifier"
)
