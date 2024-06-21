# recipiece-ai-service

A microservice that acts as a proxy for interacting with ChatGPT. This is meant to securely communicate with ChatGPT so that API Keys are not exposed on the Svelte App called recipiece.

## Running the Microservice

``` go run .```


### Interacting with the Microservice

#### URI:
``` POST \hello ```

#### Body:

``` { message: "Hello" } ```
