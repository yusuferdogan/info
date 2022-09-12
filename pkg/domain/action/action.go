package action

type Action struct{
	id		string
	actionType string
	category string

}

func NewAction(id string, actionType string, category string) *Action{
	return &Action{
		id: id,
		actionType: actionType,
		category: category,
	}
}

func (a Action) Id() string{
	return a.id
}
func (a Action) ActionType() string{
	return a.actionType
}
func (a Action) Category() string{
	return a.category
}
