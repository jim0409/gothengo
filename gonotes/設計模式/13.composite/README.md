# composite 組合 模式

組合模式統一對象和對象集，使得相同接口使用對象和對象集

組合模式常用於樹狀結構，用於統一葉子節點和樹節點的訪問，並且可以用於應用某一操作到所有子節點。

> 使用情境: 在惠團軟體中，使用者可以畫單垂的圖形外，也能夠建立複雜的徒刑，使用者可以把大量的單純`圖性`組成群組的方式來建造複雜的圖形。
e.g. Line、Text這樣的圖元物件，再加上一些物件係上這些圖源物件的容器。而這些容器所`使用的方法`，抽象出來後即為Composite的精神


1. 使用戶端呼叫簡單，用戶端可以一致的使用組合結構或其中單個物件，用戶就不必關係自己處理的事單個物件還是整個組合結構
2. 更容易在組合體內加入物件部件。用戶端不必因為加入了新的物件部件而更改代碼

```原本情境
Application ---> Canvas 
                   |
                   |
        ------------------------ 
        |          |           |
      Text      Rectangle    Line
     +draw()    +draw()      +draw()
```

```使用了Composite
Application ---> Graphic   <-----------------------------------------
                 +draw()                                            |
                    |                                               |
                    |                                               |
         ---------------------------------------------              |
         |          |              |                 |              |
       Line        Rectangle     Text            GroupGraphic  <-----
       +draw()     +draw()       +draw()         +draw()
                                                 +add(Graphic)
                                                 +remove(Graphic)
                                                 +getChild(int)
```

```只看Composite抽象
[[client]]          [[abstract]]
Application ---> Graphic<:Component>  <-------------------------------
                 +draw()<:operation>                                 |
                 +add(Graph<:Component>)                             |
                 +remove(Graph<:Component>)                          |
                 +getChild(Graph<:Component>)                        |
                         |                                           |
                         |                                           |
                ------------------------                             |
                |                       |                            |
            Leaf                 GroupGraphic<:Composite>  <-----(+children)
            +operation()         +draw()<:operation>
                                    +add(Graphic<:Component>)
                                    +remove(Graphic<:Component>)
                                    +getChild(int)
```

- Component
1. 描述在Composite中物件的interface
2. 實作一些共通的部分
3. 定義存取、管理其Child物件的介面(*實作上，可能不一定會有這項)

- Leaf(Primitive, Single, Part, Containee)
1. 表示在Composite中Leaf物件，這樣的物件沒有Child
2. 定義單一特定的物件行為

- Composite(Group, Whole, Container, Branch)
1. 定義有Child物件的行為
2. 儲存Child元件
3. 實作Child相關的物件

- Client
透過Component所定義的Interface，操作Composition物件。

# refer:
- http://www.dotspace.idv.tw/Patterns/Jdon_Composite.htm
責任鍊與組合模式的差異
- https://ithelp.ithome.com.tw/articles/10208172
- https://ithelp.ithome.com.tw/articles/10218208

# keynotes

```go
1. 定義出Component的幾種方法

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName() string
	AddChild(Component)
	Print(string)
}

2. 定義出一些基本單位元，方便後面迭代使用，

const {
    LeafNode = iota
    CompositeNode
}

type component struct {
    parent Component
    name string
}

type Leaf struct {
	component
}

type Composite struct {
    component
    childs []Component
}

3. 定義出基本單位元component需要的方法Parent()/ SetParent(parent Component)/ AddChild(Component)/ Name()/ SetName(name string)/ Print(string)

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) AddChild(Component) {}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) Print(string) {}

4. 需要定義出一個生成Leaf的函數，以及顯示該Leaf值的方法

func NewLeaf() *Leaf {
    return &Leaf{}
}

func (c *Leaf) Print(pre string) {
    fmt.Printf("%s-%s\n", pre, c.Name())
}

5. 定義出Composite來實現後面要新增的Leaf

func NewComposite() *Composite {
    return &Composite{
        childs: make([]Component, 0),
    }
}

func (c *Composite) Print(pre string) {
    fmt.Printf("%s+%s\n", pre, c.Name())
    pre += " "
    for _, comp := range c.childs {
        comp.Print(pre)
    }
}

```

# extend-refer:
composite with dependency-injection
- https://stackoverflow.com/questions/30136414/composite-pattern-and-dependency-injection