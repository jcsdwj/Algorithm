/**
  @author: 伍萬
  @since: 2024/7/15
  @desc: //TODO
**/

package leetCode

import "fmt"

//721. 账户合并

func TestLeetCode721() {
	accounts1 := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	accounts2 := [][]string{
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
		{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
		{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
		{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
		{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
	}
	fmt.Println(accountsMerge(accounts1))
	fmt.Println(accountsMerge(accounts2))
}

func accountsMerge(accounts [][]string) [][]string {
	allUsers := []string{""}
	for i := 0; i < len(accounts); i++ {
		allUsers = append(allUsers, accounts[i][0])
	}

	// 每个邮件映射一个用户
	dictRow := make(map[int]int) // 每一行映射的用户行
	dictRow[0] = 1               // 表示第一行映射第一行
	dictMail := make(map[string]int)

	for i := 0; i < len(accounts[0]); i++ {
		dictMail[accounts[0][i]] = 1 // 表示第一行的邮件都映射第一个人
	}

	for i := 1; i < len(accounts); i++ {
		curRow := i + 1
		dictRow[i] = curRow // 默认映射当前行
		// 如果出现重复的
		for j := 1; j < len(accounts[0]); j++ {
			if dictMail[accounts[i][j]] != 0 {
				// 出现重复的了
				curRow = dictMail[accounts[i][j]] // 获取重复的行号
				dictRow[i] = curRow               // 当前行的user改为之前的
				j = len(accounts[0])
			}
		}

		// 再次更新
		for j := 1; j < len(accounts); j++ {
			dictMail[accounts[i][j]] = curRow // 更新邮件
		}
	}

	// 把相同cur的放一行去重然后排序
	res := [][]string{}
	//for k, v := range dictRow {
	//	// v相同的放一行
	//}
	return res
}
