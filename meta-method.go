package meta

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	pb "buf.build/gen/go/ctfhub/meta/protocolbuffers/go"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	TaskIDSpec = `[^a-z0-9][a-z0-9_\-]{12,94}[a-z0-9]$`
	ImageSpec  = TaskIDSpec
)

func ParseFormatTask(t *pb.Task) *pb.Task {
	if t.Type == "" && t.TypeCode > 0 {
		t.Type = strings.ToLower(pb.Task_Type_name[int32(t.TypeCode)])
	}
	switch strings.ToLower(t.Type) {
	case "con", "web", "http", "pwn", "nc", "tcp", "udp":
		t.Type = "con"
		t.TypeCode = pb.Task_Con
	case "file", "attachment":
		t.Type = "file"
		t.TypeCode = pb.Task_File
	case "ext":
		t.Type = "ext"
		t.TypeCode = pb.Task_Ext
	default:
		t.Type = "file"
		t.TypeCode = pb.Task_File
	}
	// Level
	if t.Level == "" && t.LevelCode > 0 {
		t.Level = strings.ToLower(pb.Task_Level_name[int32(t.LevelCode)])
	}
	switch strings.ToLower(t.Level) {
	case "签到", "checkin", "1":
		t.LevelCode = pb.Task_Checkin
	case "简单", "初级", "easy", "2":
		t.LevelCode = pb.Task_Easy
	case "中等", "中级", "medium", "3":
		t.LevelCode = pb.Task_Medium
	case "困难", "高级", "hard", "4":
		t.LevelCode = pb.Task_Hard
	default:
		t.Level = "easy"
		t.LevelCode = pb.Task_Easy
	}
	return t
}

func VerifyTask(t *pb.Task) error {
	if t.Id == "" {
		return errors.New("task id is required")
	}
	if !regexp.MustCompile(TaskIDSpec).MatchString(t.Id) {
		return errors.New("task id is invalid. " + TaskIDSpec)
	}
	return nil
}

func VerifyContainer(c *pb.Container) error {
	if c.Image == "" {
		return errors.New("image is required")
	}
	// Image name
	if !regexp.MustCompile(ImageSpec).MatchString(c.Image) {
		return errors.New("image is invalid. " + ImageSpec)
	}
	// Resource
	_, err := resource.ParseQuantity(c.Resource.Cpu)
	if err != nil {
		return errors.New("resource cpu is invalid")
	}
	_, err = resource.ParseQuantity(c.Resource.Mem)
	if err != nil {
		return errors.New("resource mem is invalid")
	}
	return nil
}

func Verify(t *pb.Meta) error {
	if t.Task == nil {
		return errors.New("task is required")
	}
	if err := VerifyTask(t.Task); err != nil {
		return err
	}
	if len(t.Containers) > 0 {
		proxyMap := map[string]struct{}{}
		for _, c := range t.Containers {
			if c == nil {
				return errors.New("container is required")
			}
			if err := VerifyContainer(c); err != nil {
				return err
			}
			// Ports
			for _, port := range c.Ports {
				p1, p2, ok := strings.Cut(port, "/")
				if !ok {
					return errors.New("port is invalid. port/protocol")
				}
				if !slices.Contains([]string{"tcp", "udp", "http"}, p2) {
					return errors.New("port protocol is invalid. tcp/udp/http")
				}
				if _, ok := proxyMap[p1]; ok {
					return errors.New("port is duplicate")
				}
				proxyMap[p1] = struct{}{}
			}
		}
	}
	return nil
}

func ParseFormatMeta(m *pb.Meta) *pb.Meta {
	if m.Task != nil {
		ParseFormatTask(m.Task)
	}
	return m
}
