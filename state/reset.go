package state

func (s *State) Reset() {
	keys := s.Keys
	*s = *BlankState
	s.Keys = keys
}
