package meta

import "encoding/json"

func StrPtr(s string) *string { return &s }

func Template() string {
	m := Default()
	m.Task = &Task{
		Id:            "challenge_game_2024_web_abc",
		Name:          "Web题目",
		Type:          "web",
		Category:      "Web",
		Description:   "这是一个模板",
		Level:         "easy",
		AttachmentUrl: StrPtr("http://xxxxxxx.xxxx.xx/xxx.zip"),
		Refer:         StrPtr(""),
		Flag:          StrPtr("ctftrain{this_is_a_test_flag}"),
		Tags: []string{
			"web",
			"2024",
		},
		BanEgress: true,
	}
	m.Containers = []*Container{
		{
			Image: "nginx",
			Ports: []string{"80/tcp"},
			Resource: &Resource{
				Cpu: "500m",
				Mem: "512Mi",
			},
		},
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
