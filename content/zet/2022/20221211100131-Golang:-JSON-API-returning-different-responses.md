+++
title = "Golang: JSON API returning different responses"
categories = ["zet"]
tags = ["zet"]
slug = "Golang:-JSON-API-returning-different-responses"
date = "2022-12-11 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Golang: JSON API returning different responses

I am interfacing with an API which returns different responses from the endpoint
dependant on a success or error. I want to be able to handle every response and this
is how I did that.

A successful response looks something like:

```go
[
  {
    "id": 1,
    "hop": 1,
    "url": "https://danielms.site",
    "http_version": "HTTP/1.1",
    "status_code": {
      "code": "200",
      "phrase": "OK"
    }
  }
]
```

And an error:

```go
{
  "detail": {
    "error": "The URL could not be resolved.",
    "url": "https://danielms.site2",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36"
  }
}
```

In order to handle both events I created response structs for each and 
then added those a generic response struct. The generic response struct I
could then create a `UnmarshalJSON` method for which fulfils the `Unmarshaller` 
interface. 

Here is the struct and method.

```go
type ResponseTypes struct {
	Response      []RedirectResponse
	ErrorResponse ErrResponse
}

func (d *ResponseTypes) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	var m map[string]any
	var unmarshalErr error

	if err := json.Unmarshal(data, &m); err != nil {
		unmarshalErr = err
	}

	if _, ok := m["detail"]; ok {
		var errData ErrResponse
		err := json.Unmarshal(data, &errData)
		if err != nil {
			return err
		}
		d.ErrorResponse = errData
		return nil
	}

	var arr []map[string]any
	if err := json.Unmarshal(data, &arr); err != nil {
		unmarshalErr = err
	}

	if _, ok := arr[0]["status_code"]; ok {
		var respData []RedirectResponse
		err := json.Unmarshal(data, &respData)
		if err != nil {
			return err
		}
		d.Response = respData
		return nil
	}
	return unmarshalErr
}
```

The hard part was handling both slice and single element structs at the 
same time. Initially I had issues because in order to inspect the response
I needed to unmarshal it. There might be a better way but above I've just
unmarshalled again if there isn't any error.

Because Go is awesome, I don't have to do anything out of the ordinary
to use the `ResponseTypes.UnmarshalJSON` method. The interface is fulfilled
so I just call it as I would any other `UnmarshalJSON` method call.

```go
// used in another function
r, _ := io.ReadAll(response.Body)
var result ResponseTypes
err = json.Unmarshal(r, &result)
if err != nil {
  return err
}
fmt.Println(result)
```

Tags:

    #go #TIL
