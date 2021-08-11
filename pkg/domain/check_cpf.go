package domain

import "strings"

// CheckCPF checks if the cpf exists and returns nil if not, it returns an error
func CheckCPF(cpf string) (error, string) {
	if len(cpf) != 11 && len(cpf) != 14 {
		return ErrLogin, cpf
	}
	if len(cpf) == 14 {
		if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
			return nil, CpfReplace(cpf)
		} else {
			return ErrLogin, cpf
		}
	}
	return nil, cpf
}

// CpfReplace returns the CPF received in a single format
func CpfReplace(cpf string) string {
	cpfReplace := strings.Replace(cpf, ".", "", 2)
	cpfReplace = strings.Replace(cpfReplace, "-", "", 1)
	return cpfReplace
}
