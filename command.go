package pract4

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

func (c Calculator) Redo() float64 {
	c.cu.Redo()
	return c.au.Register
}

func (c Calculator) Undo() float64 {
	c.cu.Undo()
	return c.au.Register
}
