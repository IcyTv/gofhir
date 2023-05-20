package interpreter

type loop struct {
	evaluator Evaluator
	this      interface{}
	index     int
	total     interface{}
}

type Looper interface {
	Evaluator() Evaluator
	This() interface{}
	Index() int
	IncIndex(this interface{}) int
	Total() interface{}
	SetTotal(total interface{})
}

func NewLoop(evaluator Evaluator) Looper {
	return &loop{evaluator, nil, -1, nil}
}

func NewLoopWithIndex(evaluator Evaluator, index int) Looper {
	return &loop{evaluator, nil, index - 1, nil}
}

func (c *loop) Evaluator() Evaluator {
	return c.evaluator
}

func (c *loop) This() interface{} {
	if c.index < 0 {
		panic("Index not yet set")
	}
	return c.this
}

func (c *loop) Index() int {
	return c.index
}

func (c *loop) IncIndex(this interface{}) int {
	c.this = this
	c.index++
	return c.index
}

func (c *loop) Total() interface{} {
	return c.total
}

func (c *loop) SetTotal(total interface{}) {
	c.total = total
}

type Evaluator interface {
	Evaluate(ctx ContextAccessor, node interface{}, loop Looper) (interface{}, error)
}

type BaseFunction struct {
	name           string
	evaluatorParam int
	minParams      int
	maxParams      int
}

type FunctionExecutor interface {
	Name() string
	EvaluatorParam() int
	MinParams() int
	MaxParams() int
	Execute(ctx ContextAccessor, node interface{}, args []interface{}, loop Looper) (interface{}, error)
}

func NewBaseFunction(name string, evaluatorParam int, minParams int, maxParams int) BaseFunction {
	return BaseFunction{name, evaluatorParam, minParams, maxParams}
}

func (f *BaseFunction) Name() string {
	return f.name
}

func (f *BaseFunction) EvaluatorParam() int {
	return f.evaluatorParam
}

func (f *BaseFunction) MinParams() int {
	return f.minParams
}

func (f *BaseFunction) MaxParams() int {
	return f.maxParams
}
