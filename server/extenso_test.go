package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ibraimgm/extenso/server"
)

func TestExtenso(t *testing.T) {
	tests := []struct {
		url            string
		expected       string
		expectedStatus int
	}{
		// typical input
		{url: "/1", expected: `{"extenso":"um"}`},
		{url: "/0", expected: `{"extenso":"zero"}`},
		{url: "/-1", expected: `{"extenso":"menos um"}`},
		{url: "/-0", expected: `{"extenso":"zero"}`},
		{url: "/-1042", expected: `{"extenso":"menos mil e quarenta e dois"}`},
		{url: "/94587", expected: `{"extenso":"noventa e quatro mil e quinhentos e oitenta e sete"}`},
		{url: "/99999", expected: `{"extenso":"noventa e nove mil e novecentos e noventa e nove"}`},
		{url: "/100", expected: `{"extenso":"cem"}`},
		{url: "/-100", expected: `{"extenso":"menos cem"}`},
		{url: "/101", expected: `{"extenso":"cento e um"}`},
		{url: "/-101", expected: `{"extenso":"menos cento e um"}`},
		{url: "/110", expected: `{"extenso":"cento e dez"}`},
		{url: "/-110", expected: `{"extenso":"menos cento e dez"}`},
		{url: "/111", expected: `{"extenso":"cento e onze"}`},
		{url: "/-111", expected: `{"extenso":"menos cento e onze"}`},
		{url: "/1100", expected: `{"extenso":"mil e cem"}`},
		{url: "/-1100", expected: `{"extenso":"menos mil e cem"}`},
		{url: "/10100", expected: `{"extenso":"dez mil e cem"}`},
		{url: "/-10100", expected: `{"extenso":"menos dez mil e cem"}`},
		{url: "/10101", expected: `{"extenso":"dez mil e cento e um"}`},
		{url: "/-10101", expected: `{"extenso":"menos dez mil e cento e um"}`},
		{url: "/10200", expected: `{"extenso":"dez mil e duzentos"}`},
		{url: "/-10300", expected: `{"extenso":"menos dez mil e trezentos"}`},
		{url: "/10401", expected: `{"extenso":"dez mil e quatrocentos e um"}`},
		{url: "/-10501", expected: `{"extenso":"menos dez mil e quinhentos e um"}`},
		{url: "/-99999", expected: `{"extenso":"menos noventa e nove mil e novecentos e noventa e nove"}`},
		{url: "/031985", expected: `{"extenso":"trinta e um mil e novecentos e oitenta e cinco"}`},
		{url: "/70001", expected: `{"extenso":"setenta mil e um"}`},
		{url: "/71001", expected: `{"extenso":"setenta e um mil e um"}`},
		{url: "/71101", expected: `{"extenso":"setenta e um mil e cento e um"}`},
		{url: "/71111", expected: `{"extenso":"setenta e um mil e cento e onze"}`},
		// weird/wrong input
		{url: "/abcd", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
		{url: "/-999999", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
		{url: "/999999", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
		{url: "/99.999", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
		{url: "/99,999", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
		{url: "/0000099999", expected: `{"extenso":"noventa e nove mil e novecentos e noventa e nove"}`},
		{url: "/-0000099999", expected: `{"extenso":"menos noventa e nove mil e novecentos e noventa e nove"}`},
		{url: "/00000-99999", expected: `{"erro":"O path deve ser um numero inteiro entre -99999 e 99999."}`},
	}

	mux := server.CreateServerMux()

	for _, test := range tests {
		if test.expectedStatus == 0 {
			test.expectedStatus = http.StatusOK
		}

		t.Run(test.url, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, test.url, nil)
			mux.ServeHTTP(rec, req)

			// check if the http status is as expected
			if rec.Code != test.expectedStatus {
				t.Fatalf("Expected http status '%d', but received '%d'", test.expectedStatus, rec.Code)
			}

			// check if the http body is as expected
			body := strings.TrimSpace(rec.Body.String())
			if body != test.expected {
				t.Fatalf("Wrong body result\nExpected: '%s'\nReceived: '%s'", test.expected, body)
			}

		})
	}
}
