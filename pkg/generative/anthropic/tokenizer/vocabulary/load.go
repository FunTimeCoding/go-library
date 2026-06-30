package vocabulary

import "encoding/json"

func Load() (*Vocabulary, error) {
	content, e := readCached()

	if e != nil {
		content, e = download()

		if e != nil {
			return nil, e
		}
	}

	var v Vocabulary

	if f := json.Unmarshal(content, &v); f != nil {
		return nil, f
	}

	return &v, nil
}
