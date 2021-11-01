# intro

Floyd-Warshall(佛洛依德演算法)，是解決任意兩點間的`最短路徑`的一種演算法，可以正確處理`有向圖`或負權(但不可存在負權迴路)的最短路徑問題，圖繩也被用於計算有向圖的遞移封包

演算法的時間複雜度為O(N^3)，空間複雜度為O(N^2)

# 演算法描述
假設 D_(i,j,k) 表示從 i 到 j 的只以 (1..k) 集合中的節點為中間節點的最短路徑的長度
1. 若最短路徑經過點 k ，則 D_(i,j,k) = D_(i,k,k-1) + D_(k,j,k-1);
2. 若最短路徑不經過點 k ，則 D_(i,j,k) = D_(i,j,k-1)
因此，D_(i,j,k) = min( D_(i,j,k-1), D_(i,k,k-1) + D_(k,j,k-1) )

# code flow
```go
// let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)
// for each vertex v
//    dist[v][v] ← 0

// for each edge (u,v)
//    dist[u][v] ← w(u,v)  // the weight of the edge (u,v)

// for k from 1 to |V|
//    for i from 1 to |V|
//       for j from 1 to |V|
//          if dist[i][j] > dist[i][k] + dist[k][j] 
//              dist[i][j] ← dist[i][k] + dist[k][j]
//          end if

// 其中dist[i][j]表示由點 i 到點 j 的代價，當其為 ∞ 表示兩點之間沒有任何連接。
```

# refer:
- Floyd-Warshall演算法
https://zh.wikipedia.org/wiki/Floyd-Warshall%E7%AE%97%E6%B3%95

- 動態規劃
https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92

# extend-refer:
- 最長公共子序列
https://zh.wikipedia.org/wiki/%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97

- Viterbi演算法
https://zh.wikipedia.org/wiki/%E7%BB%B4%E7%89%B9%E6%AF%94%E7%AE%97%E6%B3%95