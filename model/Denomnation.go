package model

type Denomination struct {
	FiveHundred int `json:"500"`
	Hundred     int `json:"100"`
	Fifty       int `json:"50"`
	Twenty      int `json:"20"`
	Ten         int `json:"10"`
	Five        int `json:"5"`
	Coins       int `json:"coins"`
}

func (d *Denomination) GetTotal() int {
	fiveHundred := d.GetFiveHundredValue()
	hundred := d.GetHundredValue()
	fifty := d.GetFiftyValue()
	twenty := d.GetTwentyValue()
	ten := d.GetTensValue()
	five := d.GetFiveValue()

	return fiveHundred + hundred + fifty + twenty + ten + five + d.Coins
}

func (d *Denomination) GetFiveHundredValue() int {
	return d.FiveHundred * 500
}

func (d *Denomination) GetHundredValue() int {
	return d.Hundred * 100
}

func (d *Denomination) GetFiftyValue() int {
	return d.Fifty * 50
}

func (d *Denomination) GetTwentyValue() int {
	return d.Twenty * 20
}

func (d *Denomination) GetTensValue() int {
	return d.Ten * 10
}

func (d *Denomination) GetFiveValue() int {
	return d.Five * 5
}
