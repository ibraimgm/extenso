package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// the 'success' return type
type extensoJson struct {
	Extenso string `json:"extenso,omitempty"`
}

// the 'error' return type
type extensoErr struct {
	Erro string `json:"erro,omitempty"`
}

func (e extensoErr) Error() string {
	return e.Erro
}

// we always return the same message, so it makes sense to create
// a global value with it (remember, structs can't be const!)
var errPathInvalido = extensoErr{Erro: "O path deve ser um numero inteiro entre -99999 e 99999."}

// the handler, responsible to take the path, parse it,
// and pass to the translation routine
func extensoHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ignore the first "/" on path
	path := r.URL.Path[1:]
	i64, err := strconv.ParseInt(path, 10, 0)
	if err != nil {
		log.Println(err.Error()) // log the original error, for debugging purposes
		http.Error(w, marshalJson(errPathInvalido), http.StatusBadRequest)
		return
	}

	res, err := extenso(int(i64))
	if err != nil {
		log.Println(err.Error()) // log the original error, for debugging purposes
		http.Error(w, marshalJson(errPathInvalido), http.StatusBadRequest)
		return
	}

	// if we got this far, we should have a valid result
	// anyway, log it on console just for fun
	jsonStr := marshalJson(res)
	log.Printf("Request: %s, Response: %s", r.URL.Path, jsonStr)
	fmt.Fprint(w, jsonStr)
}

// base transaltion names for portuguese
var translation = map[int]string{
	0:   "zero",
	1:   "um",
	2:   "dois",
	3:   "três",
	4:   "quatro",
	5:   "cinco",
	6:   "seis",
	7:   "sete",
	8:   "oito",
	9:   "nove",
	10:  "dez",
	11:  "onze",
	12:  "doze",
	13:  "treze",
	14:  "quatorze",
	15:  "quinze",
	16:  "dezesseis",
	17:  "dezessete",
	18:  "dezoito",
	19:  "dezenove",
	20:  "vinte",
	30:  "trinta",
	40:  "quarenta",
	50:  "cinquenta",
	60:  "sessenta",
	70:  "setenta",
	80:  "oitenta",
	90:  "noventa",
	100: "cento",
	200: "duzentos",
	300: "trezentos",
	400: "quatrocentos",
	500: "quinhentos",
	600: "seiscentos",
	700: "setecentos",
	800: "oitocentos",
	900: "novecentos",
}

// extenso is the main translation routine
func extenso(n int) (extensoJson, error) {
	if n < -99999 || n > 99999 {
		return extensoJson{}, errPathInvalido
	}

	// don't worry about negative numbers
	var menos string
	if n < 0 {
		n = -n
		menos = "menos "
	}

	// check for a exact match.
	// if found, we don't need to waste time separating
	// the digits. The only odd case here is '100'
	if n == 100 {
		return extensoJson{Extenso: menos + "cem"}, nil
	} else if match, ok := translation[n]; ok {
		return extensoJson{Extenso: menos + match}, nil
	}

	milhar := n / 1000
	centena := (n - milhar*1000) / 100
	dezena := (n - (milhar*1000 + centena*100))

	// most of the time, we will have 3 pieces of data, so
	// prealloc this to make append faster
	result := make([]string, 0, 3)

	// thousands place has a corner case
	// when the digit is '1'
	if milhar == 1 {
		result = append(result, "mil")
	} else if milhar > 0 {
		result = append(result, doisDigitos(milhar)+" mil")
	}

	// hundreds place has a corner case (cem vs cento)
	if centena == 1 && dezena == 0 {
		result = append(result, "cem")
	} else {
		if centena > 0 {
			result = append(result, translation[centena*100])
		}

		if dezena > 0 {
			result = append(result, doisDigitos(dezena))
		}
	}

	return extensoJson{Extenso: menos + strings.Join(result, " e ")}, nil
}

// doisDigitos translates a two-digit number
func doisDigitos(n int) string {
	s, ok := translation[n]
	if !ok {
		dezena := n / 10
		unidade := n % 10

		ds := translation[dezena*10]
		us := translation[unidade]

		s = ds + " e " + us
	}

	return s
}

// aux. function to return a json string from a value
// if an error happens, it is logged and the function
// returns an empty string
func marshalJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return string(b)
}
