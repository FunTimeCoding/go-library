package mock_entry

import "io/fs"

func (f *Entry) Mode() fs.FileMode { return 0644 }
