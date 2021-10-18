# Token Replace
This CLI tool is created to help CI/CD pipeline developers to locally test functions normally used in pipelines where token replacements are being used.  
The tool replaces all tokens found in the input file with the environment variables of similar name.
## Syntax
Using a file as input:
```shell
$ TEST=foobar token_replace -f examples/standard.yaml
value: foobar
```
Using pipe as input:
```shell
$ cat examples/standard.yaml | TEST=foobar token_replace -f -
value: foobar
```

It uses the token format as `#{TOKEN_NAME#}#`. It is also possible to your own custom tokens.  
You can set these using the following flags `startToken` and `endToken`.
Example:
```shell
$ TEST=foobar token_replace -f examples/custom.yaml -startToken "_+" -endToken "+_"
value: foobar
```