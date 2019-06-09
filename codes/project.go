package codes

type ProjectArchitect string

const (
	ProjectSingleProcess ProjectArchitect = "SingleProcess"
	ProjectSingleGateway                  = "SingleGateway"
	ProjectMultiSystem                    = "MultiSystem"
	ProjectCommandLine                    = "CommandLine"
)

type Project interface {
	GetArchitect() ProjectArchitect
	GetName() string
	GetDescribe() string
	ClassName() string
}
