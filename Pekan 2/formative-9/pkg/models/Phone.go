package models

import "strconv"

// Struct
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// Method
func (p Phone) CetakInfoPhone() string {
	s := "name : " + p.Name + "\n" + "brand : " + p.Brand + "\n" + "year : " + strconv.FormatInt(int64(p.Year), 10) + "\n" + "colors : "
	for i, color := range p.Colors {
		s += color

		if i != len(p.Colors)-1 {
			s += ", "
		}
	}

	return s
}
