package ewo_CM

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ewo_CM struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	percentSuffix          []byte
	perMille               []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
}

// New returns a new instance of translator for the 'ewo_CM' locale
func New() locales.Translator {
	return &ewo_CM{
		locale:                 "ewo_CM",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		currencies:             [][]uint8{[]uint8{0x41, 0x44, 0x50}, []uint8{0x41, 0x45, 0x44}, []uint8{0x41, 0x46, 0x41}, []uint8{0x41, 0x46, 0x4e}, []uint8{0x41, 0x4c, 0x4b}, []uint8{0x41, 0x4c, 0x4c}, []uint8{0x41, 0x4d, 0x44}, []uint8{0x41, 0x4e, 0x47}, []uint8{0x41, 0x4f, 0x41}, []uint8{0x41, 0x4f, 0x4b}, []uint8{0x41, 0x4f, 0x4e}, []uint8{0x41, 0x4f, 0x52}, []uint8{0x41, 0x52, 0x41}, []uint8{0x41, 0x52, 0x4c}, []uint8{0x41, 0x52, 0x4d}, []uint8{0x41, 0x52, 0x50}, []uint8{0x41, 0x52, 0x53}, []uint8{0x41, 0x54, 0x53}, []uint8{0x41, 0x55, 0x44}, []uint8{0x41, 0x57, 0x47}, []uint8{0x41, 0x5a, 0x4d}, []uint8{0x41, 0x5a, 0x4e}, []uint8{0x42, 0x41, 0x44}, []uint8{0x42, 0x41, 0x4d}, []uint8{0x42, 0x41, 0x4e}, []uint8{0x42, 0x42, 0x44}, []uint8{0x42, 0x44, 0x54}, []uint8{0x42, 0x45, 0x43}, []uint8{0x42, 0x45, 0x46}, []uint8{0x42, 0x45, 0x4c}, []uint8{0x42, 0x47, 0x4c}, []uint8{0x42, 0x47, 0x4d}, []uint8{0x42, 0x47, 0x4e}, []uint8{0x42, 0x47, 0x4f}, []uint8{0x42, 0x48, 0x44}, []uint8{0x42, 0x49, 0x46}, []uint8{0x42, 0x4d, 0x44}, []uint8{0x42, 0x4e, 0x44}, []uint8{0x42, 0x4f, 0x42}, []uint8{0x42, 0x4f, 0x4c}, []uint8{0x42, 0x4f, 0x50}, []uint8{0x42, 0x4f, 0x56}, []uint8{0x42, 0x52, 0x42}, []uint8{0x42, 0x52, 0x43}, []uint8{0x42, 0x52, 0x45}, []uint8{0x42, 0x52, 0x4c}, []uint8{0x42, 0x52, 0x4e}, []uint8{0x42, 0x52, 0x52}, []uint8{0x42, 0x52, 0x5a}, []uint8{0x42, 0x53, 0x44}, []uint8{0x42, 0x54, 0x4e}, []uint8{0x42, 0x55, 0x4b}, []uint8{0x42, 0x57, 0x50}, []uint8{0x42, 0x59, 0x42}, []uint8{0x42, 0x59, 0x52}, []uint8{0x42, 0x5a, 0x44}, []uint8{0x43, 0x41, 0x44}, []uint8{0x43, 0x44, 0x46}, []uint8{0x43, 0x48, 0x45}, []uint8{0x43, 0x48, 0x46}, []uint8{0x43, 0x48, 0x57}, []uint8{0x43, 0x4c, 0x45}, []uint8{0x43, 0x4c, 0x46}, []uint8{0x43, 0x4c, 0x50}, []uint8{0x43, 0x4e, 0x58}, []uint8{0x43, 0x4e, 0x59}, []uint8{0x43, 0x4f, 0x50}, []uint8{0x43, 0x4f, 0x55}, []uint8{0x43, 0x52, 0x43}, []uint8{0x43, 0x53, 0x44}, []uint8{0x43, 0x53, 0x4b}, []uint8{0x43, 0x55, 0x43}, []uint8{0x43, 0x55, 0x50}, []uint8{0x43, 0x56, 0x45}, []uint8{0x43, 0x59, 0x50}, []uint8{0x43, 0x5a, 0x4b}, []uint8{0x44, 0x44, 0x4d}, []uint8{0x44, 0x45, 0x4d}, []uint8{0x44, 0x4a, 0x46}, []uint8{0x44, 0x4b, 0x4b}, []uint8{0x44, 0x4f, 0x50}, []uint8{0x44, 0x5a, 0x44}, []uint8{0x45, 0x43, 0x53}, []uint8{0x45, 0x43, 0x56}, []uint8{0x45, 0x45, 0x4b}, []uint8{0x45, 0x47, 0x50}, []uint8{0x45, 0x52, 0x4e}, []uint8{0x45, 0x53, 0x41}, []uint8{0x45, 0x53, 0x42}, []uint8{0x45, 0x53, 0x50}, []uint8{0x45, 0x54, 0x42}, []uint8{0x45, 0x55, 0x52}, []uint8{0x46, 0x49, 0x4d}, []uint8{0x46, 0x4a, 0x44}, []uint8{0x46, 0x4b, 0x50}, []uint8{0x46, 0x52, 0x46}, []uint8{0x47, 0x42, 0x50}, []uint8{0x47, 0x45, 0x4b}, []uint8{0x47, 0x45, 0x4c}, []uint8{0x47, 0x48, 0x43}, []uint8{0x47, 0x48, 0x53}, []uint8{0x47, 0x49, 0x50}, []uint8{0x47, 0x4d, 0x44}, []uint8{0x47, 0x4e, 0x46}, []uint8{0x47, 0x4e, 0x53}, []uint8{0x47, 0x51, 0x45}, []uint8{0x47, 0x52, 0x44}, []uint8{0x47, 0x54, 0x51}, []uint8{0x47, 0x57, 0x45}, []uint8{0x47, 0x57, 0x50}, []uint8{0x47, 0x59, 0x44}, []uint8{0x48, 0x4b, 0x44}, []uint8{0x48, 0x4e, 0x4c}, []uint8{0x48, 0x52, 0x44}, []uint8{0x48, 0x52, 0x4b}, []uint8{0x48, 0x54, 0x47}, []uint8{0x48, 0x55, 0x46}, []uint8{0x49, 0x44, 0x52}, []uint8{0x49, 0x45, 0x50}, []uint8{0x49, 0x4c, 0x50}, []uint8{0x49, 0x4c, 0x52}, []uint8{0x49, 0x4c, 0x53}, []uint8{0x49, 0x4e, 0x52}, []uint8{0x49, 0x51, 0x44}, []uint8{0x49, 0x52, 0x52}, []uint8{0x49, 0x53, 0x4a}, []uint8{0x49, 0x53, 0x4b}, []uint8{0x49, 0x54, 0x4c}, []uint8{0x4a, 0x4d, 0x44}, []uint8{0x4a, 0x4f, 0x44}, []uint8{0x4a, 0x50, 0x59}, []uint8{0x4b, 0x45, 0x53}, []uint8{0x4b, 0x47, 0x53}, []uint8{0x4b, 0x48, 0x52}, []uint8{0x4b, 0x4d, 0x46}, []uint8{0x4b, 0x50, 0x57}, []uint8{0x4b, 0x52, 0x48}, []uint8{0x4b, 0x52, 0x4f}, []uint8{0x4b, 0x52, 0x57}, []uint8{0x4b, 0x57, 0x44}, []uint8{0x4b, 0x59, 0x44}, []uint8{0x4b, 0x5a, 0x54}, []uint8{0x4c, 0x41, 0x4b}, []uint8{0x4c, 0x42, 0x50}, []uint8{0x4c, 0x4b, 0x52}, []uint8{0x4c, 0x52, 0x44}, []uint8{0x4c, 0x53, 0x4c}, []uint8{0x4c, 0x54, 0x4c}, []uint8{0x4c, 0x54, 0x54}, []uint8{0x4c, 0x55, 0x43}, []uint8{0x4c, 0x55, 0x46}, []uint8{0x4c, 0x55, 0x4c}, []uint8{0x4c, 0x56, 0x4c}, []uint8{0x4c, 0x56, 0x52}, []uint8{0x4c, 0x59, 0x44}, []uint8{0x4d, 0x41, 0x44}, []uint8{0x4d, 0x41, 0x46}, []uint8{0x4d, 0x43, 0x46}, []uint8{0x4d, 0x44, 0x43}, []uint8{0x4d, 0x44, 0x4c}, []uint8{0x4d, 0x47, 0x41}, []uint8{0x4d, 0x47, 0x46}, []uint8{0x4d, 0x4b, 0x44}, []uint8{0x4d, 0x4b, 0x4e}, []uint8{0x4d, 0x4c, 0x46}, []uint8{0x4d, 0x4d, 0x4b}, []uint8{0x4d, 0x4e, 0x54}, []uint8{0x4d, 0x4f, 0x50}, []uint8{0x4d, 0x52, 0x4f}, []uint8{0x4d, 0x54, 0x4c}, []uint8{0x4d, 0x54, 0x50}, []uint8{0x4d, 0x55, 0x52}, []uint8{0x4d, 0x56, 0x50}, []uint8{0x4d, 0x56, 0x52}, []uint8{0x4d, 0x57, 0x4b}, []uint8{0x4d, 0x58, 0x4e}, []uint8{0x4d, 0x58, 0x50}, []uint8{0x4d, 0x58, 0x56}, []uint8{0x4d, 0x59, 0x52}, []uint8{0x4d, 0x5a, 0x45}, []uint8{0x4d, 0x5a, 0x4d}, []uint8{0x4d, 0x5a, 0x4e}, []uint8{0x4e, 0x41, 0x44}, []uint8{0x4e, 0x47, 0x4e}, []uint8{0x4e, 0x49, 0x43}, []uint8{0x4e, 0x49, 0x4f}, []uint8{0x4e, 0x4c, 0x47}, []uint8{0x4e, 0x4f, 0x4b}, []uint8{0x4e, 0x50, 0x52}, []uint8{0x4e, 0x5a, 0x44}, []uint8{0x4f, 0x4d, 0x52}, []uint8{0x50, 0x41, 0x42}, []uint8{0x50, 0x45, 0x49}, []uint8{0x50, 0x45, 0x4e}, []uint8{0x50, 0x45, 0x53}, []uint8{0x50, 0x47, 0x4b}, []uint8{0x50, 0x48, 0x50}, []uint8{0x50, 0x4b, 0x52}, []uint8{0x50, 0x4c, 0x4e}, []uint8{0x50, 0x4c, 0x5a}, []uint8{0x50, 0x54, 0x45}, []uint8{0x50, 0x59, 0x47}, []uint8{0x51, 0x41, 0x52}, []uint8{0x52, 0x48, 0x44}, []uint8{0x52, 0x4f, 0x4c}, []uint8{0x52, 0x4f, 0x4e}, []uint8{0x52, 0x53, 0x44}, []uint8{0x52, 0x55, 0x42}, []uint8{0x52, 0x55, 0x52}, []uint8{0x52, 0x57, 0x46}, []uint8{0x53, 0x41, 0x52}, []uint8{0x53, 0x42, 0x44}, []uint8{0x53, 0x43, 0x52}, []uint8{0x53, 0x44, 0x44}, []uint8{0x53, 0x44, 0x47}, []uint8{0x53, 0x44, 0x50}, []uint8{0x53, 0x45, 0x4b}, []uint8{0x53, 0x47, 0x44}, []uint8{0x53, 0x48, 0x50}, []uint8{0x53, 0x49, 0x54}, []uint8{0x53, 0x4b, 0x4b}, []uint8{0x53, 0x4c, 0x4c}, []uint8{0x53, 0x4f, 0x53}, []uint8{0x53, 0x52, 0x44}, []uint8{0x53, 0x52, 0x47}, []uint8{0x53, 0x53, 0x50}, []uint8{0x53, 0x54, 0x44}, []uint8{0x53, 0x55, 0x52}, []uint8{0x53, 0x56, 0x43}, []uint8{0x53, 0x59, 0x50}, []uint8{0x53, 0x5a, 0x4c}, []uint8{0x54, 0x48, 0x42}, []uint8{0x54, 0x4a, 0x52}, []uint8{0x54, 0x4a, 0x53}, []uint8{0x54, 0x4d, 0x4d}, []uint8{0x54, 0x4d, 0x54}, []uint8{0x54, 0x4e, 0x44}, []uint8{0x54, 0x4f, 0x50}, []uint8{0x54, 0x50, 0x45}, []uint8{0x54, 0x52, 0x4c}, []uint8{0x54, 0x52, 0x59}, []uint8{0x54, 0x54, 0x44}, []uint8{0x54, 0x57, 0x44}, []uint8{0x54, 0x5a, 0x53}, []uint8{0x55, 0x41, 0x48}, []uint8{0x55, 0x41, 0x4b}, []uint8{0x55, 0x47, 0x53}, []uint8{0x55, 0x47, 0x58}, []uint8{0x55, 0x53, 0x44}, []uint8{0x55, 0x53, 0x4e}, []uint8{0x55, 0x53, 0x53}, []uint8{0x55, 0x59, 0x49}, []uint8{0x55, 0x59, 0x50}, []uint8{0x55, 0x59, 0x55}, []uint8{0x55, 0x5a, 0x53}, []uint8{0x56, 0x45, 0x42}, []uint8{0x56, 0x45, 0x46}, []uint8{0x56, 0x4e, 0x44}, []uint8{0x56, 0x4e, 0x4e}, []uint8{0x56, 0x55, 0x56}, []uint8{0x57, 0x53, 0x54}, []uint8{0x58, 0x41, 0x46}, []uint8{0x58, 0x41, 0x47}, []uint8{0x58, 0x41, 0x55}, []uint8{0x58, 0x42, 0x41}, []uint8{0x58, 0x42, 0x42}, []uint8{0x58, 0x42, 0x43}, []uint8{0x58, 0x42, 0x44}, []uint8{0x58, 0x43, 0x44}, []uint8{0x58, 0x44, 0x52}, []uint8{0x58, 0x45, 0x55}, []uint8{0x58, 0x46, 0x4f}, []uint8{0x58, 0x46, 0x55}, []uint8{0x58, 0x4f, 0x46}, []uint8{0x58, 0x50, 0x44}, []uint8{0x58, 0x50, 0x46}, []uint8{0x58, 0x50, 0x54}, []uint8{0x58, 0x52, 0x45}, []uint8{0x58, 0x53, 0x55}, []uint8{0x58, 0x54, 0x53}, []uint8{0x58, 0x55, 0x41}, []uint8{0x58, 0x58, 0x58}, []uint8{0x59, 0x44, 0x44}, []uint8{0x59, 0x45, 0x52}, []uint8{0x59, 0x55, 0x44}, []uint8{0x59, 0x55, 0x4d}, []uint8{0x59, 0x55, 0x4e}, []uint8{0x59, 0x55, 0x52}, []uint8{0x5a, 0x41, 0x4c}, []uint8{0x5a, 0x41, 0x52}, []uint8{0x5a, 0x4d, 0x4b}, []uint8{0x5a, 0x4d, 0x57}, []uint8{0x5a, 0x52, 0x4e}, []uint8{0x5a, 0x52, 0x5a}, []uint8{0x5a, 0x57, 0x44}, []uint8{0x5a, 0x57, 0x4c}, []uint8{0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},

		currencyNegativeSuffix: []byte{0xc2, 0xa0},
	}
}

// Locale returns the current translators string locale
func (ewo *ewo_CM) Locale() string {
	return ewo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ewo_CM'
func (ewo *ewo_CM) PluralsCardinal() []locales.PluralRule {
	return ewo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ewo_CM'
func (ewo *ewo_CM) PluralsOrdinal() []locales.PluralRule {
	return ewo.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ewo_CM'
func (ewo *ewo_CM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ewo_CM'
func (ewo *ewo_CM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ewo_CM'
func (ewo *ewo_CM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ewo_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ewo *ewo_CM) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ewo.decimal) + len(ewo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ewo.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ewo.group) - 1; j >= 0; j-- {
					b = append(b, ewo.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ewo.minus) - 1; j >= 0; j-- {
			b = append(b, ewo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ewo_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ewo *ewo_CM) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ewo.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ewo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ewo.minus) - 1; j >= 0; j-- {
			b = append(b, ewo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ewo.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ewo_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ewo *ewo_CM) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ewo.currencies[currency]
	l := len(s) + len(ewo.decimal) + len(ewo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ewo.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ewo.group) - 1; j >= 0; j-- {
					b = append(b, ewo.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ewo.minus) - 1; j >= 0; j-- {
			b = append(b, ewo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ewo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ewo.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}
