package lang

type EvalContextFrame struct {
	CurrentLocation int
	ShouldEvalLocation func(int) bool
}


type EvalContext struct {
	Frames []*EvalContextFrame
}

type Form int
const (
	StandardForm = iota
	DefForm
	LambdaForm
)

func StandardFormFunc(position int) bool {
	return true
}

func DefFormFunc(position int) bool {
	if position == 1 {
		return false // this is the variable name.
	} else {
		return true
	}
}

func LambdaFormFunc(position int) bool {
	return false
}

func GetFormFunc(form Form) func(int)bool {
	switch form {
	case DefForm:
		return DefFormFunc
	default:
		return StandardFormFunc
	}
}

func (ec *EvalContext) Top() *EvalContextFrame {
	return ec.Frames[len(ec.Frames)-1]
}

func (ec *EvalContext) Push(form Form) {
	ec.Frames = append(ec.Frames, 
		&EvalContextFrame{
			// This is a normal context. Eval everything
			ShouldEvalLocation: GetFormFunc(form),
		})
}

func (ec *EvalContext) Pop() {
	ec.Frames = ec.Frames[:len(ec.Frames)-1]
}

func (ecf *EvalContextFrame) ShouldEval() bool {
	return ecf.ShouldEvalLocation(ecf.CurrentLocation)
}

func (ecf *EvalContextFrame) MoveNext() {
	ecf.CurrentLocation++
}

