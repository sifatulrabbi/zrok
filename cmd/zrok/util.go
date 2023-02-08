package main

import (
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func mustGetAdminAuth() runtime.ClientAuthInfoWriter {
	adminToken := os.Getenv("ZROK_ADMIN_TOKEN")
	if adminToken == "" {
		panic("please set ZROK_ADMIN_TOKEN to a valid admin token for your zrok instance")
	}
	return httptransport.APIKeyAuth("X-TOKEN", "header", adminToken)
}

func parseUrl(in string) (string, error) {
	// parse port-only urls
	if iv, err := strconv.ParseInt(in, 10, 0); err == nil {
		return fmt.Sprintf("http://127.0.0.1:%d", iv), nil
	}

	// make sure either https:// or http:// was specified
	if !strings.HasPrefix(in, "https://") && !strings.HasPrefix(in, "http://") {
		in = "http://" + in
	}

	// parse the url
	targetEndpoint, err := url.Parse(in)
	if err != nil {
		return "", err
	}

	return targetEndpoint.String(), nil
}
