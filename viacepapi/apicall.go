package viacepapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/go-chi/chi/v5"
	"github.com/lgustavopalmieri/labs-challenge-deploy-cloud-run/weatherapi"
)

type ResultViaCep struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

func validateCEP(cep string) bool {
	regex := regexp.MustCompile(`^\d{8}$|^\d{5}-\d{3}$`)
	return regex.MatchString(cep)
}

// GetTemperature godoc
// @Summary      GetTemperature
// @Description  This API is designed to provide a seamless integration of location and weather information 
// @Tags         temperature
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true "cep"
// @Success      200   object  string   "successful response"
// @Failure      422   object  string   "invalid zipcode"
// @Failure      404   object  string   "can not find zipcode"
// @Router       /temperature/{cep} [get]
func HandleTemp(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if !validateCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusNotFound)
		log.Println("invalid zipcode", cep)
		return
	}
	url1 := "http://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data ResultViaCep
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data.Erro == "true" {
		http.Error(w, "cannot find zipcode", http.StatusNotFound)
		log.Println("cannot find zipcode", cep)
		return
	}

	cidade := data.Localidade
	temperatura, _ := weatherapi.GetWeather(url.QueryEscape(cidade))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"temperature_C": temperatura.Celsius,
		"temperature_F": temperatura.Fahrenheit,
		"temperature_K": temperatura.Kelvin,
	})
}
