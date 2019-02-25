package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
	"strings"
)

const FILE_PROPERTIES = "individupro.properties"

const URL_DS008_PROD = "url ds008"

func main() {
	fmt.Printf("Récupération des logins des espaces recruteurs \n")

	fmt.Printf("\t1 - lecture des id individu pro dans le fichier de properties \n")

	var ids []string

	ids, _ = readPropertiesFile()

	fmt.Println("\t ==>", len(ids), "identifiants récupérés")

	fmt.Printf("\t2 - appel au DS008 pour récupérer les logins des espace recruteurs des établissements\n")

	callDs008AndWriteInFile(ids)

	fmt.Printf("Fin du traitement.")
}

func callDs008AndWriteInFile(ids []string) {
	f, err := os.Create("result.csv")

	if nil != err {
		log.Fatal(err)
	}

	// Header of file
	f.WriteString("idindividupro;login;\n")

	// Call and Write.
	for i := 0; i < len(ids); i++ {
		f.WriteString(string(ids[i]) + ";")
		f.WriteString(callDs008ForIndividuPro(ids[i]) + ";")
		f.WriteString("\n")

	}

}

func callDs008ForIndividuPro(idIP string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL_DS008_PROD + "/v1/employeurs/comptes/" + idIP, nil)

	req.Header.Add("pe-id-correlation","1-9809809")
	req.Header.Add("pe-id-utilisateur","ijba1660")
	req.Header.Add("pe-nom-domaine","DS008")
	req.Header.Add("pe-id-environnement","ENV_PO002")
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-Type","application/json")

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("Erreur lors de la récupération de l'individu pro %r %s", idIP, err)
		return "Erreur lors de l'appel au DS008;"
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return getLoginFromResponseJSON(contents)
	}
}
func getLoginFromResponseJSON(response []byte) string {

	responseS := string(response)

	//indexOfLogin := strings.Index(responseJSON, "loginMnemotechnique")
	positionLogin := strings.Index(responseS, "\"loginMnemotechnique\":")

	login := response[positionLogin:]

	positionLoginVirguleEnd := strings.Index(string(login), ",")

	loginFinish := login[:positionLoginVirguleEnd]
	loginFinishWithoutQuote := strings.Replace(string(loginFinish), "\"", "", -1)
	positionPoints := strings.Index(string(loginFinishWithoutQuote), ":")
	loginFinishWithoutQuoteOnlyLogin := loginFinishWithoutQuote[positionPoints+1:]

	return string(loginFinishWithoutQuoteOnlyLogin)
}

func readPropertiesFile() ([]string, error) {
	var data []string

	file, err := os.Open(FILE_PROPERTIES)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, err
}
