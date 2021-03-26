## Topics covered

- Divide and conquer
- Graphs and trees
- Depth-first search
- Topological sort; strongly-connected components
- Breadth-first search
- Shortest paths; Dijkstra and Bellman-Ford
- Minimum spanning trees
- Union/find analysis
- Hufmann codes
- Lempel-Ziv codes
- Randomized min-cut
- Hashing
- Bloom filters
- Dynamic programming
- Linear programming; posing of combinatorial problems as LP problems
- Duality
- Network flows
- NP completeness
- Approximation algorithms
- Fast Fourier transform

## Three-Part Algorithm Format

Oftentimes, a problem will ask you to “give/design/devise/formulate/create/etc. an algorithm” or to “show how” to solve some computational task. In this case, write your solution in the 3-part algorithm format:

1. **Algorithm description**

   This can come in terms of pseudocode, or a description in English. It must be unambiguous, as short as possible (but no shorter), and precise.

   - Your pseudocode does not need to be executable. You should use notation such as “add _X_ to set _S_” or “for each edge in graph G”. Remember you are writing your pseudocode to be read by a human, not a computer.
   - See DPV for examples of pseudocode.

2. **Proof of correctness**

   Give a formal proof (as in CS 70) of the correctness of your algorithm. Intuitive arguments are not enough.

   - Again, see DPV for examples of proofs of correctness.

3. **Runtime analysis**

   You should use big-O notation for your algorithm’s runtime, and justify this runtime with a runtime analysis. This may involve a recurrence relation, or simply counting the complexity and number of operations your algorithm performs.

### Divide-and-conquer

Both merge sort and quicksort employ a common algorithmic paradigm based on recursion. This paradigm, **divide-and-conquer**, breaks a problem into subproblems that are similar to the original problem, recursively solves the subproblems, and finally combines the solutions to the subproblems to solve the original problem. Because divide-and-conquer solves subproblems recursively, each subproblem must be smaller than the original problem, and there must be a base case for subproblems. You should think of a divide-and-conquer algorithm as having three parts:

1. **Divide** the problem into a number of subproblems that are smaller instances of the same problem.
2. **Conquer** the subproblems by solving them recursively. If they are small enough, solve the subproblems as base cases.
3. **Combine** the solutions to the subproblems into the solution for the original problem.

You can easily remember the steps of a divide-and-conquer algorithm as _divide, conquer, combine_. Here's how to view one step, assuming that each divide step creates two subproblems (though some divide-and-conquer algorithms create more than two):

![img](https://cdn.kastatic.org/ka-perseus-images/98c02634ee7f970a6bfb0812cc1495bacb462282.png)

If we expand out two more recursive steps, it looks like this:

![img](https://cdn.kastatic.org/ka-perseus-images/db9d172fc33b90e905c1213b8cce660c228bb99c.png)

Because divide-and-conquer creates at least two subproblems, a divide-and-conquer algorithm makes multiple recursive calls.

