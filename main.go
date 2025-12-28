package main

import (
	"net/http"
	"net/url"
	"time"
	"fmt"
	"context"
	"log"
	"io"
)

const catalogBase = "https://go3.lv/api/products/vods/filtering"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func buildCatalogURL(tenant string, maxResults int) string {
	url, e := url.Parse(catalogBase)
	if e != nil {
		panic(e)
	}
	parametrs := url.Query()
	parametrs.Set("platform", "BROWSER")
	parametrs.Add("type[]", "VOD")
	parametrs.Set("tenant", tenant)
	parametrs.Set("maxResults", fmt.Sprintf("%d", maxResults))
	url.RawQuery = parametrs.Encode()
	return url.String()
}


func fetchBytes(ctx context.Context, client HTTPClient, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK  {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(snippet))
	}

	return io.ReadAll(resp.Body)
}



func main (){
	ctx, cancel := context.WithTimeout(context.Background(),15 * time.Second)
	defer cancel()
	url := buildCatalogURL("OM_LV",4121)
	fmt.Println("Req URL:", url)
	
	body, err := fetchBytes(ctx, httpClient, url)
	if err != nil {
		log.Fatalf("fetch failed: %v", err)
	}
	fmt.Printf("received %d bytes\n", len(body))
}
