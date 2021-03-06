# 生成排列的演算法
### 暴力法(Brute Force)
暴力法是很直接的一種分治法: 先生成`n-1`個元素的排列，加上第`n`個元素即可得到`n`個元素的排列

1. 將第`n`個元素一次交換到最後一個位置上
2. 遞迴生成前`n-1`個元素的排列
3. 加上最後一個元素，即為n個元素的排列

優點:
邏輯簡單直接，易於理解

缺點:
返回的排列數是`n!`，效能的關鍵在於係數的大小。由於暴力法每次回圈都需要交換兩個位置上的元素

，遞迴結束後又需要`再交換回來`，在`n`較大的情況下，效能較差！


### 插入法 (Insert)
插入法顧名思義就是將元素插入到一個序列中所有可能的位置生成新的序列。從第 1 個元素開始。例如要生成 `{1,2,3}` 的排列

1. 先從序列 1 開始，插入元素 2 ，有兩個位置可以插入，生成兩個序列12和21
2. 將3插入這兩個序列的所有可能位置，生成最終的6個序列

優點:
邏輯簡單直接，易於理解

缺點:
依舊消耗效能(但是比起暴力法效能會更好一些)



### 字典法 (Lexicographic)
該演算法有個前提是序列必須是有升序排列的，當然也可以微調對其它序列使用。

它通過修改當前序列得到下一個序列。我們為每個序列定義一個 權重 ，類比序列組成的數字的大小，序列升序排列時 "權重" 最小，

降序排列時 "權重" 最大。e.g. 1234 的排列按**"權重"由小到大：
```
1234
1243
1324
1342
1423
1432
2134
...
```

優點:
效能不錯

缺點:
不易理解


### SJT演算法 (Steinhaus-Johnson-Trotter)
SJT 演算法在前一個排列的基礎上通過`僅交換相鄰的兩個元素`來生成下一個排列。e.g 按照順序生成 123 的排列:
```
123(交換23) ->
132(交換13) ->
312(交換12) ->
321(交換32) ->
231(交換31) ->
213
```
1. 查詢序列中可移動的最大元素。一個元素可移動意味著它的值大於它指向的相鄰元素。
2. 交換該元素與它指向的相鄰元素。
3. 修改所有值大於該元素的元素的方向。
4. 重複以上步驟直到沒有可移動的元素。


### 堆疊法 (Heap)
Heap演算法優雅、高效。它是從`暴力法`(多次交換)演化而來的，堆演算法就是通過減少交換提升效率。
1. 如果元素個數為奇數，交換第一個和最後一個元素
2. 如果元素個數為偶數，一次交換第i個和最後一個元素

優點:
程式碼實現簡單、高效

缺點:
不易離解

# refer:
- https://www.jishuwen.com/d/p21Q/zh-tw
heap permutations
- https://en.wikipedia.org/wiki/Heap%27s_algorithm