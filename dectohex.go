func dechex (number int64, cutTrailingZeros bool) (string, string) {

	num 	:= number
	length	:= 16
	data 	:= make([]int64, length)

	neg		:= number < 0

	if neg {
		num = -1 * number
	}

	i := 0

	for num > 15 {
		tmp := num % 16
		data[i] = tmp
		num = (int64)(num / 16)
		i++
	}

	data[i] = num

	j 		:= length-1
	strPos 	:= ""
	strNeg	:= ""

	for j >= 0 {
		val := data[j];

		if val <= 9 {
			strPos += string(('0'+val))
		} else {
			strPos += string('A'+(val-10))
		}

		if val <= 5 {
			strNeg += string(('F'-val))
		} else {
			strNeg += string(('0'+(15-val)))
		}
		fmt.Println(val)

		j--
	}

	zeros := 0
	j=0

	for j < length && cutTrailingZeros {
		if strPos[j] == '0' {
			zeros++
		} else {
			break
		}
		j++
	}
	return strPos[zeros:], strNeg

}

