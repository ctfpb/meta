package meta

import (
	"strings"

	"google.golang.org/protobuf/proto"
)

func (m *Meta) R() *Meta { return proto.Clone(m).(*Meta) }

func (t *Task) ParseFormat() *Task {
	switch strings.ToLower(t.Type) {
	case "con","web", "http", "pwn", "nc", "tcp", "udp":
		t.Type = "con"
		t.TypeCode = TaskType_Con
	case "file", "attachment":
		t.Type = "file"
		t.TypeCode = TaskType_File
	case "ext":
		t.Type = "ext"
		t.TypeCode = TaskType_Ext
	default:
		t.Type = "file"
		t.TypeCode = TaskType_File
	}
	// Level
	switch strings.ToLower(t.Level) {
	case "签到", "checkin", "1":
		t.LevelCode = TaskLevel_Checkin
	case "简单", "初级", "easy", "2":
		t.LevelCode = TaskLevel_Easy
	case "中等", "中级", "medium", "3":
		t.LevelCode = TaskLevel_Medium
	case "困难", "高级", "hard", "4":
		t.LevelCode = TaskLevel_Hard
	default:
		t.Level = "easy"
		t.LevelCode = TaskLevel_Easy
	}
	return t
}

// Skill

func (m *Meta) NewSkill(id, pid, tid string, leaf int64, image, name string, level TaskLevel) *Meta {
	n := m.R()
	n.Skill = &Skill{Id: id, Pid: pid, Tid: tid, Leaf: leaf}
	n.Challenge = &Challenge{Name: name}
	n.Task = &Task{Name: image, LevelCode: level}
	return n
}

func (t *Meta) ParseFormat() *Meta {
	t.Task.ParseFormat()
	return t
}
