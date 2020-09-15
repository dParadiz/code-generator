# Code generator 

Code generator is meant to create boilerplate code for sever and client side api request/response handling.
Structure of the generator is based around encoder/decoder model. Where encoders encode specific input to context from which 
decoder can generate specific boilerplate code. 

## Encoders Decoders

Encoder and decoders are plugins for generator. 

- OpenApi encoder using `github.com/getkin/kin-openapi/openapi3` to prepare context for decoding

## Coder rendering

Code rendering is done with mustache templates

- Psr15Api decoder that decodes OpenApi context into php api based on https://github.com/dParadiz/api-frame with generated routing, middleware, handlers and models