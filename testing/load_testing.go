package testing

import (
	"fmt"
	"gometer/helpers"
	"net/http"
	"os"
	"sync"
	"time"
)

type LoadTester struct{
	url string
	concUsers int
	method string
}


func NewLoadTester(url string,concUsers int,method string) LoadTester{
	return LoadTester{
		url:url,
		concUsers:concUsers,
		method:method,
	}
}



func Test_load(loadTestingParams LoadTester){
	var wg sync.WaitGroup
	hostname,hostnameExtractionErr := helpers.ExtractHostname(loadTestingParams.url)

	if hostnameExtractionErr != nil {
		fmt.Printf("Error extracting hostname: %v", hostnameExtractionErr)
	}
	logFileName := constructLogFileName(hostname)
	createResultLogFileIfDoesNotExist(logFileName)
	for i:=0;i<loadTestingParams.concUsers;i++{
		wg.Add(1)
		go callUrl(loadTestingParams.url,loadTestingParams.method,&wg,logFileName)
	}
	wg.Wait()
}

func callUrl(url string,method string, wg *sync.WaitGroup,logFileName string){
	defer func ()  {
	wg.Done()	
	}()

	fmt.Printf("Calling url %s with method %s\n",url,method)
	var res *http.Response
	var err error

	startTime:=time.Now()

	switch method {
		case "GET":
			res,err = http.Get(url)

		// case "POST":
		// 	res,err = http.Post(url)
		default:
			fmt.Println("Invalid method")
	}

	responseTime := time.Since(startTime)
	

	if err != nil {
		fmt.Printf("Error at http call: %v", err)
	}

	logResult(res.Status, float32(responseTime)/float32(time.Second),url,method,logFileName)




}
// will implement log error function

func logResult(status string, responseTime float32,url string,method string,logFileName string){


	file, err := os.OpenFile(fmt.Sprintf("results/%s", logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	if _, err := file.WriteString(fmt.Sprintf("%s %f %s %s\n",status, responseTime,url,method)); err != nil {
		fmt.Printf("Error writing to file: %v", err)
	}

}



func constructLogFileName(hostnameToTest string) string{
	timestamp := helpers.GetFormattedTimeStampString()
	return fmt.Sprintf("%s-%s.log",hostnameToTest,timestamp)
}

func createResultLogFileIfDoesNotExist(filename string) {
	// filePath := fmt.Sprintf("results/%s", filename)
	dirName := "results"
	_,dirNotExistErr:=os.Stat(dirName)
	
	if dirNotExistErr != nil{
		os.Mkdir(dirName, 0755)
	}

	filePath := fmt.Sprintf("%s/%s",dirName, filename)

	_, filePathExistsErr := os.Stat(filePath)
	if filePathExistsErr != nil {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v", err)
		}
		defer file.Close()
	}else{
		fmt.Printf("File %s already exists",filename)
	}



 

}