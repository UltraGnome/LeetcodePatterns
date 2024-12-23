### Top K Elements

You can always sort an array and then take the first or last k elements, however the time 
complexity would be O(NLogN). A heap can pop and push elements in O(Log(K)) where K is the 
size of the heap. Therefore, instead of sorting, we can use a heap to hold the smallest or 
largest K values, and for every element in the array check whether to pop/push to the heap, 
resulting in a time complexity of O(NLogK).


When to use it?
Find the top k smallest or largest elements

Find the kth smallest or largest element

Find the k most frequent elements

LeetCode Questions
215.[ Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)

347. [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)

23. [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/description/)
