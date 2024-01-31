package meta

type Skill struct {
	ID   string `yaml:"id"`
	Pid  string `yaml:"pid"`
	Tid  string `yaml:"tid"`
	Node int    `yaml:"node"`
}

func (m *Meta) NewSkill(id, pid, tid string, node int, image, name string, level TaskLevel) *Meta {
	n := m.R()
	n.Skill.ID = id
	n.Skill.Pid = pid
	n.Skill.Tid = tid
	n.Skill.Node = node
	n.Challenge.Name = name
	n.Task.Name = image
	n.Task.LevelCode = level
	return n
}

func NewSkill(id, pid, tid string, node int, image, name string, level TaskLevel) *Meta {
	return New("陌竹", "mozhu233@outlook.com").NewSkill(id, pid, tid, node, image, name, level)
}
