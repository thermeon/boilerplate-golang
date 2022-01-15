package abstractions

type Container struct {
	Repositories
	Services
}

func NewContainer(r Repositories, s Services) *Container {
	return &Container{
		Repositories: r,
		Services:     s,
	}
}

type Repositories struct {
}

type Services struct {
	BackendSelector Selector
}
