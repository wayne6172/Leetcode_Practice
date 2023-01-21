package Q212

func findWords(board [][]byte, words []string) []string {
	// 建立字典樹
	root := &TrieNode{}
	for _, word := range words {
		root.Insert(word)
	}

	// 建立結果切片
	result := []string{}

	// 遍歷每一個格子
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, i, j, root, &result)
		}
	}

	return result
}

func dfs(board [][]byte, i int, j int, p *TrieNode, result *[]string) {
	c := board[i][j]
	if c == '#' || p.Next[c-'a'] == nil {
		return
	}
	p = p.Next[c-'a']
	if p.End {
		*result = append(*result, p.Word)
		p.End = false
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, p, result)
	}
	if j > 0 {
		dfs(board, i, j-1, p, result)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, p, result)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, p, result)
	}
	board[i][j] = c
}

type TrieNode struct {
	Next [26]*TrieNode
	End  bool
	Word string
}

func (t *TrieNode) Insert(s string) {
	p := t
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p.Next[c-'a'] == nil {
			p.Next[c-'a'] = &TrieNode{}
		}
		p = p.Next[c-'a']
	}
	p.End = true
	p.Word = s
}
