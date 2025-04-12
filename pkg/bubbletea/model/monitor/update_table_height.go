package monitor

func (m *Model) updateTableHeight(
	addOne bool,
	removeOne bool,
) {
	// addOne: add 1 to height on resize OR when toasts go back to 0
	// why: first toast is hidden, or a blank line adds at bottom
	// does not happen on startup

	// removeOne: remove 1 from height on resize when toasts are shown
	// why: resize when a toast is shown hides first toast

	// top bar 1, bottom bar 1, table 3 for header
	// table bottom border does not count
	static := 5
	toasts := len(m.toast)

	// first toast already makes table smaller, only increase static for each more than 1
	if toasts > 1 {
		static += toasts - 1
	}

	height := m.height - static

	if addOne {
		height += 1
	}

	if removeOne {
		height -= 1
	}

	m.table.SetHeight(height)
}
