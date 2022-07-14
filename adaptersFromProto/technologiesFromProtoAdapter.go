package adaptersfromproto

type TechnologiesFromProtoAdapter struct {
	Technologies []Technology `json:"technologies"`
}

type Technology struct {
	Label    string `json:"label"`
	Category string `json:"category"`
}
