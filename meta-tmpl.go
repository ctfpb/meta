package meta

import (
	"encoding/json"

	pb "buf.build/gen/go/ctfhub/meta/protocolbuffers/go"
)

func varPtr[T any](v T) *T { return &v }

func Template() string {
	m := Default()
	m.Task = &pb.Task{
		Id:            "challenge_game_2024_web_abc",
		Name:          "Web题目",
		Type:          "web",
		Category:      "Web",
		Description:   "这是一个模板",
		Level:         "easy",
		AttachmentUrl: varPtr("http://xxxxxxx.xxxx.xx/xxx.zip"),
		Refer:         varPtr(""),
		Flag:          varPtr("ctftrain{this_is_a_test_flag}"),
		Tags: []string{
			"web",
			"2024",
		},
		Egress: varPtr(true),
	}
	m.Containers = []*pb.Container{
		{
			Image: "nginx",
			Ports: []string{"80/tcp"},
			Resource: &pb.Resource{
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
