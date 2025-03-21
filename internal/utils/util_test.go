package utils

import "testing"

func Test_ValidateEmail(t *testing.T) {
	tests := []struct {
		have string
		ok   bool
	}{
		{have: "user@example.com", ok: true},
		{have: "user.name@example.com", ok: true},
		{have: "user+name@example.com", ok: true},
		{have: "user-name@example.com", ok: true},
		{have: "user_name@example.com", ok: true},
		{have: "user123@example.com", ok: true},
		{have: "user@sub.example.com", ok: true},
		{have: "user@example.co.uk", ok: true},
		{have: "user@example.io", ok: true},
		{have: "user@123.example.com", ok: true},
		{have: "user@.com", ok: false},
		{have: "user@com", ok: false},
		{have: "user@example..com", ok: false},
		{have: "user@example.c", ok: false},
		{have: "user@example.#com", ok: false},
		{have: "user@example_com", ok: false},
		{have: "user@example,com", ok: false},
		{have: "user@example com", ok: false},
		{have: "user@.example.com", ok: false},
		{have: "user@example..com", ok: false},
		{have: "user@example.c.o.m", ok: false},
		{have: "user@example.c@om", ok: false},
		{have: "user@example.com ", ok: false},
		{have: " user@example.com", ok: false},
		{have: "user@example .com", ok: false},
		{have: "user@.com", ok: false},
		{have: "user@example", ok: false},
		{have: "user@example.", ok: false},
		{have: "user@example..com", ok: false},
		{have: "user@example.c", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			if valid := ValidateEmail(tt.have); valid != tt.ok {
				t.Errorf("%s != %t", tt.have, tt.ok)
			}
		})
	}

}

func Test_ValidatePhone(t *testing.T) {
	tests := []struct {
		have string
		ok   bool
	}{
		{have: "+79161234567", ok: true},
		{have: "89161234567", ok: true},
		{have: "+7 916 123 45 67", ok: true},
		{have: "8 (916) 123-45-67", ok: true},
		{have: "+7-916-123-45-67", ok: true},
		{have: "9161234567", ok: true},
		{have: "4951234567", ok: true},
		{have: "8-495-123-45-67", ok: true},
		{have: "+7 (495) 123-45-67", ok: true},
		{have: "8 916 123 45 67", ok: true},
		{have: "+7(916)123-45-67", ok: true},
		{have: "8(916)1234567", ok: true},
		{have: "+7 916 1234567", ok: true},
		{have: "8 916 123-45-67", ok: true},
		{have: "+7 495 1234567", ok: true},

		{have: "+7916123456", ok: false},
		{have: "891612345", ok: false},
		{have: "+7 916 123 45", ok: false},
		{have: "8 (916) 123-45", ok: false},
		{have: "+7 916 123 45 6789", ok: false},
		{have: "8 916 123 45 6789", ok: false},
		{have: "+7 916 123 45 67a", ok: false},
		{have: "8 916 123 45 67!", ok: false},
		{have: "+7 916 123 45 67 ", ok: false},
		{have: " +7 916 123 45 67", ok: false},
		{have: "+7 916 123  45 67", ok: false},
		{have: "+7 916 123--45-67", ok: false},
		{have: "+7 (916) 123-45-67)", ok: false},
		{have: "+7 916) 123-45-67", ok: false},
		{have: "+7 (916 123-45-67", ok: false},
		{have: "+7 916 (123)-45-67", ok: false},
		{have: "+7 916 123-45-67-", ok: false},
		{have: "+7-916-123-45-67-", ok: false},
		{have: "+7 916 123 45 67-", ok: false},
		{have: "+7 916 123 45 67 ", ok: false},
		{have: " +7 916 123 45 67", ok: false},
		{have: "+7 916 123  45 67", ok: false},
		{have: "+7 916 123--45-67", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			if valid := ValidatePhone(tt.have); valid != tt.ok {
				t.Errorf("%s != %t", tt.have, tt.ok)
			}
		})
	}

}
