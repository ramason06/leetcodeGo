package hard

import "strconv"

func findExpressions(num string, target int) []string {
	var result []string
	dfs(num, target, 0, "", 0, 0, &result)
	return result
}

func dfs(num string, target, pos int, path string, eval, multed int, result *[]string) {
	if pos == len(num) {
		if eval == target {
			*result = append(*result, path)
		}
		return
	}

	for i := pos; i < len(num); i++ {
		if i != pos && num[pos] == '0' {
			break // 숫자는 0으로 시작할 수 없습니다.
		}
		cur, _ := strconv.Atoi(num[pos : i+1])
		if pos == 0 {
			// 초기 호출
			dfs(num, target, i+1, path+strconv.Itoa(cur), cur, cur, result)
		} else {
			dfs(num, target, i+1, path+"+"+strconv.Itoa(cur), eval+cur, cur, result)
			dfs(num, target, i+1, path+"-"+strconv.Itoa(cur), eval-cur, -cur, result)
			dfs(num, target, i+1, path+"*"+strconv.Itoa(cur), eval-multed+multed*cur, multed*cur, result)
		}
	}
}
