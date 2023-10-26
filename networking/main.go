package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// *** Working with URLs ***

	surl := "https://www.example.com:8080/user?username=joemarini"

	result, _ := url.Parse(surl)
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("path:", result.Path)
	//fmt.Println("Raw Path:", result.RawPath)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw query:", result.RawQuery)

	vals := result.Query()
	fmt.Println("Query param username:", vals["username"])

	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}

	fmt.Println("New url string", newurl.String())

	newurl.Host = "joemarini.com"

	fmt.Println("New url with host updated:", newurl.String())

	// Encode query string
	newvals := url.Values{}
	newvals.Add("x", "100")
	newvals.Add("y", "some str value with &")

	newurl.RawQuery = newvals.Encode()

	fmt.Println("New url with encoded query:", newurl.String())

	//getRequestTest()
	//postRequestTest()
	//encodeJson()
	//decodeJson()
	//encodeXml()
	decodeXml()
}

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	resp, err := http.Get(httpbin)
	handleError(err)

	defer resp.Body.Close() // Caller needs to close response

	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Protocol:", resp.Proto)
	fmt.Println("Content Length:", resp.ContentLength)

	var sb strings.Builder
	content, _ := io.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)

	fmt.Println("Byte count:", bytecount)
	fmt.Println("Content:", sb.String())
}

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"
	reqBody := strings.NewReader(`{
		"field1": "This is field 1",
		"field2": 250
	}`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	handleError(err)

	defer resp.Body.Close() // Caller needs to close response

	content, _ := io.ReadAll(resp.Body)
	fmt.Printf("Post response:\n%s\n", content)

	// Post form
	data := url.Values{}
	data.Add("field1", "This is form field 1")
	data.Add("field2", "350")

	formResp, err := http.PostForm(httpbin, data)
	handleError(err)

	defer formResp.Body.Close() // Caller needs to close response

	formRespContent, _ := io.ReadAll(formResp.Body)
	fmt.Printf("Form Post response:\n%s\n", formRespContent)
}

// Fields need to be public for json serialization to work
type person struct {
	XMLName   xml.Name `xml:"persondata"`
	Name      string   `json:"fullname" xml:"full_name"`
	Address   string   `json:"addr" xml:"address"`
	Age       int      `json:"age" xml:"age,attr"`
	FavColors []string `json:"favcolors,omitempty" xml:"favourite_colors"`
	SSN       string   `json:"-" xml:"-"` // Do not serialize this field
}

func encodeJson() {
	people := []person{
		{xml.Name{Local: "persondata"}, "Jane Doe", "123 Anywhere street", 35, nil, "123456"},
		{xml.Name{Local: "persondata"}, "Jane Public", "456 Everywhere street", 31, []string{"Purple", "Yellow"}, "789012"},
	}

	//jsonStr, err := json.Marshal(people)
	jsonStr, err := json.MarshalIndent(people, "", "\t")
	handleError(err)

	fmt.Printf("%s\n", jsonStr)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func decodeJson() {
	data := []byte(`{
		"fullname": "John Q Public",
		"addr": "987 Main St",
		"age": 45,
		"favcolors": ["Purple", "Gold"]
	}`)

	var p person
	valid := json.Valid(data)

	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

	// Decode in a map
	var jsonMap map[string]interface{}
	json.Unmarshal(data, &jsonMap)
	fmt.Printf("%#v\n", jsonMap)

	for k, v := range jsonMap {
		fmt.Printf("key (%v), value (%T : %v)\n", k, v, v)
	}
}

func encodeXml() {
	people := []person{
		{xml.Name{Local: "persondata"}, "Jane Doe", "123 Anywhere street", 35, nil, "123456"},
		{xml.Name{Local: "persondata"}, "Jane Public", "456 Everywhere street", 31, []string{"Purple", "Yellow"}, "789012"},
	}

	//xmlStr, err := xml.Marshal(people)
	xmlStr, err := xml.MarshalIndent(people, "", "\t")
	handleError(err)

	fmt.Printf("%s%s\n", xml.Header, xmlStr)
}

func decodeXml() {
	data := []byte(`
		<persondata age="31">
				<full_name>Jane Public</full_name>
				<address>456 Everywhere street</address>
				<favourite_colors>Purple</favourite_colors>
				<favourite_colors>Yellow</favourite_colors>
		</persondata>
	`)

	var p person
	xml.Unmarshal(data, &p)
	fmt.Printf("%#v\n", p)
}
