package interpreter

var EmptyColl = NewEmptyCol()

var colTypeSpec = UndefinedTypeSpec
var colTypeInfo = NewListTypeInfo(nil)

type ColAccessor interface {
	AnyAccessor
	ItemTypeSpec() TypeSpecAccessor
	Empty() bool
	Count() int
	Get(i int) interface{}
	Contains(item interface{}) bool
}

type ColModifier interface {
	ColAccessor
	Add(item interface{})
	AddUnique(item interface{}) bool
	AddAll(collection ColAccessor) int
	AddAllUnique(collection ColAccessor) int
}

type baseColType struct {
	baseAnyType
	adapter      ModelAdapter
	itemTypeSpec TypeSpecAccessor
	items        []interface{}
}

type colType struct {
	baseColType
}

type sysArrayCol struct {
	baseColType
}

type emptyCol struct {
	baseAnyType
}

type colDelegate struct {
	delegate ColAccessor
}

func IsCol(node interface{}) bool {
	_, ok := node.(ColAccessor)
	return ok
}

func NewCol(adapter ModelAdapter) ColModifier {
	return NewColWithSource(adapter, nil)
}

func NewColWithItem(adapter ModelAdapter, item interface{}) ColModifier {
	c := newCollection(adapter, nil, nil)
	c.items = make([]interface{}, 1)
	c.items[0] = item
	return c
}

func NewColWithSource(adapter ModelAdapter, source interface{}) ColModifier {
	return newCollection(adapter, nil, source)
}

func NewColWithSpec(adapter ModelAdapter, itemTypeSpec TypeSpecAccessor) ColModifier {
	return newCollection(adapter, itemTypeSpec, nil)
}

func NewEmptyCol() ColAccessor {
	return NewEmptyColWithSource(nil)
}

func NewEmptyColWithSource(source interface{}) ColAccessor {
	return &emptyCol{baseAnyType{source}}
}

func NewColDelegate(delegate ColAccessor) ColAccessor {
	return &colDelegate{delegate}
}

func newCollection(adapter ModelAdapter, itemTypeSpec TypeSpecAccessor, source interface{}) *colType {
	if adapter == nil {
		panic("no adapter has been provided")
	}

	return &colType{
		baseColType{
			baseAnyType{source},
			adapter,
			itemTypeSpec,
			nil,
		},
	}
}

func NewSysArrayCol(itemTypeSpec TypeSpecAccessor, items []interface{}) ColAccessor {
	return &sysArrayCol{
		baseColType{
			itemTypeSpec: itemTypeSpec,
			items:        items,
		},
	}
}

func (c *baseColType) DataType() DataTypes {
	return ColDataType
}

func (c *baseColType) TypeSpec() TypeSpecAccessor {
	return colTypeSpec
}

func (c *baseColType) TypeInfo() TypeInfoAccessor {
	if c.itemTypeSpec == nil {
		return colTypeInfo
	}
	return NewListTypeInfo(NewString(c.itemTypeSpec.String()))
}

func (c *baseColType) ItemTypeSpec() TypeSpecAccessor {
	return c.itemTypeSpec
}

func (c *baseColType) Empty() bool {
	return c.Count() == 0
}

func (c *baseColType) Count() int {
	return len(c.items)
}

func (c *baseColType) Get(i int) interface{} {
	if c.items == nil {
		panic("collection is empty")
	}
	return c.items[i]
}

func (c *baseColType) Contains(item interface{}) bool {
	if c.items == nil {
		return false
	}

	for _, o := range c.items {
		if ModelEqual(c.adapter, item, o) {
			return true
		}
	}
	return false
}

func (c *baseColType) Equal(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			colDeepEqual(c.adapter, c, o)
	}
}

func (c *baseColType) Equivalent(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			colDeepEquivalent(c.adapter, c, o)
	}
}

func colDeepEqual(adapter ModelAdapter, c1 ColAccessor, c2 ColAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEqual(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func colDeepEquivalent(adapter ModelAdapter, c1 ColAccessor, c2 ColAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEquivalent(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func (c *colType) Add(item interface{}) {
	c.add(item)
}

func (c *colType) add(item interface{}) {
	if c.items == nil {
		c.items = make([]interface{}, 0)
	}
	c.items = append(c.items, item)
}

func (c *colType) AddUnique(item interface{}) bool {
	if c.items == nil {
		c.add(item)
		return true
	}

	if item == nil {
		for _, o := range c.items {
			if o == nil {
				return false
			}
		}
	} else {
		adapter := c.adapter
		for _, o := range c.items {
			if o != nil && ModelEqual(adapter, item, o) {
				return false
			}
		}
	}

	c.items = append(c.items, item)
	return true
}

func (c *colType) AddAll(collection ColAccessor) int {
	count := collection.Count()
	for i := 0; i < count; i++ {
		c.add(collection.Get(i))
	}
	return count
}

func (c *colType) AddAllUnique(collection ColAccessor) int {
	added := 0
	count := collection.Count()
	for i := 0; i < count; i++ {
		a := c.AddUnique(collection.Get(i))
		if a {
			added = added + 1
		}
	}
	return added
}

func (c *emptyCol) DataType() DataTypes {
	return ColDataType
}

func (c *emptyCol) TypeSpec() TypeSpecAccessor {
	return colTypeSpec
}

func (c *emptyCol) TypeInfo() TypeInfoAccessor {
	return colTypeInfo
}

func (c *emptyCol) Equal(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return o.Empty()
	}
}

func (c *emptyCol) Equivalent(item interface{}) bool {
	return c.Equal(item)
}

func (c *emptyCol) ItemTypeSpec() TypeSpecAccessor {
	return UndefinedTypeSpec
}

func (c *emptyCol) Empty() bool {
	return true
}

func (c *emptyCol) Count() int {
	return 0
}

func (c *emptyCol) Get(int) interface{} {
	panic("cannot get an item from an empty collection")
}

func (c *emptyCol) Contains(interface{}) bool {
	return false
}

func (c *colDelegate) DataType() DataTypes {
	return c.delegate.DataType()
}

func (c *colDelegate) TypeSpec() TypeSpecAccessor {
	return c.delegate.TypeSpec()
}

func (c *colDelegate) TypeInfo() TypeInfoAccessor {
	return c.delegate.TypeInfo()
}

func (c *colDelegate) Source() interface{} {
	return c.delegate.Source()
}

func (c *colDelegate) Equal(node interface{}) bool {
	return c.delegate.Equal(node)
}

func (c *colDelegate) Equivalent(node interface{}) bool {
	return c.delegate.Equivalent(node)
}

func (c *colDelegate) ItemTypeSpec() TypeSpecAccessor {
	return c.delegate.ItemTypeSpec()
}

func (c *colDelegate) Empty() bool {
	return c.delegate.Empty()
}

func (c *colDelegate) Count() int {
	return c.delegate.Count()
}

func (c *colDelegate) Get(i int) interface{} {
	return c.delegate.Get(i)
}

func (c *colDelegate) Contains(item interface{}) bool {
	return c.delegate.Contains(item)
}
