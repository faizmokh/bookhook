package hooks

import "strings"

// FindHashtag takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindHashtag(slice []Hashtag, val string) (int, bool) {
	for i, item := range slice {
		if strings.ToLower(item.Text) == strings.ToLower(val) {
			return i, true
		}
	}
	return -1, false
}

// FindUrl takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindUrl(slice []URL) (int, bool) {
	for i, item := range slice {
		if len(item.ExpandedURL) > 0 {
			return i, true
		}
	}
	return -1, false
}
