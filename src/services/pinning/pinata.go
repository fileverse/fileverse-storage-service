package main

import (
	"fmt"
	"net/http"

	"github.com/wabarc/ipfs-pinner/pkg/pinata"
)

var (
	pinataApiKey = "73d01142fc4f2ffdb1e8"
	pinataSecret = "e6396be9b0eb2a07b027e1bcdcfe78a8c6dfb33d138b377baaa44a51a35ab05d"
	pinataJwt    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiI1Zjk5YTljMy05M2M0LTQxMmYtYTMxZi03Y2Q1ZTRjNTg2NDYiLCJlbWFpbCI6ImFudWJoYXZAZmlsZXZlcnNlLmlvIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjEsImlkIjoiRlJBMSJ9LHsiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjEsImlkIjoiTllDMSJ9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjczZDAxMTQyZmM0ZjJmZmRiMWU4Iiwic2NvcGVkS2V5U2VjcmV0IjoiZTYzOTZiZTliMGViMmEwN2IwMjdlMWJjZGNmZTc4YThjNmRmYjMzZDEzOGIzNzdiYWFhNDRhNTFhMzVhYjA1ZCIsImV4cCI6MTc1Mzc4NDAxMn0.AS10i98PyqSjNVVp6j8x1L26s7pvVE9oTS4WpjbdNvE"
)

func unpinFile(cid string) (bool, error) {
	url := "https://api.pinata.cloud/pinning/unpin/" + cid

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+pinataJwt)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	if (res.StatusCode != http.StatusOK) && (res.StatusCode != http.StatusAccepted) {
		return false, fmt.Errorf("unpin failed: %d", res.StatusCode)
	}
	return true, nil
}

func main() {
	pnt := pinata.Pinata{
		Apikey: pinataApiKey,
		Secret: pinataSecret,
	}

	cid, err := pnt.PinFile("./go.mod")
	if err != nil {
		fmt.Sprintln(err)
		return
	}

	fmt.Printf("https://blush-calm-lobster-485.mypinata.cloud/ipfs/%s\n", cid)

	res, err := unpinFile(cid)
	if err != nil {
		fmt.Println("failed to unpin: ", err)
	} else {
		fmt.Println("unpin result: ", res)
	}

}
