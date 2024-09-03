+++
title = "How to invalidate tanstack query key in same component"
categories = ["zet"]
tags = ["zet"]
slug = "How-to-invalidate-tanstack-query-key-in-same-component"
date = "2024-09-03 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# How to invalidate tanstack query key in same component

I am new to TanStack query, router, form etc.

I wanted to have a table and a form in the same component (or page containing
many components).

Create a thing in the form, invalidate the cache and have the table re-fetch
table data rendering the table with the new data.

At the same time I am using Orval to generate all my TanStack Query (TQ from now
on) methods.

It took me a while to figure this out but AFAICT, you do not invalidate a cache
from Orval's generated spec. You instead use TQ directly.

```tsx
export function Foo() {
    const queryClient = useQueryClient()
    const { status, data } = useTQListMethod()
    const { mutate } = useMutateTQListMethod({
        mutation: {
            onSuccess: () => {
                // this is how I invalidated the cache
                queryClient.invalidateQueries({
                    queryKey: ["/formdata"] // matches the queryKey from useTQListMethod
                })
            }
        }
    })

    if (status === "pending") return <p>Loading...</p>;
    if (status === "error") return <p>Error :(</p>;


    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        const formData = new FormData(e.currentTarget)
        // truncated logic

        // does the mutation (sending POST data) but does not invalidate the cache!
        mutate({data: myData})
    }
    return (
    <>
        <form onSubmit={handleSubmit}>This is truncated and not that important</form>
        <MyTable data={data.data}
    </>
    )
}
```

Useing `queryClient.invalidateQueries` let me force my cache query to refetch
and successfully re-render the table (`MyTable`) with the new entry.

Tags:

    #react #tanstack
