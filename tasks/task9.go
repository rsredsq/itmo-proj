package tasks

type Command interface {
	Execute()
	UnExecute()
}

type ArithmeticUnit struct {
	Register float64
}

func (u *ArithmeticUnit) Run(opCode rune, operand float64) {
	switch opCode {
	case '+':
		u.Register += operand
	case '-':
		u.Register -= operand
		break
	case '*':
		u.Register *= operand
	case '/':
		u.Register /= operand
	}
}

type ControlUnit struct {
	commands []Command
	current  int
}

func (u *ControlUnit) StoreCommand(cmd Command) {
	u.commands = append(u.commands, cmd)
}

func (u *ControlUnit) ExecuteCommand() {
	u.commands[u.current].Execute()
	u.current++
}

func (u *ControlUnit) Undo() {
	u.commands[u.current-1].UnExecute()
}

func (u *ControlUnit) Redo() {
	u.commands[u.current-1].Execute()
}

func (u *ControlUnit) UndoN(n int) {
	commands := u.commands[len(u.commands)-n:]
	for i := len(commands) - 1; i >= 0; i-- {
		commands[i].UnExecute()
	}
}

func (u *ControlUnit) RedoN(n int) {
	for _, cmd := range u.commands[len(u.commands)-n:] {
		cmd.Execute()
	}
}

type Add struct {
	unit    *ArithmeticUnit
	operand float64
}

func (a Add) Execute() {
	a.unit.Run('+', a.operand)
}

func (a Add) UnExecute() {
	a.unit.Run('-', a.operand)
}

type Sub struct {
	unit    *ArithmeticUnit
	operand float64
}

func (s Sub) Execute() {
	s.unit.Run('-', s.operand)
}

func (s Sub) UnExecute() {
	s.unit.Run('+', s.operand)
}

type Mul struct {
	unit    *ArithmeticUnit
	operand float64
}

func (m Mul) Execute() {
	m.unit.Run('*', m.operand)
}

func (m Mul) UnExecute() {
	m.unit.Run('/', m.operand)
}

type Div struct {
	unit    *ArithmeticUnit
	operand float64
}

func (d Div) Execute() {
	d.unit.Run('/', d.operand)
}

func (d Div) UnExecute() {
	d.unit.Run('*', d.operand)
}

type Calculator struct {
	au *ArithmeticUnit
	cu *ControlUnit
}

func NewCalculator() Calculator {
	return Calculator{
		au: &ArithmeticUnit{},
		cu: &ControlUnit{},
	}
}

func (c Calculator) Run(command Command) float64 {
	c.cu.StoreCommand(command)
	c.cu.ExecuteCommand()
	return c.au.Register
}

func (c Calculator) Add(operand float64) float64 {
	return c.Run(Add{c.au, operand})
}

func (c Calculator) Sub(operand float64) float64 {
	return c.Run(Sub{c.au, operand})
}

func (c Calculator) Mul(operand float64) float64 {
	return c.Run(Mul{c.au, operand})
}

func (c Calculator) Div(operand float64) float64 {
	return c.Run(Div{c.au, operand})
}

func (c Calculator) Redo() float64 {
	c.cu.Redo()
	return c.au.Register
}

func (c Calculator) Undo() float64 {
	c.cu.Undo()
	return c.au.Register
}

func (c Calculator) RedoN(n int) float64 {
	c.cu.RedoN(n)
	return c.au.Register
}

func (c Calculator) UndoN(n int) float64 {
	c.cu.UndoN(n)
	return c.au.Register
}
