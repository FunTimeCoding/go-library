package missing

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/text/dictionary"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/constant"
	"io/fs"
	"path/filepath"
	"strings"
)

func Run() {
	f := constant.File
	wordUsage := make(map[string]*dictionary.WordUsage)
	totalWords := 0

	for _, category := range dictionary.Read(f) {
		for _, word := range category.Words {
			wordKey := strings.ToLower(word)
			wordUsage[wordKey] = &dictionary.WordUsage{
				Word:     word,
				Category: category.Name,
				Used:     false,
			}
			totalWords++
		}
	}

	fmt.Printf("Check %d words\n", totalWords)
	filesScanned := 0
	errors.PanicOnError(
		filepath.WalkDir(
			".",
			func(
				path string,
				d fs.DirEntry,
				err error,
			) error {
				if err != nil {
					return err
				}

				if d.IsDir() {
					if constant.Skip[d.Name()] {
						return filepath.SkipDir
					}

					return nil
				}

				if !dictionary.IncludeFile(d.Name()) {
					return nil
				}

				dictionary.ScanFile(path, wordUsage)
				filesScanned++

				return nil
			},
		),
	)

	fmt.Printf("Scanned %d files\n", filesScanned)
	unused := make(map[string][]string)
	usedCount := 0

	for _, usage := range wordUsage {
		if usage.Used {
			usedCount++
		} else {
			if unused[usage.Category] == nil {
				unused[usage.Category] = make([]string, 0)
			}
			unused[usage.Category] = append(
				unused[usage.Category],
				usage.Word,
			)
		}
	}

	fmt.Printf("Results:\n")
	fmt.Printf(
		"Used: %d/%d words (%.1f%%)\n",
		usedCount,
		totalWords,
		float64(usedCount)/float64(totalWords)*100,
	)
	fmt.Printf("Unused: %d words\n", totalWords-usedCount)

	if len(unused) == 0 {
		fmt.Println("No unused dictionary words")

		return
	}

	fmt.Println("Unused words by category:")

	for category, words := range unused {
		fmt.Printf("\n# %s (%d unused):\n", category, len(words))

		for _, word := range words {
			fmt.Printf("  %s\n", word)
		}
	}
}
