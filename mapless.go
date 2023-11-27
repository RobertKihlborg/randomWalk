package main

import "math/rand"

func CollideOrAddSort(a, b int, value vec2, list *[]vec2) bool {
	//fmt.Printf("Trying to add %v to %v, (a,b) = (%v,%v)\n", value, (*list)[a:b], a, b)
	if value == 0 {
		return true
	}
	if a == b {
		(*list)[a] = value
		return false
	}
	l, r := a, b-1

	// Binary search to find value in part of list
	for l+1 < r {
		mid := (l + r) / 2
		midVal := (*list)[mid]
		if value == midVal {
			return true
		}

		if value > midVal {
			l = mid
		} else {
			r = mid
		}
	}

	// Check if binary search resulted in finding value or finding where to insert value
	if value == (*list)[l] || value == (*list)[r] {
		return true
	}
	//fmt.Printf("%v, %v\n", int32(value), int32((*list)[l]))
	insertPos := 0
	if value > (*list)[r] {
		insertPos = r + 1
	} else if value < (*list)[l] {
		insertPos = l
	} else {
		insertPos = r
	}
	//fmt.Printf("l, r = %v, %v, Inserting at %v\n", l, r, insertPos)

	// Insert into sorted list
	for i := insertPos; i <= b; i++ {
		tmp := (*list)[i]
		(*list)[i] = value
		value = tmp
	}

	//fmt.Println("Insert successful")
	return false
}

func MaplessSAOld(n, blockSize int) *[]vec2 {
	res := make([]vec2, n+1)
	sorted := make([]vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	maplessOldHelper(1, n+1, blockSize, &res, &sorted, r)
	return &res
}

func maplessOldHelper(a, b, blockSize int, pos *[]vec2, sortedPos *[]vec2, r *rand.Rand) {
	if b-a <= blockSize {
		for true {
			direction := r.Intn(4)
			fail := false
			currPos := vec2(0)
			for i := a; i < b; i++ {
				direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
				currPos += unitVectors[direction]
				(*pos)[i] = currPos // Take a step
				fail = CollideOrAddSort(a, i, currPos, sortedPos)
				if fail {
					break
				}
			}
			if !fail {
				return
			}
		}
	}

	// Case for when to do recursive call
	mid := (a + b) / 2

	for true {
		fail := false
		maplessOldHelper(a, mid, blockSize, pos, sortedPos, r)
		maplessOldHelper(mid, b, blockSize, pos, sortedPos, r)

		// Note which nodes were visited in the first part (return this directly maybe?)
		lastPosOfFirst := (*pos)[mid-1]

		// Check which nodes were visited in the second part if we put it at the end of the first, and check for collisions
		for i := mid; i < b; i++ {
			(*pos)[i] += lastPosOfFirst
			fail = CollideOrAddSort(a, i, (*sortedPos)[i]+lastPosOfFirst, sortedPos)
			if fail {
				break
			}
		}
		if !fail {
			return
		}
	}
}

func DefaultMapless(n int) *[]vec2 {
	return MaplessSA(n, 30)
}

func MaplessSA(n, blockSize int) *[]vec2 {
	res := make([]vec2, n+1)
	sorted := make([]vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	maplessHelper(r.Intn(4), 1, n+1, blockSize, &res, &sorted, r)
	return &res
}

func maplessHelper(prevDir, a, b, blockSize int, pos *[]vec2, sortedPos *[]vec2, r *rand.Rand) int {
	if b-a <= blockSize {
		for true {
			direction := prevDir
			fail := false
			currPos := vec2(0)
			for i := a; i < b; i++ {
				direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
				currPos += unitVectors[direction]
				(*pos)[i] = currPos // Take a step
				fail = CollideOrAddSort(a, i, currPos, sortedPos)
				if fail {
					break
				}
			}
			if !fail {
				return direction
			}
		}
	}

	// Case for when to do recursive call
	mid := (a + b) / 2

	for true {
		fail := false
		direction := maplessHelper(prevDir, a, mid, blockSize, pos, sortedPos, r)
		direction = maplessHelper(direction, mid, b, blockSize, pos, sortedPos, r)

		// Note which nodes were visited in the first part (return this directly maybe?)
		lastPosOfFirst := (*pos)[mid-1]

		// Check which nodes were visited in the second part if we put it at the end of the first, and check for collisions
		for i := mid; i < b; i++ {
			(*pos)[i] += lastPosOfFirst
			fail = CollideOrAddSort(a, i, (*sortedPos)[i]+lastPosOfFirst, sortedPos)
			if fail {
				break
			}
		}
		if !fail {
			return direction
		}
	}
	return 0 // Unreachable anyway
}
