# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        if not head or left == right:
            return head
        
        dummy = ListNode(0)
        prev = dummy
        prev.next = head
        
        for _ in range(left-1):
            prev = prev.next
        
        prev_sub = prev.next
        curr = prev.next.next
        for _ in range(right-left):
            next_node = curr.next
            curr.next = prev_sub
            prev_sub = curr
            curr = next_node
        
        prev.next.next = curr
        prev.next = prev_sub

        return dummy.next