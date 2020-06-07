package hide

const Width int = 320
const Height int = 240
const Scale int = 2

func (h *Hide) Layout(ow, oh int) (int, int) {
	return Width, Height
}
