package sv_SE

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sv_SE struct {
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

// New returns a new instance of translator for the 'sv_SE' locale
func New() locales.Translator {
	return &sv_SE{
		locale:          "sv_SE",
		pluralsCardinal: []locales.PluralRule{2, 6},
		pluralsOrdinal:  []locales.PluralRule{2, 6},
		decimal:         []byte{},
		group:           []byte{},
		minus:           []byte{0xe2, 0x80, 0x8f, 0xe2, 0x88, 0x92},
		percent:         []byte{},
		perMille:        []byte{},
		currencies:      [][]uint8{[]uint8{0x41, 0x44, 0x50}, []uint8{0x41, 0x45, 0x44}, []uint8{0x41, 0x46, 0x41}, []uint8{0x41, 0x46, 0x4e}, []uint8{0x41, 0x4c, 0x4b}, []uint8{0x41, 0x4c, 0x4c}, []uint8{0x41, 0x4d, 0x44}, []uint8{0x41, 0x4e, 0x47}, []uint8{0x41, 0x4f, 0x41}, []uint8{0x41, 0x4f, 0x4b}, []uint8{0x41, 0x4f, 0x4e}, []uint8{0x41, 0x4f, 0x52}, []uint8{0x41, 0x52, 0x41}, []uint8{0x41, 0x52, 0x4c}, []uint8{0x41, 0x52, 0x4d}, []uint8{0x41, 0x52, 0x50}, []uint8{0x41, 0x52, 0x53}, []uint8{0x41, 0x54, 0x53}, []uint8{0x41, 0x55, 0x44}, []uint8{0x41, 0x57, 0x47}, []uint8{0x41, 0x5a, 0x4d}, []uint8{0x41, 0x5a, 0x4e}, []uint8{0x42, 0x41, 0x44}, []uint8{0x42, 0x41, 0x4d}, []uint8{0x42, 0x41, 0x4e}, []uint8{0x42, 0x42, 0x44}, []uint8{0x42, 0x44, 0x54}, []uint8{0x42, 0x45, 0x43}, []uint8{0x42, 0x45, 0x46}, []uint8{0x42, 0x45, 0x4c}, []uint8{0x42, 0x47, 0x4c}, []uint8{0x42, 0x47, 0x4d}, []uint8{0x42, 0x47, 0x4e}, []uint8{0x42, 0x47, 0x4f}, []uint8{0x42, 0x48, 0x44}, []uint8{0x42, 0x49, 0x46}, []uint8{0x42, 0x4d, 0x44}, []uint8{0x42, 0x4e, 0x44}, []uint8{0x42, 0x4f, 0x42}, []uint8{0x42, 0x4f, 0x4c}, []uint8{0x42, 0x4f, 0x50}, []uint8{0x42, 0x4f, 0x56}, []uint8{0x42, 0x52, 0x42}, []uint8{0x42, 0x52, 0x43}, []uint8{0x42, 0x52, 0x45}, []uint8{0x42, 0x52, 0x4c}, []uint8{0x42, 0x52, 0x4e}, []uint8{0x42, 0x52, 0x52}, []uint8{0x42, 0x52, 0x5a}, []uint8{0x42, 0x53, 0x44}, []uint8{0x42, 0x54, 0x4e}, []uint8{0x42, 0x55, 0x4b}, []uint8{0x42, 0x57, 0x50}, []uint8{0x42, 0x59, 0x42}, []uint8{0x42, 0x59, 0x52}, []uint8{0x42, 0x5a, 0x44}, []uint8{0x43, 0x41, 0x44}, []uint8{0x43, 0x44, 0x46}, []uint8{0x43, 0x48, 0x45}, []uint8{0x43, 0x48, 0x46}, []uint8{0x43, 0x48, 0x57}, []uint8{0x43, 0x4c, 0x45}, []uint8{0x43, 0x4c, 0x46}, []uint8{0x43, 0x4c, 0x50}, []uint8{0x43, 0x4e, 0x58}, []uint8{0x43, 0x4e, 0x59}, []uint8{0x43, 0x4f, 0x50}, []uint8{0x43, 0x4f, 0x55}, []uint8{0x43, 0x52, 0x43}, []uint8{0x43, 0x53, 0x44}, []uint8{0x43, 0x53, 0x4b}, []uint8{0x43, 0x55, 0x43}, []uint8{0x43, 0x55, 0x50}, []uint8{0x43, 0x56, 0x45}, []uint8{0x43, 0x59, 0x50}, []uint8{0x43, 0x5a, 0x4b}, []uint8{0x44, 0x44, 0x4d}, []uint8{0x44, 0x45, 0x4d}, []uint8{0x44, 0x4a, 0x46}, []uint8{0x44, 0x4b, 0x4b}, []uint8{0x44, 0x4f, 0x50}, []uint8{0x44, 0x5a, 0x44}, []uint8{0x45, 0x43, 0x53}, []uint8{0x45, 0x43, 0x56}, []uint8{0x45, 0x45, 0x4b}, []uint8{0x45, 0x47, 0x50}, []uint8{0x45, 0x52, 0x4e}, []uint8{0x45, 0x53, 0x41}, []uint8{0x45, 0x53, 0x42}, []uint8{0x45, 0x53, 0x50}, []uint8{0x45, 0x54, 0x42}, []uint8{0x45, 0x55, 0x52}, []uint8{0x46, 0x49, 0x4d}, []uint8{0x46, 0x4a, 0x44}, []uint8{0x46, 0x4b, 0x50}, []uint8{0x46, 0x52, 0x46}, []uint8{0x47, 0x42, 0x50}, []uint8{0x47, 0x45, 0x4b}, []uint8{0x47, 0x45, 0x4c}, []uint8{0x47, 0x48, 0x43}, []uint8{0x47, 0x48, 0x53}, []uint8{0x47, 0x49, 0x50}, []uint8{0x47, 0x4d, 0x44}, []uint8{0x47, 0x4e, 0x46}, []uint8{0x47, 0x4e, 0x53}, []uint8{0x47, 0x51, 0x45}, []uint8{0x47, 0x52, 0x44}, []uint8{0x47, 0x54, 0x51}, []uint8{0x47, 0x57, 0x45}, []uint8{0x47, 0x57, 0x50}, []uint8{0x47, 0x59, 0x44}, []uint8{0x48, 0x4b, 0x44}, []uint8{0x48, 0x4e, 0x4c}, []uint8{0x48, 0x52, 0x44}, []uint8{0x48, 0x52, 0x4b}, []uint8{0x48, 0x54, 0x47}, []uint8{0x48, 0x55, 0x46}, []uint8{0x49, 0x44, 0x52}, []uint8{0x49, 0x45, 0x50}, []uint8{0x49, 0x4c, 0x50}, []uint8{0x49, 0x4c, 0x52}, []uint8{0x49, 0x4c, 0x53}, []uint8{0x49, 0x4e, 0x52}, []uint8{0x49, 0x51, 0x44}, []uint8{0x49, 0x52, 0x52}, []uint8{0x49, 0x53, 0x4a}, []uint8{0x49, 0x53, 0x4b}, []uint8{0x49, 0x54, 0x4c}, []uint8{0x4a, 0x4d, 0x44}, []uint8{0x4a, 0x4f, 0x44}, []uint8{0x4a, 0x50, 0x59}, []uint8{0x4b, 0x45, 0x53}, []uint8{0x4b, 0x47, 0x53}, []uint8{0x4b, 0x48, 0x52}, []uint8{0x4b, 0x4d, 0x46}, []uint8{0x4b, 0x50, 0x57}, []uint8{0x4b, 0x52, 0x48}, []uint8{0x4b, 0x52, 0x4f}, []uint8{0x4b, 0x52, 0x57}, []uint8{0x4b, 0x57, 0x44}, []uint8{0x4b, 0x59, 0x44}, []uint8{0x4b, 0x5a, 0x54}, []uint8{0x4c, 0x41, 0x4b}, []uint8{0x4c, 0x42, 0x50}, []uint8{0x4c, 0x4b, 0x52}, []uint8{0x4c, 0x52, 0x44}, []uint8{0x4c, 0x53, 0x4c}, []uint8{0x4c, 0x54, 0x4c}, []uint8{0x4c, 0x54, 0x54}, []uint8{0x4c, 0x55, 0x43}, []uint8{0x4c, 0x55, 0x46}, []uint8{0x4c, 0x55, 0x4c}, []uint8{0x4c, 0x56, 0x4c}, []uint8{0x4c, 0x56, 0x52}, []uint8{0x4c, 0x59, 0x44}, []uint8{0x4d, 0x41, 0x44}, []uint8{0x4d, 0x41, 0x46}, []uint8{0x4d, 0x43, 0x46}, []uint8{0x4d, 0x44, 0x43}, []uint8{0x4d, 0x44, 0x4c}, []uint8{0x4d, 0x47, 0x41}, []uint8{0x4d, 0x47, 0x46}, []uint8{0x4d, 0x4b, 0x44}, []uint8{0x4d, 0x4b, 0x4e}, []uint8{0x4d, 0x4c, 0x46}, []uint8{0x4d, 0x4d, 0x4b}, []uint8{0x4d, 0x4e, 0x54}, []uint8{0x4d, 0x4f, 0x50}, []uint8{0x4d, 0x52, 0x4f}, []uint8{0x4d, 0x54, 0x4c}, []uint8{0x4d, 0x54, 0x50}, []uint8{0x4d, 0x55, 0x52}, []uint8{0x4d, 0x56, 0x50}, []uint8{0x4d, 0x56, 0x52}, []uint8{0x4d, 0x57, 0x4b}, []uint8{0x4d, 0x58, 0x4e}, []uint8{0x4d, 0x58, 0x50}, []uint8{0x4d, 0x58, 0x56}, []uint8{0x4d, 0x59, 0x52}, []uint8{0x4d, 0x5a, 0x45}, []uint8{0x4d, 0x5a, 0x4d}, []uint8{0x4d, 0x5a, 0x4e}, []uint8{0x4e, 0x41, 0x44}, []uint8{0x4e, 0x47, 0x4e}, []uint8{0x4e, 0x49, 0x43}, []uint8{0x4e, 0x49, 0x4f}, []uint8{0x4e, 0x4c, 0x47}, []uint8{0x4e, 0x4f, 0x4b}, []uint8{0x4e, 0x50, 0x52}, []uint8{0x4e, 0x5a, 0x44}, []uint8{0x4f, 0x4d, 0x52}, []uint8{0x50, 0x41, 0x42}, []uint8{0x50, 0x45, 0x49}, []uint8{0x50, 0x45, 0x4e}, []uint8{0x50, 0x45, 0x53}, []uint8{0x50, 0x47, 0x4b}, []uint8{0x50, 0x48, 0x50}, []uint8{0x50, 0x4b, 0x52}, []uint8{0x50, 0x4c, 0x4e}, []uint8{0x50, 0x4c, 0x5a}, []uint8{0x50, 0x54, 0x45}, []uint8{0x50, 0x59, 0x47}, []uint8{0x51, 0x41, 0x52}, []uint8{0x52, 0x48, 0x44}, []uint8{0x52, 0x4f, 0x4c}, []uint8{0x52, 0x4f, 0x4e}, []uint8{0x52, 0x53, 0x44}, []uint8{0x52, 0x55, 0x42}, []uint8{0x52, 0x55, 0x52}, []uint8{0x52, 0x57, 0x46}, []uint8{0x53, 0x41, 0x52}, []uint8{0x53, 0x42, 0x44}, []uint8{0x53, 0x43, 0x52}, []uint8{0x53, 0x44, 0x44}, []uint8{0x53, 0x44, 0x47}, []uint8{0x53, 0x44, 0x50}, []uint8{0x53, 0x45, 0x4b}, []uint8{0x53, 0x47, 0x44}, []uint8{0x53, 0x48, 0x50}, []uint8{0x53, 0x49, 0x54}, []uint8{0x53, 0x4b, 0x4b}, []uint8{0x53, 0x4c, 0x4c}, []uint8{0x53, 0x4f, 0x53}, []uint8{0x53, 0x52, 0x44}, []uint8{0x53, 0x52, 0x47}, []uint8{0x53, 0x53, 0x50}, []uint8{0x53, 0x54, 0x44}, []uint8{0x53, 0x55, 0x52}, []uint8{0x53, 0x56, 0x43}, []uint8{0x53, 0x59, 0x50}, []uint8{0x53, 0x5a, 0x4c}, []uint8{0x54, 0x48, 0x42}, []uint8{0x54, 0x4a, 0x52}, []uint8{0x54, 0x4a, 0x53}, []uint8{0x54, 0x4d, 0x4d}, []uint8{0x54, 0x4d, 0x54}, []uint8{0x54, 0x4e, 0x44}, []uint8{0x54, 0x4f, 0x50}, []uint8{0x54, 0x50, 0x45}, []uint8{0x54, 0x52, 0x4c}, []uint8{0x54, 0x52, 0x59}, []uint8{0x54, 0x54, 0x44}, []uint8{0x54, 0x57, 0x44}, []uint8{0x54, 0x5a, 0x53}, []uint8{0x55, 0x41, 0x48}, []uint8{0x55, 0x41, 0x4b}, []uint8{0x55, 0x47, 0x53}, []uint8{0x55, 0x47, 0x58}, []uint8{0x55, 0x53, 0x44}, []uint8{0x55, 0x53, 0x4e}, []uint8{0x55, 0x53, 0x53}, []uint8{0x55, 0x59, 0x49}, []uint8{0x55, 0x59, 0x50}, []uint8{0x55, 0x59, 0x55}, []uint8{0x55, 0x5a, 0x53}, []uint8{0x56, 0x45, 0x42}, []uint8{0x56, 0x45, 0x46}, []uint8{0x56, 0x4e, 0x44}, []uint8{0x56, 0x4e, 0x4e}, []uint8{0x56, 0x55, 0x56}, []uint8{0x57, 0x53, 0x54}, []uint8{0x58, 0x41, 0x46}, []uint8{0x58, 0x41, 0x47}, []uint8{0x58, 0x41, 0x55}, []uint8{0x58, 0x42, 0x41}, []uint8{0x58, 0x42, 0x42}, []uint8{0x58, 0x42, 0x43}, []uint8{0x58, 0x42, 0x44}, []uint8{0x58, 0x43, 0x44}, []uint8{0x58, 0x44, 0x52}, []uint8{0x58, 0x45, 0x55}, []uint8{0x58, 0x46, 0x4f}, []uint8{0x58, 0x46, 0x55}, []uint8{0x58, 0x4f, 0x46}, []uint8{0x58, 0x50, 0x44}, []uint8{0x58, 0x50, 0x46}, []uint8{0x58, 0x50, 0x54}, []uint8{0x58, 0x52, 0x45}, []uint8{0x58, 0x53, 0x55}, []uint8{0x58, 0x54, 0x53}, []uint8{0x58, 0x55, 0x41}, []uint8{0x58, 0x58, 0x58}, []uint8{0x59, 0x44, 0x44}, []uint8{0x59, 0x45, 0x52}, []uint8{0x59, 0x55, 0x44}, []uint8{0x59, 0x55, 0x4d}, []uint8{0x59, 0x55, 0x4e}, []uint8{0x59, 0x55, 0x52}, []uint8{0x5a, 0x41, 0x4c}, []uint8{0x5a, 0x41, 0x52}, []uint8{0x5a, 0x4d, 0x4b}, []uint8{0x5a, 0x4d, 0x57}, []uint8{0x5a, 0x52, 0x4e}, []uint8{0x5a, 0x52, 0x5a}, []uint8{0x5a, 0x57, 0x44}, []uint8{0x5a, 0x57, 0x4c}, []uint8{0x5a, 0x57, 0x52}},
		percentSuffix:   []byte{0xc2, 0xa0},

		currencyPositiveSuffix: []byte{0xc2, 0xa0},

		currencyNegativeSuffix: []byte{0xc2, 0xa0},
	}
}

// Locale returns the current translators string locale
func (sv *sv_SE) Locale() string {
	return sv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sv_SE'
func (sv *sv_SE) PluralsCardinal() []locales.PluralRule {
	return sv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sv_SE'
func (sv *sv_SE) PluralsOrdinal() []locales.PluralRule {
	return sv.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sv_SE'
func (sv *sv_SE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sv_SE'
func (sv *sv_SE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if (nMod10 == 1 || nMod10 == 2) && (nMod100 != 11 && nMod100 != 12) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sv_SE'
func (sv *sv_SE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sv_SE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sv *sv_SE) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sv.decimal) + len(sv.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(sv.decimal) - 1; j >= 0; j-- {
				b = append(b, sv.decimal[j])
			}

			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sv.group) - 1; j >= 0; j-- {
					b = append(b, sv.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sv_SE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sv *sv_SE) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sv.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(sv.decimal) - 1; j >= 0; j-- {
				b = append(b, sv.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sv.percentSuffix...)

	b = append(b, sv.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sv_SE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sv *sv_SE) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sv.currencies[currency]
	l := len(s) + len(sv.decimal) + len(sv.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(sv.decimal) - 1; j >= 0; j-- {
				b = append(b, sv.decimal[j])
			}

			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sv.group) - 1; j >= 0; j-- {
					b = append(b, sv.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sv.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}
