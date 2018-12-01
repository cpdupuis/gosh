package lang

type EvalContextFrame struct {
	CurrentLocation int
	ShouldEvalLocation func(int) bool
}


type EvalContext struct {
	Frames []*EvalContextFrame
}

func StandardForm(position int) bool {
	return true
}

func DefForm(position int) bool {
	if position == 1 {
		return false // this is the variable name.
	} else {
		return true
	}
}

func (ec *EvalContext) Top() *EvalContextFrame {
	return ec.Frames[len(ec.Frames)-1]
}

func (ec *EvalContext) Push() {
	ec.Frames = append(ec.Frames, 
		&EvalContextFrame{
			// This is a normal context. Eval everything
			ShouldEvalLocation: StandardForm,
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

