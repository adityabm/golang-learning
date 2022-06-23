package entity

type Siswa struct {
	ID    int
	Name  string
	Class string
}

func (s Siswa) BayarSpp() string {
	var status string

	if s.Name == "Anton" {
		status = "Bayar SPP"
	} else if s.Name == "Adit" {
		status = "Administrasi belum lengkap"
	}

	return status
}