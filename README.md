# url-shortener
A simple url shortener written in Go, with multiple backing stores as choices, such as:
- Memory
- Redis
(More choices to come)

URLs are shortened to cgo.pl/abcdef (6-character alphanumeric). This can be changed.

# How to use
- To shorten a url use:
`curl http://localhost:8010/shorten?url=www.google.com`
- To resolve a short url use:
`curl http://localhost:8010/unmask?url=cgo.pl/abcdef`

# TODO
- Tests
- Add more backing stores (for eg a relational one)
