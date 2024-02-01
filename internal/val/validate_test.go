package val

import (
	"testing"
)

func Test_ValidatePno(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		tests := [...]string{
			"4112098357",
			"3311305498",
			"7112050112",
			"7412188927",
			"3812145997",
			"6701203199",
			"8004113877",
			"9806275906",
			"2802107041",
			"0007060833",
			"991227-0171",
			"720813-5637",
			"370123-4878",
			"350717-7297",
			"871010-5225",
			"651214-3329",
			"041008-1400",
			"750930-7117",
			"821010-9701",
			"200303178578",
			"193902179021",
			"198801266217",
			"197511269453",
			"200505206169",
			"199304288872",
			"195012150594",
			"200003170933",
			"196908054858",
			"199702051559",
			"200010201317",
			"196106122234",
			"19410308-0778",
			"19341028-8843",
			"19840928-9744",
			"19340815-3322",
			"19921220-7667",
			"20000104-6416",
			"19721218-8754",
			"20020226-3695",
		}

		for _, pno := range tests {
			valid, checksum, err := ValidatePno(pno)
			if err != nil {
				t.Error(err)
			}

			if !valid {
				t.Errorf("'%s' should be a valid PNO. checksum is '%s'", pno, checksum)
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		tests := [...]string{
			"198003257629",
			"19800325762",
			"1980032576",
			"198003257",
			"19800325",
			"1980032",
			"198003",
			"19800",
			"1980",
			"198",
			"19",
			"1",
		}

		for _, pno := range tests {
			valid, checksum, _ := ValidatePno(pno)

			if valid {
				t.Errorf("'%s' should be a invalid PNO. checksum is '%s'", pno, checksum)
			}
		}
	})
}
