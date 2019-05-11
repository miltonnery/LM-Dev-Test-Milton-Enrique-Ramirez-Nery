package gateway

import (
	"lb_authorization_svc/configs"
	"lb_authorization_svc/internal/model/config"
	"regexp"
	"strings"
)

func Find(reqMethod string, reqUri string) (bool, *config.Uri) {
	if !strings.HasSuffix(reqUri, "/") {
		reqUri += "/"
	}
	path, req := splitPath(reqUri)
	if req == "/" {
		req = ""
	}
	uri := configs.Endpoints[path]
	if (uri.Method == reqMethod) && (uri.Enabled) {
		if uri.Method == "GET" {
			uri.RedirectionPath = uri.RedirectionPath + req
		}
		return true, &uri
	} else {
		return false, &config.Uri{}
	}
}

func splitPath(reqUri string) (string, string) {
	lbPattern, _ := regexp.Compile("((?:\\/lifebank\\/[^\\/]+))(\\S+)")

	//r := lbPattern.FindStringSubmatch(reqUri)
	//fmt.Print(r)

	if lbPattern.MatchString(reqUri) {
		return lbPattern.FindStringSubmatch(reqUri)[1], lbPattern.FindStringSubmatch(reqUri)[2]
	}

	return "base_path_not_found", "no_request_params_found"
}
