package types

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g Gender) IsValid() bool {
	switch g {
	case Male:
		return true
	case Female:
		return true
	default:
		return false
	}
}
