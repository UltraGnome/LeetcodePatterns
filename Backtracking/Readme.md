### Backtracking

Backtracking is closely related to DFS, but with a focus on finding solutions while validating their correctness. 
If a solution doesn't work, you backtrack by returning to the previous recursive state and trying a different option. 
Additionally, backtracking uses constraints to eliminate branches that cannot lead to a valid solution, making 
the search more efficient.

When to use it?
Combinatorial problems (combinations, permutations, subsets)

Constraint satisfaction (Sudoku, N-Queens)

Prune paths using constraints to reduce search space

LeetCode Questions
78. [Subsets](https://leetcode.com/problems/subsets/description/)

46. [Permutations](https://leetcode.com/problems/permutations/description/)

39. [Combination Sum](https://leetcode.com/problems/combination-sum/description/)

37.[ Sudoku Solver](https://leetcode.com/problems/sudoku-solver/description/)

51. [N-Queens](https://leetcode.com/problems/n-queens/description/)
