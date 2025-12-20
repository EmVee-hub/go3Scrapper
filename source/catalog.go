package catalog

import (
	"context"
	"net/http"
	"time"
	"fmt"
)

//FIXME: NAMING CATALOG_ENDPOINT
const CATALOG_ENDPOINT= "https://go3.lv/api/products/vods/filtering?platform=BROWSER&type[]=VOD&tenant=OM_LV&maxResults=4121"


var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func buildCatalogURL() string {
	return CATALOG_ENDPOINT
}


func fetchBytes(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url,nil)
	if err != nil {
		return nil, fmt.ErrorF("creating request: %w", err)
	}
}

