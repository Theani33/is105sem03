package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; KSN")

func Krypter(input []rune, key []rune, chiffer int) []rune {
    output := make([]rune, len(input))
    for i, r := range input {
        index := sokIAlfabetet(r, key)
        if index < 0 {
            // character not found in key
            output[i] = r
        } else {
            encryptedIndex := (index + chiffer) % len(key)
            output[i] = key[encryptedIndex]
        }
    }
    return output
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
    for i := 0; i < len(alfabet); i++ {
        if symbol == alfabet[i] {
            return i
            break
        }
    }
    return -1
}