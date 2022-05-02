package schemas

import "regexp"

func (k *Kierowca) validate() (bool, error) {
	peselCheck := regexp.MustCompile(`[0-9]{11}`)
	isPeselValid := len(peselCheck.FindAllString(k.Pesel, -1)) == 1

	isNameValid, err := regexp.MatchString(`[A-Za-z]+`, k.Imie)
	if err != nil {
		return false, err
	}
	isSurnameValid, err := regexp.MatchString(`[A-Za-z]+`, k.Nazwisko)
	if err != nil {
		return false, err
	}
	isHasloValid := len(k.Haslo) > 6
	isLoginValid := len(k.Login) > 0

	return isHasloValid && isLoginValid && isSurnameValid && isNameValid && isPeselValid, nil
}
