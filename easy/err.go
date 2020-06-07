package easy

type _err int

func (e _err) Error() string {
	return ""
}

const (
	Exit _err = _err(iota)
)
