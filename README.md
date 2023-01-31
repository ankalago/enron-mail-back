# Test

## Paul Jacome - Quito - Ecuador

**[Linkedin Profile](https://www.linkedin.com/in/pauljacome/)**

![Product_Png](https://raw.githubusercontent.com/ankalago/enron-mail-back/main/screenshot.png)

## Download emails

```
wget http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz
```

## Load data like cURL

```
curl https://api.zincsearch.com/api/enterprise_16O476E9609y7Lz/_bulk -i -u ankalago@gmail.com:Qb7xFe2169z8g4y30c5r --data-binary "@enron_mail_20110402.json"
```

## Install dependencies

```
go mod tidy
```

## Run indexer (with 250 go routines)

```
go run indexer.go
```

## Run server and endpoint
```
go run main.go
```

### With out user

`
http://localhost:3000/feed
`

### With user

`
http://localhost:3000/feed/arora-h
`
