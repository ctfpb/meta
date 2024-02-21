package meta

import (
	"strings"

	"google.golang.org/protobuf/proto"
)

func (m *Meta) R() *Meta { return proto.Clone(m).(*Meta) }

func (t *Task) parseFormat() *Task {
	switch strings.ToLower(t.Type) {
	case "con", "container", "web", "pwn":
		t.Type = "con"
		t.TypeCode = TaskType_Con
	case "file":
		t.Type = "file"
		t.TypeCode = TaskType_File
	case "ext":
		t.Type = "ext"
		t.TypeCode = TaskType_Ext
	default:
		// 默认 容器 con
		t.Type = "con"
		t.TypeCode = TaskType_Con
	}
	// Level
	switch strings.ToLower(t.Level) {
	case "签到", "checkin":
		t.LevelCode = TaskLevel_Checkin
	case "简单", "初级", "easy":
		t.LevelCode = TaskLevel_Easy
	case "中等", "中级", "medium":
		t.LevelCode = TaskLevel_Medium
	case "困难", "高级", "hard":
		t.LevelCode = TaskLevel_Hard
	default:
		// 默认 简单easy
		t.Level = "checkin"
		t.LevelCode = TaskLevel_Checkin
	}
	return t
}

// Skill

func (m *Meta) NewSkill(id, pid, tid string, leaf int64, image, name string, level TaskLevel) *Meta {
	n := m.R()
	n.Skill.Id = id
	n.Skill.Pid = pid
	n.Skill.Tid = tid
	n.Skill.Leaf = leaf
	n.Challenge.Name = name
	n.Task.Name = image
	n.Task.LevelCode = level
	return n
}

// Deploy

func (t *Meta) parseFormat() *Meta {
	t.Task.parseFormat()
	return t
}
