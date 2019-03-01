package main

import Trie from "./trie"

func main() {
	array := []string{"dog", "deer", "deal"}
	dictTrie := buildDictionary(array)
}
func buildDictionary(array []string) *Trie {
	return trie.New()
}
