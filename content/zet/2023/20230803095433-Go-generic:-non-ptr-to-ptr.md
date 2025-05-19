+++
title = "Go generic: non-ptr to ptr"
categories = ["zet"]
tags = ["zet"]
slug = "go-generic:-non-ptr-to-ptr"
date = "2023-08-03 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go generic: non-ptr to ptr

I'm using `oapi-codegen` and any `required` fields are set as pointers. This is great
for doing `nil` checks. Initialising structs was more difficult because you cannot
create a struct with `*string` fields like this:

```golang
type Foo struct {
    ID *string
}

func main() {
    f := Foo{ID: "123"}
}
```

This fails because its not a pointer. Usually I'd create a variable like `var s *string` and then
populate the struct using `s`.

Then I found this great use of generics and am now using it everywhere.

```golang
// Ptr takes in non-pointer and returns a pointer
func Ptr[T any](v T) *T {
	return &v
}

// in use 
// a snippet from my tests
{name: "200: OK", method: "GET", url: "/api/v1/devices/device_000000000000", body: "", status: 200,
    want: repository.DeviceResponseJSON{
        Data: oapi.DeviceResponse{
            ApiVersion:      Ptr("not provided"),
            CreatedBy:       Ptr("guest"),
            DeviceActivated: Ptr(false),
            DeviceId:        Ptr("device_000000000000"),
            DeviceName:      Ptr("test_device"),
            DeviceVersion:   Ptr("2.7"),
            HostAddress:     Ptr("192.168.0.1"),
            HttpPort:        Ptr(int32(8080)),
            SshPort:         Ptr(int32(22)),
            TeamId:          Ptr("team_000000000000"),
        },
    },
},
```

Just made writing tests so much easier.

Tags:

    #TIL #generics #golang
