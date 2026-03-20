package example

type ContextWrapper struct{} // ok

type CtxWrapper struct{} // want `avoid "ctx" in name, use a more specific term`

type ReqBody struct{} // want `use "request" instead of "req" in ReqBody`

type HandlerServer struct{} // want `avoid "handler" in name, use a more specific term`
