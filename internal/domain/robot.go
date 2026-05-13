package domain

type RobotID int

type Robot struct {
	ID            RobotID
	Name          string
	Weight        float64
	Width         float64
	Height        float64
	Length        float64
	IsValid       bool
	InvalidReason string
	Autonomous    string
	PowerButton   string
	InternalPower string
	Status        string
	TeamID        TeamID
	CategoryID    CategoryID
}

func NewRobot(
	name string,
	weight float64,
	width float64,
	height float64,
	length float64,
	autonomous string,
	powerButton string,
	internalPower string,
	teamID TeamID,
	categoryID CategoryID,
) (*Robot, error) {
	// Verificar nombre vacio
	if name == "" {
		return nil, ErrEmpty
	}

	if weight <= 0 {
		return nil, ErrInvalidWeight
	}

	if width <= 0 || height <= 0 || length <= 0 {
		return nil, ErrInvalidDimensions
	}

	if teamID <= 0 {
		return nil, ErrInvalidID
	}

	if categoryID <= 0 {
		return nil, ErrInvalidID
	}

	if powerButton == "" {
		return nil, ErrMissingPowerButton
	}

	return &Robot{
		Name:          name,
		Weight:        weight,
		Width:         width,
		Height:        height,
		Length:        length,
		Autonomous:    autonomous,
		PowerButton:   powerButton,
		InternalPower: internalPower,
		IsValid:       false,
		Status:        "pending",
		TeamID:        teamID,
		CategoryID:    categoryID,
	}, nil
}
