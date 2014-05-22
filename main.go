package main

import (
    "code.google.com/p/go.net/publicsuffix"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/cookiejar"
	"net/url"
)

const (
	loginUrl = "http://elas.las.vsbnet.be/postlogin.php"
	transactionUrl = "http://elas.las.vsbnet.be/transactions/posttransaction.php"
)

type ElasMember struct{

}

type ElasGateway struct{
	baseUrl String, 
	client http.Client,
	dataAdapter 
}

func (this *ElasGateway) members() ElasMember[]{

}

type MembersRepository struct{
}



func main() {
    options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        log.Fatal(err)
    }
	proxyUrl, err := url.Parse("http://amarisadsl:3128")
    client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, Jar: jar}
	login(client)
//	makeTransaction(client)
}

func makeTransaction(client *http.Client){
	values := url.Values{}
	values.Set("letsgroup", "1")
	values.Set("letscode_from", "140")
	values.Set("letscode_to", "138")
	values.Set("amount", "1")
	values.Set("minlimit", "-500")
	values.Set("balance", "500")
	values.Set("description", "automatische test")
    resp, err := client.PostForm(transactionUrl, values)
	printBody(resp, err)
}

func login(client *http.Client){
	values := url.Values{}
	values.Set("login", "timbosschaerts@gmail.com")
	values.Set("password", "Wortel7")
    resp, err := client.PostForm(loginUrl, values)
	printBody(resp, err)
}


func printBody(resp *http.Response, e error){
    if e != nil {
        log.Fatal(e)
    }
    data, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
	log.Println(string(data))
}
