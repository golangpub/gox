package gox

type Event struct {
	Type  int    `json:"type,omitempty"`
	Major int64  `json:"major,omitempty"`
	Minor int64  `json:"minor,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
