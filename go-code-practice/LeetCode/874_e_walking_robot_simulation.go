package main

/*
A robot on an infinite XY-plane starts at point (0, 0) and faces north.
The robot can receive one of three possible types of commands:

-2: turn left 90 degrees,
-1: turn right 90 degrees, or
1 <= k <= 9: move forward k units.
Some of the grid squares are obstacles. The ith obstacle is at grid point obstacles[i] = (xi, yi).

If the robot would try to move onto them, the robot stays on the previous grid square instead
(but still continues following the rest of the route.)

Return the maximum Euclidean distance that the robot will be from the origin squared (i.e. if the distance is 5, return 25).

Note:
North means +Y direction.
East means +X direction.
South means -Y direction.
West means -X direction.


Example 1:
Input: commands = [4,-1,3], obstacles = []
Output: 25
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 3 units to (3, 4).
The furthest point away from the origin is (3, 4), which is 3^2 + 4^2 = 25 units away.

Example 2:
Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
Output: 65
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 1 unit and get blocked by the obstacle at (2, 4), robot is at (1, 4).
4. Turn left.
5. Move north 4 units to (1, 8).
The furthest point away from the origin is (1, 8), which is 1^2 + 8^2 = 65 units away.


Constraints:
1 <= commands.length <= 104
commands[i] is one of the values in the list [-2,-1,1,2,3,4,5,6,7,8,9].
0 <= obstacles.length <= 104
-3 * 104 <= xi, yi <= 3 * 104
The answer is guaranteed to be less than 2^31
*/

// 題目意思，機器人會依據指令遊走在一個平面上。平面上會有可能會有數個障礙物。求最後機器人相對於原點(0,0)所展開的最大行徑面積
func robotSim(commands []int, obstacles [][]int) int {
	// 宣告一個存放障礙物的字典陣列
	obs := make(map[[2]int]struct{})
	for _, v := range obstacles {
		// 將障礙物的位置賦予字典陣列一個固定值
		obs[[2]int{v[0], v[1]}] = struct{}{}
	}

	// 定義平面座標x,y
	x, y := 0, 0
	// 定義方向向量 上 右 下 左
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// 當前座標位置的方向 0,1,2,3 分別表示 上 右 下 左
	cur := 0
	// 最大行徑面積
	res := 0
	// 將指令依次輸入，控制`cur`位置
	for _, v := range commands {
		// 當指令為`-2`或`-1`時，做`左`或`右`轉
		if v == -2 {
			cur = (cur + 3) % 4 // 當指令為左轉時，加3後除4的餘數
		} else if v == -1 {
			cur = (cur + 1) % 4 // 當指令為右轉時，加1後除4的餘數
		} else {
			for i := 0; i < v; i++ {
				xx := x + dir[cur][0]
				yy := y + dir[cur][1]

				// 如果遇到障礙物就中斷，否則 x,y = xx, yy
				if _, ok := obs[[2]int{xx, yy}]; ok {
					break
				} else {
					x, y = xx, yy
				}
			}
			// 如果當下位置的面積大於最大行徑面積，更新最大行徑面積
			if x*x+y*y > res {
				res = x*x + y*y
			}
		}
	}
	return res
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/walking-robot-simulation/discuss/580982/go-clean-code-44ms-beats-100

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	data := make(map[string]struct{})
	for _, obstacle := range obstacles {
		data[strconv.Itoa(obstacle[0])+" "+strconv.Itoa(obstacle[1])] = struct{}{}
	}
	var ret, x, y, i int
	for _, cmd := range commands {
		switch cmd {
		case -1:
			i = (i + 1) % 4
		case -2:
			i = (i + 3) % 4
		default:
			for ; cmd > 0; cmd-- {
				if _, ok := data[strconv.Itoa(x+dirs[i][0])+" "+strconv.Itoa(y+dirs[i][1])]; ok {
					break
				}
				x, y = x+dirs[i][0], y+dirs[i][1]

			}
			ret = int(math.Max(float64(ret), float64(x*x+y*y)))
		}
	}
	return ret
}
*/
