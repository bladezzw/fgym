type Array struct {
	Ls     []float64
	Length int
}

func (this *Array) SetLength(l int) int {
	this.Length = l
	if this.len() > l {
		this.Ls = this.Ls[:l]
	}
	return l
}

func (this *Array) len() int {
	return len(this.Ls)
}
func (this *Array) max() (max float64) {
	max = 0
	for i := 0; i < len(this.Ls); i++ {
		if this.Ls[i] > max {
			max = this.Ls[i]
		}
	}
	return max
}
func (this *Array) min() (min float64) {
	min = 0
	for i := 0; i < len(this.Ls); i++ {
		if this.Ls[i] < min {
			min = this.Ls[i]
		}
	}
	return min
}

func (this *Array) append(f float64) {
	this.Ls = append(this.Ls, f)
	if this.len() > this.Length {
		this.Ls = this.Ls[1:]
	}
}
