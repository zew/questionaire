package qst

func init() {

	fourPermutationsOf6x3x3[3] = []preferences3x3T{
		{
			ID: 1, // Frage 1
			Ppls: [][]int{
				{0, 2, 3},
				{3, 0, 2},
				{2, 3, 0},
			},
		},
		{
			ID: 2, // Frage 2
			Ppls: [][]int{
				{0, 3, 2},
				{3, 0, 2},
				{2, 2, 1},
			},
		},
		{
			ID: 3, // Frage 3
			Ppls: [][]int{
				{3, 0, 2},
				{2, 0, 3},
				{0, 5, 0},
			},
		},
		{
			ID: 4, // Frage 4
			Ppls: [][]int{
				{1, 0, 4},
				{3, 1, 1},
				{1, 4, 0},
			},
		},
		{
			ID: 5, // Frage 5
			Ppls: [][]int{
				{2, 1, 2},
				{3, 0, 2},
				{0, 4, 1},
			},
		},
		{
			ID: 6, // Frage 6
			Ppls: [][]int{
				{1, 0, 4},
				{4, 0, 1},
				{0, 5, 0},
			},
		},
	}

}
