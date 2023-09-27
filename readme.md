# TIS API exposes data from TIS system and allows insert and update outages and manipulation

API that exposes data about EEE function locations and equipment and signals of protective devices. Enables input, updating and deletion of outages and manipulation in the TIS system

- Building in GO version go1.18.2 windows/amd64
- Database is Oracle 
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [crypto](https://golang.org/x/crypto) 
- Uses [golang-jwt](https://github.com/golang-jwt/jwt/v4) and [pascaldekloe](https://github.com/pascaldekloe/jwt) to generate JWT tokens
- Uses [godror Oracle driver](https://github.com/godror/godror)