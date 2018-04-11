package padArchiver

import (
	"fmt"
	"os"

	sh "github.com/ipfs/go-ipfs-api"
)

//GettedPads is the directory where are stored the pads that are getted from IPFS
const IpfsStorage = "ipfsStorage"
const GettedPads = "ipfsStorage/gettedPads"

//Add gets the content from the etherpad specified in the link, and downloads it in the format of the specified extension, and then, puts it into IPFS
func IpfsAdd(path string) (string, error) {
	//connect to ipfs shell
	s := sh.NewShell("localhost:5001")
	//save the file into IPFS
	ipfsHash, err := s.AddDir(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ipfsHash, nil
}

//Get gets the content from IPFS for a given hash, and saves it into a file
func IpfsGet(hash string, filename string) error { //create the pads directory
	//create the pads directory
	_ = os.Mkdir(IpfsStorage, os.ModePerm)
	_ = os.Mkdir(GettedPads, os.ModePerm)

	//connect to ipfs shell
	s := sh.NewShell("localhost:5001")
	err := s.Get(hash, GettedPads+"/"+filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
