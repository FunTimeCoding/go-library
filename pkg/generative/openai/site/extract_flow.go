package site

import "fmt"

func (s *Site) ExtractFlow(verbose bool) string {
	if false {
		// Unstable selector
		s.printNewButton()
	}

	s.clickNew()

	if false {
		s.printProfileButton()
	}

	s.clickProfile()

	s.clickSettings()
	s.clickPersonalize()
	s.clickMemories()
	result := s.readMemories()

	if verbose {
		fmt.Printf("Memories: %d\n", len(result))
	}

	if false {
		// Unstable selector
		s.printCloseMemories()
	}

	s.clickCloseMemories()

	if false {
		s.printCloseSettings()
	}

	s.clickCloseSettings()

	return result
}
