+++
title = "AsyncClient: python quick example"
categories = ["zet"]
tags = ["zet"]
slug = "AsyncClient:-python-quick-example"
date = "2024-01-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# AsyncClient: python quick example

I have started a new job in another python shop.

Over the last 6-12 months, I wrote a lot more Go and yaml than python as my primary role
in my last job was kubernetes and infra management. So I am reacquainting myself
in python.

This is a quick async snippet for future reference, stolen mostly from Rednafi's
[blog post](https://rednafi.com/misc/eschewing_black_box_api_calls/) and changed to
suit my needs.

This code will fetch multiple `product_id`'s from an API using asyncio and httpx.
Pydantic is used to marshall things nicely.

```python
import httpx
from pydantic import BaseModel
import asyncio

urls = [1, 2, 3, 4, 5, 6, 7, 8, 9]


async def fetch_data(product_id: int):
    async with httpx.AsyncClient() as client:
        response = await client.get(f"https://dummyjson.com/products/{product_id}")
        response.raise_for_status()
        data = response.json()
        product = Product(**data)
        print(f"{product.id}: {product.title}, {product.description}")


class Product(BaseModel):
    id: int
    title: str
    description: str


async def main() -> None:
    tasks = []
    for _id in urls:
        tasks.append(fetch_data(_id))
    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
```

Example output:

```shell
2: iPhone X, SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...
4: OPPOF19, OPPO F19 is officially announced on April 2021.
3: Samsung Universe 9, Samsung's new variant which goes beyond Galaxy to the Universe
6: MacBook Pro, MacBook Pro 2021 with mini-LED display may launch between September, November
9: Infinix INBOOK, Infinix Inbook X1 Ci3 10th 8GB 256GB 14 Win10 Grey – 1 Year Warranty
1: iPhone 9, An apple mobile which is nothing like apple
8: Microsoft Surface Laptop 4, Style and speed. Stand out on HD video calls backed by Studio Mics. Capture ideas on the vibrant touchscreen.
7: Samsung Galaxy Book, Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched
5: Huawei P30, Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.
```

Tags:

  #python #async #httpx
