# GoStart
A start page loader using FCGI in Go.


Uses FCGI to run:

```go
func main() {
	http.HandleFunc("/", handler)
	if err := fcgi.Serve(nil, nil); err != nil {
		panic(err)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "This was generated by Go running as a FastCGI app")
}
```


Reads in files or templates to display 

```Go
func readFile(w http.ResponseWriter) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Fprintln(w, "File reading error", err)
		return
	}
	fmt.Fprintln(w, "Contents of file:", string(data))
}
```

Templates and other useful info: https://golang.org/doc/articles/wiki/

Download the files to display (update button?)

Github Zip?

```Go
// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
   }
   defer resp.Body.Close()
   
   // Create the file
	 out, err := os.Create(filepath)
   if err != nil {
    return err
   }
   
   defer out.Close()

  // Write the body to file
  _, err = io.Copy(out, resp.Body)
  return err
}
```
