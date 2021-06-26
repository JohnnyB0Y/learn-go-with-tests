package main

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) string {
	return d[word]
}

type IntDictionary map[string]int

func (d IntDictionary) ValueForKey(key string) int {
	return d[key]
}
