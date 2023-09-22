package generator

type RootOutput struct {
	OutputBasedir        string
	ImportPackages       map[string]string
	Providers            []string
	InjectedComponents   []string
	InjectedControllers  []string
	RoutesCode           []string
	RunGopringParameters []string
}
