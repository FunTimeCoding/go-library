package site

import (
	"fmt"
	"time"
)

func (s *Site) ExtractFlow(verbose bool) string {
	s.NewChat()

	if false {
		s.printProfile()
	}

	s.clickProfile()

	s.clickSettings()
	s.clickPersonalize()

	if false {
		s.printMemories()
	}

	s.clickMemories()

	time.Sleep(2 * time.Second)

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
