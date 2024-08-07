package goucan

import "testing"

func TestUcan(t *testing.T) {
	t.Log("Custom test")
	token := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCIsInVjdiI6IjAuOC4xIn0.eyJhdWQiOiJkaWQ6a2V5OnphYmNkZS4uLiIsImF0dCI6W3sid2l0aCI6ImJ1Y2s6YWxwaGEiLCJjYW4iOiJzdWNoL3ByaXZhdGUvaG9seS9obW1idWcifV0sImV4cCI6MTcyMjkyOTMyNywiaXNzIjoiZGlkOmtleTp6Nk1raVlXOVVMaVdIQ1BZM0RpQWs1UERMRG5LcGZrYXZURDJnWXJ5c2RRMkE5MlAiLCJwcmYiOltdfQ.ftvTMTq0P73ngfzPtw6Coacx_3HX3KLtoWJVv4wh1bZbJq185sCsaZLWaZWKd0RAQpExAN66JZZPsmbAjb5UAA"
	ucan, err := DecodeUcanString(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ucan.Header)
	t.Log(ucan.Payload)

	checkPayload := &UcanPayload{
		Iss: "did:key:z6MkiYW9ULiWHCPY3DiAk5PDLDnKpfkavTD2gYrysdQ2A92P",
		Aud: "did:key:zabcde...",
		Att: []Capabilities{
			{
				With: map[string]interface{}{
					"schema":   "buck",
					"hierPart": "alpha",
				},
				Can: map[string]interface{}{
					"namespace": "such",
					"segments":  []string{"private", "holy", "hmmbug"},
				},
			},
		},
	}

	ok, err := ucan.Verify(checkPayload)
	if err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatal("Verification failed")
	}

	t.Log(ucan)
}

/*
 capability:
*/
