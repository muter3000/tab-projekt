package schemas

import "regexp"

func (p *Pracownik) validate() (bool, error) {
	peselCheck := regexp.MustCompile(`[0-9]{11}`)
	isPeselValid := len(peselCheck.FindAllString(p.Pesel, -1)) == 1

	isNameValid, err := regexp.MatchString(`[A-Za-z]+`, p.Imie)
	if err != nil {
		return false, err
	}
	isSurnameValid, err := regexp.MatchString(`[A-Za-z]+`, p.Nazwisko)
	if err != nil {
		return false, err
	}
	isHasloValid := len(p.Haslo) > 6
	isLoginValid := len(p.Login) > 0

	return isHasloValid && isLoginValid && isSurnameValid && isNameValid && isPeselValid, nil
}
