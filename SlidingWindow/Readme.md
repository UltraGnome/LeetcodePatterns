Sliding Window 

This pattern uses two pointers, P1 and P2

Move P2 as far as possible until the condition is false.

Move P1 closer to P2 until the condition is true.

Keep track of min and max length of the subarray.

When to use:
1.  Linear data structures (arrays, lists, strings)
2. If you must scan a subarray or string.
3. If subarray must satisfy a condition
4. Improve time complexity?

Big O:  O(n)

LeetCode Questions
3.[ Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

424. [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/description/)

[1876. Substrings of Size Three with Distinct Characters](https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/description/)

76. [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/) HARD