package servererror

import (
	"io/ioutil"
	"net/http"
	"os"
)

func GetFileContent()  {
	http.HandleFunc("/practice/aes1/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/practice/aes1/"):]//  /practice/aes1/
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
