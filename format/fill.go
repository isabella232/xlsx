package format

import (
	"github.com/plandem/xlsx/format/internal/color"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/styles"
)

type patternOption byte
type gradientOption byte

type fillOption struct {
	Pattern  patternOption
	Gradient gradientOption
}

//Fill is a 'namespace' for all possible settings for fill
var Fill fillOption

func (f *fillOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		f.Pattern.Color(rgb)
	}
}

func (f *fillOption) Background(rgb string) option {
	return func(s *StyleFormat) {
		f.Pattern.Background(rgb)
	}
}

func (f *fillOption) Type(pt styles.PatternType) option {
	return func(s *StyleFormat) {
		f.Pattern.Type(pt)
	}
}

func (p *patternOption) Color(rgb string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Color = color.New(rgb)
	}
}

func (p *patternOption) Background(rgb string) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Background = color.New(rgb)
	}
}

func (p *patternOption) Type(pt styles.PatternType) option {
	return func(s *StyleFormat) {
		s.Fill.Pattern.Type = pt
	}
}

func (g *gradientOption) Type(gt styles.GradientType) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Type = gt
	}
}

func (g *gradientOption) Degree(degree float64) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Degree = degree
	}
}

func (g *gradientOption) Left(left float64) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Left = left
	}
}

func (g *gradientOption) Right(right float64) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Right = right
	}
}

func (g *gradientOption) Top(top float64) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Top = top
	}
}

func (g *gradientOption) Bottom(bottom float64) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Bottom = bottom
	}
}

func (g *gradientOption) Stop(position float64, rgb string) option {
	return func(s *StyleFormat) {
		s.Fill.Gradient.Stop = append(s.Fill.Gradient.Stop, &ml.GradientStop{Position: position, Color: color.New(rgb)})
	}
}
