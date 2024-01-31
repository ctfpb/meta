package meta

type Deploy struct {
	CPU      int    `yaml:"cpu"`      // -> KubeDeployment
	Mem      int    `yaml:"mem"`      // -> KubeDeployment
	Protocol string `yaml:"protocol"` // TCP || UDP || ALL -> Ingress
}
