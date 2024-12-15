package main

func rangeSum(low, high, nums int64) int64 {
	if nums%2 == 0 {
		return (low + high) * (nums / 2)
	}
	return nums * (low + high) / 2
}

func numFromByte(b byte) int64 {
	return int64(b - '0')
}

func numToByte(num int64) byte {
	return byte(num + '0')
}

func solve9a() {
	var s string
	scan(&s)

	b := []byte(s)

	maxIdx := int64(len(b)/2) + int64(len(b)%2) - 1

	var res int64
	var i, j int64 = 0, int64(len(b)) - 1
	var numI, numJ int64 = 0, maxIdx
	var posIdx int64 = 0
	for i < j {
		if i%2 == 0 {
			fileLen := numFromByte(b[i])
			res += numI * rangeSum(posIdx, posIdx+fileLen-1, fileLen)
			numI++
			i++
			posIdx += fileLen
			continue
		}

		if j%2 == 1 {
			j--
			continue
		}

		emptyI := numFromByte(b[i])
		fileJ := numFromByte(b[j])

		canAlloc := min(emptyI, fileJ)
		b[i] = numToByte(emptyI - canAlloc)
		b[j] = numToByte(fileJ - canAlloc)

		res += numJ * rangeSum(posIdx, posIdx+canAlloc-1, canAlloc)
		posIdx += canAlloc

		if b[i] == '0' {
			i++
		}
		if b[j] == '0' {
			j--
			numJ--
		}
	}

	if j%2 == 0 {
		fileLen := numFromByte(b[j])
		res += numJ * rangeSum(posIdx, posIdx+fileLen-1, fileLen)
	}

	printToOut(res)
}

type freeSpaceInfo struct {
	free   int64
	posIdx int64
}

func getFreeSpaceForFileLen(freeSpaces [][]*freeSpaceInfo, curPos int64, fileLen int64) *freeSpaceInfo {
	for i := fileLen; i <= 9; i++ {
		for len(freeSpaces[i]) > 0 {
			if freeSpaces[i][0].free >= fileLen && freeSpaces[i][0].posIdx < curPos {
				return freeSpaces[i][0]
			}

			freeSpaces[i] = freeSpaces[i][1:]
		}
	}

	return nil
}

func solve9b() {
	var s string
	scan(&s)

	b := []byte(s)

	maxIdx := int64(len(b)/2) + int64(len(b)%2) - 1

	freeSpaces := make([][]*freeSpaceInfo, 10)
	var posIdx int64 = 0
	for i := 0; i < len(b); i++ {
		n := numFromByte(b[i])
		if i%2 == 1 {
			freeSpace := &freeSpaceInfo{
				free:   n,
				posIdx: posIdx,
			}
			for j := n; j > 0; j-- {
				freeSpaces[j] = append(freeSpaces[j], freeSpace)
			}
		}
		posIdx += n
	}

	var res int64
	var numJ = maxIdx

	for j := int64(len(b)) - 1; j > 0; j-- {
		fileLen := numFromByte(b[j])
		posIdx -= fileLen

		if j%2 == 1 {
			continue
		}

		freeSpace := getFreeSpaceForFileLen(freeSpaces, posIdx, fileLen)

		curPosIdx := posIdx
		if freeSpace != nil {
			curPosIdx = freeSpace.posIdx
			freeSpace.free -= fileLen
			freeSpace.posIdx += fileLen
		}

		res += numJ * rangeSum(curPosIdx, curPosIdx+fileLen-1, fileLen)
		numJ--
	}

	printToOut(res)
}
