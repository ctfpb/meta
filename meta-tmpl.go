package meta

import (
	"gopkg.in/yaml.v3"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

func Template() string {
	m := Default()
	m.Task = &Task{
		Name:        "challenge_game_2023_web_abc",
		Type:        "web",
		Category:    "Web",
		Description: "这是一个模板",
		Level:       "easy",
		Flag:        "ctftrain{this_is_a_test_flag}",
		Hints:       []string{"这是一个模板", "没有提示"},
	}
	m.Challenge = &Challenge{
		Name:  "Web题目",
		Refer: "https://www.ctfhub.com",
		Tags:  []string{"web", "2024"},
	}
	m.Containers = []*Container{
		{
			Image: "nginx",
			Ports: []string{"80/tcp"},
			Resource: &Resource{
				Cpu: resource.NewQuantity(500, resource.DecimalSI),
				Mem: resource.NewQuantity(512, resource.BinarySI)},
		},
	}
	buf, err := yaml.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
