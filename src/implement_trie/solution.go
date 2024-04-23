package implement_trie

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a
dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:
Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix,
and false otherwise.


Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True


Constraints:

1 <= word.length, prefix.length <= 2000
word and prefix consist only of lowercase English letters.
At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.
*/

type Trie struct {
	childNodes [26]*Trie
	wordEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	currentNode := t
	for _, k := range word {
		idx := k - 'a'
		if currentNode.childNodes[idx] == nil {
			currentNode.childNodes[idx] = &Trie{}
		}
		currentNode = currentNode.childNodes[idx]
	}
	currentNode.wordEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t
	for _, k := range word {
		idx := k - 'a'
		if currentNode.childNodes[idx] == nil {
			return false
		}
		currentNode = currentNode.childNodes[idx]
	}
	return currentNode.wordEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t
	for _, k := range prefix {
		idx := k - 'a'
		if currentNode.childNodes[idx] == nil {
			return false
		}
		currentNode = currentNode.childNodes[idx]
	}
	return true
}
