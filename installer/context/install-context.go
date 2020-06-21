package context

type InstallContext struct {
	variables map[string]string
}

func New() *InstallContext {
	return &InstallContext{
		variables: make(map[string]string),
	}
}

func (c InstallContext) SetVar(name string, value string) {
	c.variables[name] = value
}

func (c InstallContext) GetVar(name string) string {
	return c.variables[name]
}
