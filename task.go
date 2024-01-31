package meta

import "strings"

type TaskLevel int32

const (
	LevelCheckin TaskLevel = iota + 1
	LevelEasy
	LevelMedium
	LevelHard
)

type TaskType int32

const (
	TypeWeb  TaskType = iota + 1 // 浏览器访问
	TypeNc                       // NC 访问
	TypeFile                     // 纯附件
	TypeExt                      // 外部题目
)

type Task struct {
	Name          string    `yaml:"name"`                // 镜像名称
	Type          string    `yaml:"type"`                // 镜像类型
	TypeCode      TaskType  `yaml:"type_code,omitempty"` // 题目分类
	Category      string    `yaml:"category"`            // 题目分类
	Description   string    `yaml:"description"`
	Level         string    `yaml:"level"`
	LevelCode     TaskLevel `yaml:"level_code,omitempty"`
	Flag          string    `yaml:"flag"`
	AttachmentURL string    `yaml:"attachment_url"`
	Hints         []string  `yaml:"hints"`
}

func (t *Task) parseFormat() *Task {
	// Type 默认 web
	if t.Type != "" {
		t.Type = strings.ToLower(t.Type)
	} else {
		t.Type = "web"
	}
	// Level 默认 简单easy
	switch strings.ToLower(t.Level) {
	case "签到", "checkin":
		t.LevelCode = LevelCheckin
	case "简单", "easy":
		t.LevelCode = LevelEasy
	case "中等", "中级", "medium":
		t.LevelCode = LevelMedium
	case "困难", "高级", "hard":
		t.LevelCode = LevelHard
	default:
		t.Level = "easy"
		t.LevelCode = LevelEasy
	}
	return t
}

func (t *Task) Format() *Task {
	// Level
	if t.LevelCode != 0 && t.Level == "" {
		switch t.LevelCode {
		case LevelCheckin:
			t.Level = "签到"
		case LevelEasy:
			t.Level = "简单"
		case LevelMedium:
			t.Level = "中等"
		case LevelHard:
			t.Level = "困难"
		default:
		}
		t.LevelCode = 0
	} else {
		switch strings.ToLower(t.Level) {
		case "签到", "checkin":
			t.LevelCode = LevelCheckin
		case "简单", "easy":
			t.LevelCode = LevelEasy
		case "中等", "中级", "medium":
			t.LevelCode = LevelMedium
		case "困难", "高级", "hard":
			t.LevelCode = LevelHard
		default:
			t.LevelCode = 0
		}
		t.Level = ""
	}
	return t
}
