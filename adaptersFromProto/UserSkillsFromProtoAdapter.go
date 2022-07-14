package adaptersfromproto

type UserSkillsFromProtoAdapter struct {
	Data []Users `json:"data"`
}

type Users struct {
	Email      string       `json:"id"`
	English    uint         `json:"english"`
	UserSkills []UserSkills `json:"skills"`
}

type UserSkills struct {
	UserID    string `json:"id"`
	Details   string `json:"details"`
	SkillName string `json:"name"`
	Level     uint   `json:"level"`
}
