package tundra

type Object struct {
	Description   string
	containedObjs map[string]*Object
	commands      map[string]Command
}
type ObjectOption func(*Object)

func NewObject(options ...ObjectOption) *Object {

	o := &Object{}
	for _, option := range options {
		option(o)
	}
	return o
}

func WithDescription(description string) ObjectOption {
	return func(o *Object) {
		o.Description = description
	}
}

func (o *Object) AddCommand(name string, command Command) {
	if o.commands == nil {
		o.commands = make(map[string]Command, 0)
	}
	o.commands[name] = command
}

func (o *Object) GetCommand(name string) Command {
	if o.commands == nil {
		return nil
	}
	return o.commands[name]
}

func (o *Object) RemoveCommand(name string) {
	if o.commands == nil {
		return
	}
	delete(o.commands, name)
}

func (o *Object) AddObject(name string, object *Object) {
	if o.containedObjs == nil {
		o.containedObjs = make(map[string]*Object, 0)
	}
	o.containedObjs[name] = object
}

func (o *Object) GetObject(name string) *Object {
	if o.containedObjs == nil {
		return nil
	}
	return o.containedObjs[name]
}

func (o *Object) RemoveObject(name string) {
	if o.commands == nil {
		return
	}
	delete(o.containedObjs, name)
}
