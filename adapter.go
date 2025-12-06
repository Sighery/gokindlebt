package gokindlebt

type Adapter struct{}

func NewAdapter() (Adapter, error) {
	return Adapter{}, nil
}
