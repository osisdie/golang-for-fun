package codechallenge

import (
	"reflect"
	"testing"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = NewTrieNode()
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return true
}

func (t *Trie) InsertIfNoPrefix(word string) bool {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = NewTrieNode()
		}
		current = current.children[ch]

		// If we find a word in the Trie that is a prefix of the current word
		if current.isEnd {
			return false
		}
	}

	// If the current word is a prefix of an already inserted word
	if len(current.children) > 0 {
		return false
	}

	current.isEnd = true
	return true
}

// Trie
func TrieNoPrefix(words []string) []string {
	trie := NewTrie()

	for _, word := range words {
		if !trie.InsertIfNoPrefix(word) {
			// fmt.Println("BAD SET")
			// fmt.Println(word)
			return []string{"BAD SET", word}
		}
	}

	// fmt.Println("GOOD SET")
	return []string{"BAD SET"}
}

func TestTrieNoPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "Test case 1",
			input: []string{"hgiiccfchbeadgebc",
				"biiga",
				"edchgb",
				"ccfdbeajaeid",
				"ijgbeecjbj",
				"bcfbbacfbfcfbhcbfjafibfhffac",
				"ebechbfhfcijcjbcehbgbdgbh",
				"ijbfifdbfifaidje",
				"acgffegiihcddcdfjhhgadfjb",
				"ggbdfdhaffhghbdh",
				"dcjaichjejgheiaie",
				"d",
				"jeedfch",
				"ahabicdffbedcbdeceed",
				"fehgdfhdiffhegafaaaiijceijdgbb",
				"beieebbjdgdhfjhj",
				"ehg",
				"fdhiibhcbecddgijdb",
				"jgcgafgjjbg",
				"c",
				"fiedahb",
				"bhfhjgcdbjdcjjhaebejaecdheh",
				"gbfbbhdaecdjaebadcggbhbchfjg",
				"jdjejjg",
				"dbeedfdjaghbhgdhcedcj",
				"decjacchhaciafafdgha",
				"a",
				"hcfibighgfggefghjh",
				"ccgcgjgaghj",
				"bfhjgehecgjchcgj",
				"bjbhcjcbbhf",
				"daheaggjgfdcjehidfaedjfccdafg",
				"efejicdecgfieef",
				"ciidfbibegfca",
				"jfhcdhbagchjdadcfahdii",
				"i",
				"abjfjgaghbc",
				"bddeejeeedjdcfcjcieceieaei",
				"cijdgbddbceheaeececeeiebafgi",
				"haejgebjfcfgjfifhihdbddbacefd",
				"bhhjbhchdiffb",
				"jbbdhcbgdefifhafhf",
				"ajhdeahcjjfie",
				"idjajdjaebfhhaacecb",
				"bhiehhcggjai",
				"bjjfjhiice",
				"aif",
				"gbbfjedbhhhjfegeeieig",
				"fefdhdaiadefifjhedaieaefc",
				"hgaejbhdebaacbgbgfbbcad",
				"heghcb",
				"eggadagajjgjgaihjdigihfhfbijbh",
				"jadeehcciedcbjhdeca",
				"ghgbhhjjgghgie",
				"ibhihfaeeihdffjgddcj",
				"hiedaegjcdai",
				"bjcdcafgfjdejgiafdhfed",
				"fgdgjaihdjaeefejbbijdbfabeie",
				"aeefgiehgjbfgidcedjhfdaaeigj",
				"bhbiaeihhdafgaciecb",
				"igicjdajjdegbceibgebedghihihh",
				"baeeeehcbffd",
				"ajfbfhhecgaghgfdadbfbb",
				"ahgaccehbgajcdfjihicihhc",
				"bbjhih",
				"a",
				"cdfcdejacaicgibghgddd",
				"afeffehfcfiefhcagg",
				"ajhebffeh",
				"e",
				"hhahehjfgcj",
				"ageaccdcbbcfidjfc",
				"gfcjahbbhcbggadcaebae",
				"gi",
				"edheggceegiedghhdfgabgcd",
				"hejdjjbfacggdccgahiai",
				"ffgeiadgjfgecdbaebagij",
				"dgaiahge",
				"hdbaifh",
				"gbhccajcdebcig",
				"ejdcbbeiiebjcagfhjfdahbif",
				"g",
				"ededbjaaigdhb",
				"ahhhcibdjhidbgefggdjebfcf",
				"bdigjaehfchebiedajcjdh",
				"fjehjgbdbaiifi",
				"fbgigbdcbcgffdicfcidfdafghajc",
				"ccajeeijhhb",
				"gaaagfacgiddfahejhbgdfcfbfeedh",
				"gdajaigfbjcdegeidgaccjfi",
				"fghechfchjbaebcghfcfbdicgaic",
				"cfhigaciaehacdjhfcgajgbhhgj",
				"edhjdbdjccbfihiaddij",
				"cbbhagjbcadegicgifgghai",
				"hgdcdhieji",
				"fbifgbhdhagch",
				"cbgcdjea",
				"dggjafcajhbbbaja",
				"bejihed",
				"eeahhcggaaidifdigcfjbficcfhjj",
			},
			expected: []string{"BAD SET", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrieNoPrefix(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
