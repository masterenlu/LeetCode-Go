package Problem

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 && (p1 != nil || p2 != nil) {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-06-04 09:27:59

// Given the heads of two singly linked-lists headA and headB, return the node at
// which the two lists intersect. If the two linked lists have no intersection at
// all, return null.
//
// For example, the following two linked lists begin to intersect at node c1:
//
// It is guaranteed that there are no cycles anywhere in the entire linked struc
// ture.
//
// Note that the linked lists must retain their original structure after the fun
// ction returns.
//
//
// Example 1:
//
//
// Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2
// , skipB = 3
// Output: Intersected at '8'
// Explanation: The intersected node's value is 8 (note that this must not be 0 i
// f the two lists intersect).
// From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [
// 5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 no
// des before the intersected node in B.
//
//
// Example 2:
//
//
// Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skip
// B = 1
// Output: Intersected at '2'
// Explanation: The intersected node's value is 2 (note that this must not be 0 i
// f the two lists intersect).
// From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [
// 3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node bef
// ore the intersected node in B.
//
//
// Example 3:
//
//
// Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// Output: No intersection
// Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it r
// eads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, whi
// le skipA and skipB can be arbitrary values.
// Explanation: The two lists do not intersect, so return null.
//
//
//
// Constraints:
//
//
// The number of nodes of listA is in the m.
// The number of nodes of listB is in the n.
// 0 <= m, n <= 3 * 104
// 1 <= Node.val <= 105
// 0 <= skipA <= m
// 0 <= skipB <= n
// intersectVal is 0 if listA and listB do not intersect.
// intersectVal == listA[skipA + 1] == listB[skipB + 1] if listA and listB inter
// sect.
//
//
//
// Follow up: Could you write a solution that runs in O(n) time and use only O(1)
// memory?
// 👍 1162 👎 0
