package state

func (s *State) Exists() bool {
	if s.Name == "" {
		return false
	}

	if s.Scene == "" {
		return false
	}

	return true
}
