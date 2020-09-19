package pract4

import "fmt"

type EventHandler struct {
	subscribers []func(interface{})
}

func (e *EventHandler) AddHandler(handler func(interface{})) {
	e.subscribers = append(e.subscribers, handler)
}

func (e EventHandler) NotifyAll(sender interface{}) {
	for _, s := range e.subscribers {
		s(sender)
	}
}

type Drive struct {
	EventHandler
	Twist string
}

func (d *Drive) SetTwist(twist string) {
	d.Twist = twist
	d.NotifyAll(d)
}

func (d *Drive) Drive() {
	d.SetTwist("исходная позиция")
}

func (d *Drive) TurnLeft() {
	d.SetTwist("поворот налево")
}

func (d *Drive) TurnRight() {
	d.SetTwist("поворот направо")
}

func (d *Drive) Stop() {
	d.SetTwist("стоп")
}

func (d Drive) String() string {
	return fmt.Sprintf("Привод: %v", d.Twist)
}

type Power struct {
	EventHandler
	MicrowavePower int
}

func (p *Power) SetMicrowavePower(power int) {
	p.MicrowavePower = power
	p.NotifyAll(p)
}

func (p Power) String() string {
	return fmt.Sprintf("Задана мощность %vw", p.MicrowavePower)
}

type Notification struct {
	EventHandler
	Message string
}

func (n *Notification) SetMessage(message string) {
	n.Message = message
	n.NotifyAll(n)
}

func (n *Notification) StartNotification() {
	n.SetMessage("Операция началась")
}

func (n *Notification) StopNotification() {
	n.SetMessage("Операция закончилась")
}

func (n Notification) String() string {
	return fmt.Sprintf("Информация: %v", n.Message)
}

type Microwave struct {
	drive        *Drive
	power        *Power
	notification *Notification
}

func (m *Microwave) Defrost() {
	m.notification.StartNotification()
	m.power.SetMicrowavePower(1000)
	m.drive.TurnRight()
	m.drive.TurnRight()
	m.power.SetMicrowavePower(500)
	m.drive.Stop()
	m.drive.TurnLeft()
	m.drive.TurnLeft()
	m.power.SetMicrowavePower(200)
	m.drive.Stop()
	m.drive.TurnRight()
	m.drive.TurnRight()
	m.drive.Stop()
	m.power.SetMicrowavePower(0)
	m.notification.StopNotification()
}
