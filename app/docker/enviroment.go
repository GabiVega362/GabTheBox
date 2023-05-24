package docker

type Enviroment struct {
	Image string

	Title       string
	Description string
	Deployed    bool

	Port int
	User string
}

func NewEnviroment(image string, title string, desc string) *Enviroment {
	return &Enviroment{
		Image:       image,
		Title:       title,
		Description: desc,
		Deployed:    false,
		Port:        0,
		User:        "",
	}

}
