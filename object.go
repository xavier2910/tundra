package tundra

type Object struct {
	Description   string
	ContainedObjs map[string]*Object
	Commands      map[string]Command
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
	if o.Commands == nil {
		o.Commands = make(map[string]Command, 0)
	}
	o.Commands[name] = command
}

func (o *Object) RemoveCommand(name string) {
	delete(o.Commands, name)
}

func (o *Object) AddObject(name string, object *Object) {
	if o.ContainedObjs == nil {
		o.ContainedObjs = make(map[string]*Object, 0)
	}
	o.ContainedObjs[name] = object
}

func (o *Object) RemoveObject(name string) {
	delete(o.ContainedObjs, name)
}
