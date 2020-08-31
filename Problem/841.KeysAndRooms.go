package Problem

//leetcode submit region begin(Prohibit modification and deletion)
func canVisitAllRooms(rooms [][]int) bool {
	isVisited := make([]bool, len(rooms))

	isVisited = getKeys(0, rooms, isVisited)

	for _, val := range isVisited {
		if val == false {
			return false
		}
	}
	return true
}

func getKeys(key int, rooms [][]int, isVisited []bool) []bool {
	isVisited[key] = true
	for _, val := range rooms[key] {
		if !isVisited[val] {
			isVisited[val] = true
			isVisited = getKeys(val, rooms, isVisited)
		}
	}
	return isVisited
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-31 09:51:34

//There are N rooms and you start in room 0. Each room has a distinct number in
//0, 1, 2, ..., N-1, and each room may have some keys to access the next room.
//
// Formally, each room i has a list of keys rooms[i], and each key rooms[i][j] i
//s an integer in [0, 1, ..., N-1] where N = rooms.length. A key rooms[i][j] = v o
//pens the room with number v.
//
// Initially, all the rooms start locked (except for room 0).
//
// You can walk back and forth between rooms freely.
//
// Return true if and only if you can enter every room.
//
//
//
//
// Example 1:
//
//
//Input: [[1],[2],[3],[]]
//Output: true
//Explanation:
//We start in room 0, and pick up key 1.
//We then go to room 1, and pick up key 2.
//We then go to room 2, and pick up key 3.
//We then go to room 3.  Since we were able to go to every room, we return true.
//
//
//
// Example 2:
//
//
//Input: [[1,3],[3,0,1],[2],[0]]
//Output: false
//Explanation: We can't enter the room with number 2.
//
//
// Note:
//
//
// 1 <= rooms.length <= 1000
// 0 <= rooms[i].length <= 1000
// The number of keys in all rooms combined is at most 3000.
//
//
// 👍 110 👎 0
