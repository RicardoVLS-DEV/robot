package domain

type CategoryID int

type Category struct {
	ID		CategoryID
	Name 	string
	RuleID	int
}

func NewCategory(name string,  ruleID int) (*Category, error) {
	if name == "" {
		return nil, ErrEmpty
	}

	return &Category{
		Name: name,
		RuleID: ruleID,
	}, nil
}

func (c *Category) AddRule(ruleID int) {
	c.RuleID = ruleID
}