package recursion

import (
	"fmt"
)

func FindWay() {
	//获取地图
	var mymap [8][8]int
	//生成围墙
	for i := 0; i < len(mymap); i++ {
		mymap[i][0] = 1;
		mymap[0][i] = 1;
		mymap[7][i] = 1;
		mymap[i][7] = 1;
	}

	mymap[3][2] = 1;
	mymap[3][1] = 1;
	mymap[3][3] = 1;
	//打印地图
	printMap(mymap)

	settWay(&mymap,1,1)
	fmt.Println("====================================")

	printMap(mymap)
}

//定义一些规则0没走过1墙2位通路3表示该点走过不通
//探路的规则下->ringt->top->left
func settWay(mp *[8][8]int, i int, j int) bool {
	if mp[6][6] == 2 {
		return true
	} else {
		//bottom->right->top->left
		if mp[i][j] == 0 { //这个点没走过
			mp[i][j] = 2 //该点可以走
			if settWay(mp, i+1, j) {
				return true
			} else if settWay(mp, i, j+1) {
				return true
			} else if settWay(mp, i-1, j) {
				return true
			} else if settWay(mp, i, j-1) {
				return true
			} else {
				mp[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func printMap(mp [8][8]int) {
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			fmt.Print(mp[i][j], " ")
		}
		fmt.Println("")
	}
}
