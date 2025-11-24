package main

import "sort"

// GenerateAllCandidates は全ての候補セット（ラン・グループ）を生成する
func GenerateAllCandidates(tiles []Tile) [][]Tile {
	var candidates [][]Tile

	// ジョーカーを分離
	var normalTiles []Tile
	var jokers []Tile
	for _, tile := range tiles {
		if tile.IsJoker {
			jokers = append(jokers, tile)
		} else {
			normalTiles = append(normalTiles, tile)
		}
	}

	// ランの候補を生成（ジョーカー含む）
	runs := generateRunsWithJokers(normalTiles, jokers)
	candidates = append(candidates, runs...)

	// グループの候補を生成（ジョーカー含む）
	groups := generateGroupsWithJokers(normalTiles, jokers)
	candidates = append(candidates, groups...)

	return candidates
}

// generateRuns は同じ色で連続する数字の組み合わせを生成
func generateRuns(tiles []Tile) [][]Tile {
	var runs [][]Tile

	// 色ごとにタイルを分類
	byColor := make(map[Color][]Tile)
	for _, tile := range tiles {
		if tile.IsJoker {
			continue // ジョーカーは一旦スキップ
		}
		byColor[tile.Color] = append(byColor[tile.Color], tile)
	}

	// 各色について連続する数字を探す
	for _, colorTiles := range byColor {
		// 数字でソート
		sort.Slice(colorTiles, func(i, j int) bool {
			return colorTiles[i].Number < colorTiles[j].Number
		})

		// 連続部分列からランを生成（3枚以上13枚以下）
		for start := 0; start < len(colorTiles); start++ {
			var run []Tile
			var prevNum TileNumber = -1

			for i := start; i < len(colorTiles); i++ {
				tile := colorTiles[i]

				// 連続していない場合は終了
				if prevNum != -1 && tile.Number != prevNum+1 {
					break
				}

				// 重複をスキップ
				if prevNum == tile.Number {
					continue
				}

				run = append(run, tile)
				prevNum = tile.Number

				// 3枚以上ならランとして追加
				if len(run) >= 3 {
					candidate := make([]Tile, len(run))
					copy(candidate, run)
					runs = append(runs, candidate)
				}
			}
		}
	}

	return runs
}

// generateRunsWithJokers はジョーカーを含むランの候補を生成
func generateRunsWithJokers(tiles []Tile, jokers []Tile) [][]Tile {
	var runs [][]Tile

	// まずジョーカーなしのランを生成
	runs = append(runs, generateRuns(tiles)...)

	if len(jokers) == 0 {
		return runs
	}

	// 色ごとにタイルを分類
	byColor := make(map[Color][]Tile)
	for _, tile := range tiles {
		byColor[tile.Color] = append(byColor[tile.Color], tile)
	}

	// 各色について、ジョーカーを使ったランを生成
	colors := []Color{Red, Blue, Yellow, Black}
	for _, color := range colors {
		colorTiles := byColor[color]

		// 数字でソート
		sort.Slice(colorTiles, func(i, j int) bool {
			return colorTiles[i].Number < colorTiles[j].Number
		})

		// 存在する数字のセット
		hasNumber := make(map[TileNumber]bool)
		for _, t := range colorTiles {
			hasNumber[t.Number] = true
		}

		// 1-13の各開始位置から、ジョーカーを使ったランを探す
		for startNum := TileNumber(1); startNum <= 11; startNum++ {
			for length := TileNumber(3); length <= 13-startNum+1; length++ {
				var run []Tile
				jokersNeeded := 0
				valid := true

				for num := startNum; num < startNum+length; num++ {
					if hasNumber[num] {
						// 対応するタイルを見つける
						for _, t := range colorTiles {
							if t.Number == num {
								run = append(run, t)
								break
							}
						}
					} else {
						jokersNeeded++
						if jokersNeeded > len(jokers) {
							valid = false
							break
						}
					}
				}

				if valid && jokersNeeded > 0 && jokersNeeded <= len(jokers) && len(run)+jokersNeeded >= 3 {
					// ジョーカーを追加
					for i := 0; i < jokersNeeded; i++ {
						run = append(run, jokers[i])
					}
					runs = append(runs, run)
				}
			}
		}
	}

	return runs
}

// generateGroups は同じ数字で異なる色の組み合わせを生成
func generateGroups(tiles []Tile) [][]Tile {
	var groups [][]Tile

	// 数字ごとにタイルを分類
	byNumber := make(map[TileNumber][]Tile)
	for _, tile := range tiles {
		if tile.IsJoker {
			continue
		}
		byNumber[tile.Number] = append(byNumber[tile.Number], tile)
	}

	// 各数字について色の組み合わせを生成
	for _, numTiles := range byNumber {
		// 色の重複を除去
		colorSet := make(map[Color]Tile)
		for _, tile := range numTiles {
			colorSet[tile.Color] = tile
		}

		var uniqueTiles []Tile
		for _, tile := range colorSet {
			uniqueTiles = append(uniqueTiles, tile)
		}

		// 3枚以上の組み合わせを生成
		if len(uniqueTiles) >= 3 {
			// 3枚の組み合わせ
			combinations := getCombinations(uniqueTiles, 3)
			groups = append(groups, combinations...)

			// 4枚の組み合わせ
			if len(uniqueTiles) >= 4 {
				combinations = getCombinations(uniqueTiles, 4)
				groups = append(groups, combinations...)
			}
		}
	}

	return groups
}

// generateGroupsWithJokers はジョーカーを含むグループの候補を生成
func generateGroupsWithJokers(tiles []Tile, jokers []Tile) [][]Tile {
	var groups [][]Tile

	// まずジョーカーなしのグループを生成
	groups = append(groups, generateGroups(tiles)...)

	if len(jokers) == 0 {
		return groups
	}

	// 数字ごとにタイルを分類
	byNumber := make(map[TileNumber][]Tile)
	for _, tile := range tiles {
		byNumber[tile.Number] = append(byNumber[tile.Number], tile)
	}

	// 各数字について、ジョーカーを使ったグループを生成
	for num := TileNumber(1); num <= 13; num++ {
		numTiles := byNumber[num]

		// 色の重複を除去
		colorSet := make(map[Color]Tile)
		for _, tile := range numTiles {
			colorSet[tile.Color] = tile
		}

		var uniqueTiles []Tile
		for _, tile := range colorSet {
			uniqueTiles = append(uniqueTiles, tile)
		}

		existingColors := len(uniqueTiles)

		// 2枚 + 1ジョーカー = 3枚のグループ
		if existingColors == 2 && len(jokers) >= 1 {
			group := make([]Tile, len(uniqueTiles))
			copy(group, uniqueTiles)
			group = append(group, jokers[0])
			groups = append(groups, group)
		}

		// 2枚 + 2ジョーカー = 4枚のグループ
		if existingColors == 2 && len(jokers) >= 2 {
			group := make([]Tile, len(uniqueTiles))
			copy(group, uniqueTiles)
			group = append(group, jokers[0], jokers[1])
			groups = append(groups, group)
		}

		// 3枚 + 1ジョーカー = 4枚のグループ
		if existingColors == 3 && len(jokers) >= 1 {
			group := make([]Tile, len(uniqueTiles))
			copy(group, uniqueTiles)
			group = append(group, jokers[0])
			groups = append(groups, group)
		}

		// 1枚 + 2ジョーカー = 3枚のグループ
		if existingColors == 1 && len(jokers) >= 2 {
			group := make([]Tile, len(uniqueTiles))
			copy(group, uniqueTiles)
			group = append(group, jokers[0], jokers[1])
			groups = append(groups, group)
		}
	}

	return groups
}

// getCombinations はn個からr個を選ぶ組み合わせを生成
func getCombinations(tiles []Tile, r int) [][]Tile {
	var result [][]Tile
	n := len(tiles)

	var backtrack func(start int, current []Tile)
	backtrack = func(start int, current []Tile) {
		if len(current) == r {
			combo := make([]Tile, r)
			copy(combo, current)
			result = append(result, combo)
			return
		}
		for i := start; i < n; i++ {
			backtrack(i+1, append(current, tiles[i]))
		}
	}

	backtrack(0, []Tile{})
	return result
}

// TileKey はタイルを一意に識別するためのキー
type TileKey struct {
	Number  int
	Color   Color
	IsJoker bool
	Index   int // 同じタイルが複数ある場合の区別用
}

// CandidateInfo は候補セットの情報
type CandidateInfo struct {
	tiles   []Tile
	indices []int
}

// SolveCheckmate は詰み判定を行い、解があれば解を返す
func SolveCheckmate(board Board, hand Hand) (bool, []Meld) {
	// 全タイルを収集し、IDを付与
	var allTiles []Tile
	var id TileID = 0
	for _, meld := range board.Melds {
		for _, tile := range meld {
			tile.ID = id
			allTiles = append(allTiles, tile)
			id++
		}
	}
	for _, tile := range hand.Tiles {
		tile.ID = id
		allTiles = append(allTiles, tile)
		id++
	}

	// タイルがない場合は詰み（出し切っている）
	if len(allTiles) == 0 {
		return true, nil
	}

	// 全候補セットを生成
	candidates := GenerateAllCandidates(allTiles)

	// Exact Coverで解を探索
	solution := exactCover(allTiles, candidates)
	if solution == nil {
		return false, nil
	}

	// 解をMeldに変換
	var melds []Meld
	for _, candidate := range solution {
		melds = append(melds, Meld(candidate))
	}
	return true, melds
}

// exactCover はバックトラッキングでExact Cover問題を解く
func exactCover(tiles []Tile, candidates [][]Tile) [][]Tile {
	// IDからインデックスへのマップを作成
	idToIndex := make(map[TileID]int)
	for i, tile := range tiles {
		idToIndex[tile.ID] = i
	}

	// 各候補が使うタイルのインデックスを計算
	var candidateInfos []CandidateInfo

	for _, candidate := range candidates {
		var indices []int
		valid := true

		for _, tile := range candidate {
			if idx, ok := idToIndex[tile.ID]; ok {
				indices = append(indices, idx)
			} else {
				valid = false
				break
			}
		}

		if valid && len(indices) == len(candidate) {
			candidateInfos = append(candidateInfos, CandidateInfo{
				tiles:   candidate,
				indices: indices,
			})
		}
	}

	// バックトラッキング
	used := make([]bool, len(tiles))
	var solution [][]Tile
	if backtrack(candidateInfos, used, 0, &solution) {
		return solution
	}
	return nil
}

// backtrack はExact Coverのバックトラッキング探索
func backtrack(candidates []CandidateInfo, used []bool, covered int, solution *[][]Tile) bool {
	// 全てカバーできたら成功
	if covered == len(used) {
		return true
	}

	// 最初の未使用タイルを見つける
	firstUncovered := -1
	for i, u := range used {
		if !u {
			firstUncovered = i
			break
		}
	}

	if firstUncovered == -1 {
		return true
	}

	// このタイルを含む候補を試す
	for _, candidate := range candidates {
		containsFirst := false
		for _, idx := range candidate.indices {
			if idx == firstUncovered {
				containsFirst = true
				break
			}
		}
		if !containsFirst {
			continue
		}

		// この候補が使えるかチェック
		canUse := true
		for _, idx := range candidate.indices {
			if used[idx] {
				canUse = false
				break
			}
		}

		if canUse {
			// 使用済みにマーク
			for _, idx := range candidate.indices {
				used[idx] = true
			}

			// 再帰
			if backtrack(candidates, used, covered+len(candidate.indices), solution) {
				*solution = append(*solution, candidate.tiles)
				return true
			}

			// 元に戻す
			for _, idx := range candidate.indices {
				used[idx] = false
			}
		}
	}

	return false
}
