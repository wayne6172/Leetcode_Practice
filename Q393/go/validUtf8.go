package validUtf8

func identityLength(mask int) int {
	i := 7
	for ; i >= 0; i-- {
		if mask&(1<<i) == 0 {
			break
		}
	}

	switch i {
	case 7:
		return 1
	case 5:
		return 2
	case 4:
		return 3
	case 3:
		return 4
	default:
		return -1
	}
}

func validUtf8(data []int) bool {
	for length := len(data); length > 0; {
		bytes := identityLength(data[0])

		if bytes == -1 || length < bytes {
			return false
		} else {
			split := data[:bytes]

			for i := 1; i < bytes; i++ {
				if (split[i]&(1<<7)) == 0 || (split[i]&(1<<6)) > 0 {
					return false
				}
			}
		}

		data = data[bytes:]
		length = len(data)
	}

	return true
}
