package main

import (
	"database/sql"
	"log"
	"new/learning/user/api"
	db "new/learning/user/db/sqlc"
	"new/learning/user/util"
)

func main() {
	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}

	// Establish a database connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Printf("cannot connect to db: %v", err)
		return
	}
	defer conn.Close() // Close the database connection when main exits

	// Create a store instance using the database connection
	store := db.NewStore(conn)
	server,err := api.NewServer(config, store)
	if err != nil{
		log.Fatal("cannot connect to db", err)
	}

	// Create a new server instance by passing the configuration and store
	// server, err := api.NewServer(config, store)
	// if err != nil {
	// 	log.Printf("cannot create server: %v", err)
	// 	return
	// }

	// Start the server using the configured address
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Printf("cannot start server: %v", err)
		return
	}
}

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5432/user?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

// func main() {
// 	conn, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db", err)
// 	}

// 	store := db.NewStore(conn)
// 	server := api.NewServer(store)

// 	err = server.Start(serverAddress)
// 	if err != nil {
// 		log.Fatal("cannot start server", err)
// 	}
// }

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"

// 	shell "github.com/ipfs/go-ipfs-api"
// )

// func addFile(sh *shell.Shell, text string) (string, error) {
// 	return sh.Add(strings.NewReader(text))
// }

// func readFile(sh *shell.Shell, cid string) (*string, error) {
// 	reader, err := sh.Cat(fmt.Sprintf("/ipfs/%s", cid))
// 	if err != nil {
// 		return nil, fmt.Errorf("Error reaidng the file %s", err.Error())
// 	}
// 	bytes, err := io.ReadAll(reader)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error reaidng the bytes %s", err.Error())
// 	}
// 	text := string(bytes)
// 	return &text, nil
// }

// func main() {
// 	// Define the IPFS node connection details
// 	IPFS_NODE_HOST := "34.70.152.120"
// 	IPFS_NODE_PORT := "5001"
// 	IPFS_NODE_PROTOCOL := "http"

// 	// Construct the API endpoint URL
// 	ipfsURL := IPFS_NODE_PROTOCOL + "://" + IPFS_NODE_HOST + ":" + IPFS_NODE_PORT

// 	// Connect to the specified IPFS node
// 	sh := shell.NewShell(ipfsURL)

// 	// Add a string content to IPFS
// 	cid, err := sh.Add(strings.NewReader("hello world!"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("added %s\n", cid)

// 	// Get the data from IPFS and output the contents into `string` format
// 	data, err := sh.Cat(cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}

//		// Convert the data to a string
//		buf := new(bytes.Buffer)
//		_, err = buf.ReadFrom(data)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "error: %s", err)
//			os.Exit(1)
//		}
//		newStr := buf.String()
//		fmt.Printf("data: %s\n", newStr)
//	}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	shell "github.com/ipfs/go-ipfs-api"
// )

// // TimeSeriesDatum is the structure used to store a single time series data
// type TimeSeriesDatum struct {
// 	Id    uint64 `json:"id"`
// 	Value uint64 `json:"value"`
// }

// // YourLocalPath represents the local path on your computer
// const YourLocalPath = "./downloads" // Update this path as needed

// func main() {
// 	// Define the IPFS node connection details
// 	IPFS_NODE_HOST := "34.70.152.120"
// 	IPFS_NODE_PORT := "5001"
// 	IPFS_NODE_PROTOCOL := "http"

// 	// Construct the API endpoint URL
// 	ipfsURL := IPFS_NODE_PROTOCOL + "://" + IPFS_NODE_HOST + ":" + IPFS_NODE_PORT

// 	// Connect to the specified IPFS node
// 	sh := shell.NewShell(ipfsURL)

// 	// Create a TimeSeriesDatum object
// 	tsd := &TimeSeriesDatum{
// 		Id:    1,
// 		Value: 123,
// 	}

// 	// Marshal TimeSeriesDatum object to JSON and create a reader from it
// 	tsdBin, _ := json.Marshal(tsd)
// 	reader := bytes.NewReader(tsdBin)

// 	// Add the JSON content to IPFS
// 	cid, err := sh.Add(reader)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("added %s\n", cid)

// 	// Get the data from IPFS and unmarshal it into a TimeSeriesDatum object
// 	data, err := sh.Cat(cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %s", err)
// 		os.Exit(1)
// 	}

// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(data)
// 	newStr := buf.String()

// 	res := &TimeSeriesDatum{}
// 	json.Unmarshal([]byte(newStr), &res)
// 	fmt.Println(res)

// 	// Download file from IPFS to local path
// 	err = downloadFile(sh, cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error downloading file: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("File downloaded successfully")
// }

// // downloadFile downloads a file from IPFS to the specified local path
//
//	func downloadFile(sh *shell.Shell, cid string) error {
//		return sh.Get(cid, YourLocalPath)
//	}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"

// 	shell "github.com/ipfs/go-ipfs-api"
// 	"github.com/manifoldco/promptui"
// )

// // TimeSeriesDatum is the structure used to store a single time series data
// type TimeSeriesDatum struct {
// 	Id    uint64 `json:"id"`
// 	Value uint64 `json:"value"`
// }

// // YourLocalPath represents the local path on computer
// const YourLocalPath = "./downloads"

// func main() {
// 	// Define the IPFS node connection details
// 	IPFS_NODE_HOST := "34.70.152.120"
// 	IPFS_NODE_PORT := "5001"
// 	IPFS_NODE_PROTOCOL := "http"

// 	// Construct the API endpoint URL
// 	ipfsURL := IPFS_NODE_PROTOCOL + "://" + IPFS_NODE_HOST + ":" + IPFS_NODE_PORT

// 	// Connect to the specified IPFS node
// 	sh := shell.NewShell(ipfsURL)

// 	// Select a file to upload to IPFS
// 	filePath, err := selectFile()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error selecting file: %s", err)
// 		os.Exit(1)
// 	}

// 	// Read the selected file
// 	fileContent, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error reading file: %s", err)
// 		os.Exit(1)
// 	}

// 	// Add the file content to IPFS
// 	cid, err := sh.Add(bytes.NewReader(fileContent))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error uploading file to IPFS: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("File uploaded to IPFS with CID: %s\n", cid)

// 	// Download file from IPFS to local path
// 	err = downloadFile(sh, cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error downloading file: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("File downloaded successfully")
// }

// // selectFile prompts the user to select a file and returns its path
// func selectFile() (string, error) {
// 	prompt := promptui.Select{
// 		Label: "Select a file to upload",
// 		Items: listFiles("C:/Users/asus/Downloads"), // Update the directory path if needed
// 	}

// 	_, filePath, err := prompt.Run()

// 	if err != nil {
// 		return "", err
// 	}

// 	return filePath, nil
// }

// // listFiles lists all files in the given directory
// func listFiles(dirPath string) []string {
// 	var files []string

// 	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
// 		if !info.IsDir() {
// 			files = append(files, path)
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return files
// }

// // downloadFile downloads a file from IPFS to the specified local path

//	func downloadFile(sh *shell.Shell, cid string) error {
//		return sh.Get(cid, YourLocalPath)
//	}
// package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"

// 	shell "github.com/ipfs/go-ipfs-api"
// 	"github.com/manifoldco/promptui"
// )

// // TimeSeriesDatum is the structure used to store a single time series data
// type TimeSeriesDatum struct {
// 	Id    uint64 `json:"id"`
// 	Value uint64 `json:"value"`
// }

// // YourLocalPath represents the local path on the computer
// const YourLocalPath = "./downloads"

// type PinInfo struct {
// 	Type string
// }

// type PinStreamInfo struct {
// 	Cid  string
// 	Type string
// }

// type PinType string

// const (
// 	DirectPin    PinType = "direct"
// 	RecursivePin PinType = "recursive"
// 	IndirectPin  PinType = "indirect"
// )

// // Define the IPFS node connection details
// const (
// 	IPFS_NODE_HOST     = "34.70.152.120"
// 	IPFS_NODE_PORT     = "5001"
// 	IPFS_NODE_PROTOCOL = "http"
// )

// var ipfsURL = IPFS_NODE_PROTOCOL + "://" + IPFS_NODE_HOST + ":" + IPFS_NODE_PORT

// func main() {
// 	// Connect to the specified IPFS node
// 	sh := shell.NewShell(ipfsURL)

// 	// Select a file to upload to IPFS
// 	filePath, err := selectFile()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error selecting file: %s", err)
// 		os.Exit(1)
// 	}

// 	// Read the selected file
// 	fileContent, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error reading file: %s", err)
// 		os.Exit(1)
// 	}

// 	// Add the file content to IPFS
// 	cid, err := sh.Add(bytes.NewReader(fileContent))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error uploading file to IPFS: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("File uploaded to IPFS with CID: %s\n", cid)

// 	// Download file from IPFS to local path and obtain its URL
// 	downloadedURL, err := downloadFileAndGetURL(sh, cid)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error downloading file or getting URL: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("File downloaded successfully")
// 	fmt.Println("Downloaded File URL:", downloadedURL)

// 	// Pin the uploaded file
// 	err = pinFile(sh, cid, RecursivePin)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error pinning file on IPFS: %s", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("File pinned on IPFS")
// }

// // selectFile prompts the user to select a file and returns its path
// func selectFile() (string, error) {
// 	prompt := promptui.Select{
// 		Label: "Select a file to upload",
// 		Items: listFiles("C:/Users/asus/Downloads"), // Update the directory path if needed
// 	}

// 	_, filePath, err := prompt.Run()

// 	if err != nil {
// 		return "", err
// 	}

// 	return filePath, nil
// }

// // listFiles lists all files in the given directory
// func listFiles(dirPath string) []string {
// 	var files []string

// 	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
// 		if !info.IsDir() {
// 			files = append(files, path)
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return files
// }

// // downloadFileAndGetURL downloads a file from IPFS to the specified local path and returns its URL
// func downloadFileAndGetURL(sh *shell.Shell, cid string) (string, error) {
// 	localPath := filepath.Join(YourLocalPath, cid) // File will be downloaded with its CID as the filename

// 	if err := sh.Get(cid, localPath); err != nil {
// 		return "", err
// 	}

// 	// Constructing the URL for the downloaded file
// 	downloadedURL := fmt.Sprintf("%s/ipfs/%s", ipfsURL, cid)

// 	return downloadedURL, nil
// }

// // pinFile pins a file on IPFS with the specified pinning type
// func pinFile(sh *shell.Shell, cid string, pinType PinType) error {
// 	pinTypeStr := string(pinType)
// 	_, err := sh.Request("pin/add").Arguments(cid, "--recursive", "--progress=false", "--type="+pinTypeStr).Send(context.Background())
// 	return err
// }

// unpinFile unpins a file from IPFS
// func unpinFile(sh *shell.Shell, cid string) error {
// 	_, err := sh.Request("pin/rm").Arguments(cid).Send(context.Background())
// 	return err
// }
