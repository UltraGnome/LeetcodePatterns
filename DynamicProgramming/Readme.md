### Dynamic Programming

Dynamic Programming is used when you need to solve a problem that depends on previous results from subproblems. 
You can effectively “cache” these previous result values when you calculate them for the first time to be re-used
later. 

Dynamic Programming has 2 main techniques:

Top Down - Recursion (DFS) with Memoization. Memoization is a fancy word for a hashmap that can cache the values 
previously calculated. In the top down approach you start with the global problem and the recursively split it 
into subproblems to then solve the global problem.

Bottom Up - Iteratively performed by using an array/matrix to store previous values. In the bottom up approach we 
start with base cases and then build up to the global solution iteratively.

Many times bottom up is preferred since you can reduce the space complexity if you don’t need access to all 
subproblems and can store the last couple of subproblem results using variables.

When to use it?
Overlapping subproblems and optimal substructure

Optimization problems (min/max distance, profit, etc.)

Sequence problems (longest increasing subsequence)

Combinatorial problems (number of ways to do something)

Reduce time complexity from exponential to polynomial

LeetCode Questions
70. [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)

322. [Coin Change](https://leetcode.com/problems/coin-change/description/)

1143. [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/description/)

300. [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/description/)

72.[ Edit Distance](https://leetcode.com/problems/edit-distance/description/)
