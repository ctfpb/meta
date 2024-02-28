package meta

import (
	"errors"
	"regexp"
	"strings"

	"google.golang.org/protobuf/proto"
)

const (
	TaskIDSpec = `[^a-z0-9][a-z0-9_]{12,94}[a-z0-9]$`
)

func (m *Meta) R() *Meta { return proto.Clone(m).(*Meta) }

func (t *Task) ParseFormat() *Task {
	switch strings.ToLower(t.Type) {
	case "con", "web", "http", "pwn", "nc", "tcp", "udp":
		t.Type = "con"
		t.TypeCode = Task_Con
	case "file", "attachment":
		t.Type = "file"
		t.TypeCode = Task_File
	case "ext":
		t.Type = "ext"
		t.TypeCode = Task_Ext
	default:
		t.Type = "file"
		t.TypeCode = Task_File
	}
	// Level
	switch strings.ToLower(t.Level) {
	case "签到", "checkin", "1":
		t.LevelCode = Task_Checkin
	case "简单", "初级", "easy", "2":
		t.LevelCode = Task_Easy
	case "中等", "中级", "medium", "3":
		t.LevelCode = Task_Medium
	case "困难", "高级", "hard", "4":
		t.LevelCode = Task_Hard
	default:
		t.Level = "easy"
		t.LevelCode = Task_Easy
	}
	return t
}

func (t *Task) Check() error {
	if t.Id == "" {
		return errors.New("task id is required")
	}
	if !regexp.MustCompile(TaskIDSpec).MatchString(t.Id) {
		return errors.New("task id is invalid. " + TaskIDSpec)
	}
	return nil
}

func (t *Meta) Check() error {
	if t.Task == nil {
		return errors.New("task is required")
	}
	if err := t.Task.Check(); err != nil {
		return err
	}
	return nil
}

func (t *Meta) ParseFormat() *Meta {
	if t.Task != nil {
		t.Task.ParseFormat()
	}
	return t
}
