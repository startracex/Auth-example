package utils

func SafeInfo(mapping *map[string]any, attach ...string) {
	// TODO delete sensitive information
	delete(*mapping, "password")
	if len(attach) > 0 {
		for _, v := range attach {
			delete(*mapping, v)
		}
	}
}
