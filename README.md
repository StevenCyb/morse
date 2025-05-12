# Morse
This is a simple morse translator written in go. 
It can be used like:
```bash
# Text to Morse
$ go run main.go send "hello world"
.... . .-.. .-.. ---  .-- --- .-. .-.. -..
# Morse to Text
$ go run main.go read ".... . .-.. .-.. ---  .-- --- .-. .-.. -.."
HELLO WORLD
```
