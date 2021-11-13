# intro
Binary Tree: 左邊大於父節點，右邊小於父節點
```
            A(4)
        /         \
    B(2)          C(5)
    /   \         /
D(1)    E(8)     F(3)
        / \       \
    G(7)  H(9)     I(6)
```

# 引薦之前python學的筆記
- PreOrder/ InOrder/ PostOrder
```python
# ABC, Preorder : root to left to right (根 -> 左子樹 -> 右子樹) .. 從頂點開始向左尋，如果不是一顆完整樹(虛構補齊)
# ABDE -> GH -> CFI
def preorder(ptr):
    if ptr != None:
        print('[%2d]' % ptr.data, end='')
        preorder(ptr.left)
        preorder(ptr.right)

# BAC, Inorder : left to root to right (左子樹 -> 根 -> 右子樹) .. 從最左邊的數字開始，如果不是一顆完整樹(虛構補齊)
# DB -> GE -> H -> A -> F -> I -> C
def inorder(ptr):
    if ptr != None:
        inorder(ptr.left)
        print('[%2d]' % ptr.data, end='')
        inorder(ptr.right)

# BCA, Postorder : left to right to root (左子樹 -> 右子樹 -> 根) .. 從最左邊的數字開始，跳過根向右邊查，如果不是一顆完整樹(虛構補齊)
# D -> GH ->E -> B -> I -> F -> C -> A
def postorder(ptr):
    if ptr !=None:
        postorder(ptr.left)
        postorder(ptr.right)
        print('[%2d]' % ptr.data, end='')
```

# refer:
- https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%8F%89%E6%A0%91
- http://alrightchiu.github.io/SecondRound/binary-tree-traversalxun-fang.html
- https://cloud.tencent.com/developer/article/1412883