package gototo

type Draw interface {
	GetDate() string
	GetWinningNumbers() []int
	GetAdditionalNumber() int
}

type DrawModel struct {
	Date             string
	WinningNumbers   []int
	AdditionalNumber int
}

func (r DrawModel) GetDate() string {
	return r.Date
}

func (r DrawModel) GetWinningNumbers() []int {
	return r.WinningNumbers
}

func (r DrawModel) GetAdditionalNumber() int {
	return r.AdditionalNumber
}

type NextDraw interface {
	GetDate() string
	GetPrize() string
}

type NextDrawModel struct {
	Date  string
	Prize string
}

func (n NextDrawModel) GetDate() string {
	return n.Date
}

func (n NextDrawModel) GetPrize() string {
	return n.Prize
}
