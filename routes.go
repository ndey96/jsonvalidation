package jsonvalidation

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "DownloadSchema",
    "GET",
    "/schema/{schemaId}",
    DownloadSchema,
  },
  Route{
    "UploadSchema",
    "POST",
    "/schema/{schemaId}",
    UploadSchema,
  },
  Route{
    "ValidateDocument",
    "POST",
    "/validate/{schemaId}",
    ValidateDocument,
  },
}
