package fsutils

type firstErrorHolder struct {
	innerErr error
}

func (h firstErrorHolder) firstErr() error {
	return h.innerErr
}

func (h *firstErrorHolder) updateErr(anotherErr error) {
	if h.innerErr == nil {
		h.innerErr = anotherErr
	}
}
