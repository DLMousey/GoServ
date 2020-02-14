package tokeniser

import (
	"regexp"
	"strings"
)

type HttpRequest struct {
	Path string
	Method string
	Version string
	Headers []RequestHeader
}

type RequestHeader struct {
	Name string
	Value string
}

func TokeniseRequest(request string) (httpRequest HttpRequest){

	regex, error := regexp.Compile("[^A-Za-z0-9:;/,.\\s]+")
	if error != nil {
		panic("Failed to compile regex! " + error.Error())
	}

	request = regex.ReplaceAllString(request, "")
	httpRequest.Headers = make([]RequestHeader, 0)

	lines := strings.Split(request, "\n")
	firstLine := lines[0]
	parts := strings.Split(firstLine, " ")

	httpRequest.Method = parts[0]
	httpRequest.Path = parts[1]
	httpRequest.Version = parts[2]

	for _, line := range lines {
		// If we're on the first line (don't know a better way of skipping this) of the headers
		// we'll skip it since we've already parsed this above
		if line == httpRequest.Method + " " + httpRequest.Path + " " + httpRequest.Version {
			continue
		}

		// If we've reached the separator between the headers and the body - we'll break out of this loop
		if line == "\r" || line == "" {
			continue
		}

		// Otherwise we must be at a header - so we'll process it
		var lineHeader RequestHeader

		lineParts := strings.Split(line, ":")
		lineHeader.Name = lineParts[0]
		lineHeader.Value = lineParts[1]

		httpRequest.Headers = append(httpRequest.Headers, lineHeader)
	}

	return
}
