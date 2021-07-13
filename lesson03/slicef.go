package main

func max1(s []string) string {
	ans := s[0]
	mx := len(ans)
	for _, k := range s {
		if len(k) > mx {
			ans = k
			mx = len(k)
		}
	}
	return ans
}
