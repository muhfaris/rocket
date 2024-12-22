package hexagonal

type RouterGroup struct {
	GroupName string
	GroupPath string
	Routes    []ChildRouterGroup
}

type ChildRouterGroup struct {
	Method  string
	Path    string
	Handler string
}
