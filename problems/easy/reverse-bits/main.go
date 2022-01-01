package easy_reverse_bits

/*
https://leetcode.com/problems/reverse-bits/
*/
func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}

	rBits := uint32(0)
	shift := 31
	for {
		if num&1 == 1 {
			rBits ^= 1
		}
		if num == 0 || num == 1 {
			break
		}
		shift--
		num >>= 1
		rBits <<= 1
	}
	rBits <<= shift

	return rBits
}
