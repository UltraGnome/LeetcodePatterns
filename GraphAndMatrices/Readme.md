### Graph and Matrices

DFS (Depth-First Search) traverses as deep as possible along each branch before backtracking, 
prioritizing visiting nodes or cells in a recursive or stack-based manner. 
BFS (Breadth-First Search) explores all neighbors of a node or cell before moving deeper, 
traversing level by level using a queue. 
For DFS use recursion with a visited set to keep track of visited nodes. 
For BFS use iteration with a queue and a visited set to keep track of visited nodes. 
In a graph, neighbors are found in the adjacency list. 
In a matrix, neighbors are up/down/left/right cells, with some examples including diagonals too.

When to use it?
Search graphs or matrices

DFS: Explore all possible paths (e.g., maze)

BFS: Find the shortest path

Topological Sort: Order tasks based on dependencies

LeetCode Questions
79. [Word Search](https://leetcode.com/problems/word-search/description/)

207. [Course Schedule](https://leetcode.com/problems/course-schedule/description/)

994.[ Rotting Oranges](https://leetcode.com/problems/rotting-oranges/description/)

417. [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/description/)

127. [Word Ladder](https://leetcode.com/problems/word-ladder/description/)
