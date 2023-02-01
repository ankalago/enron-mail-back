# Test

## Paul Jácome - Quito / Ecuador - Software Web Developer

**[Linkedin Profile](https://bit.ly/paul-jacome-linkedin)**

**[GitHub Profile](https://bit.ly/paul-jacome-github)**

**[Bitbucket Profile](https://bit.ly/paul-jacome-bitbucket)**

**[Whatsapp Profile](https://bit.ly/paul-jacome-whatsapp)**

![Product_Png](https://raw.githubusercontent.com/ankalago/enron-mail-back/main/screenshot.png)

## Download emails

```
mkdir emails && cd emails && wget http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz && gzip -d  olympics.ndjson.gz && cd ..
```

## Load data like cURL (as documentation)

```
curl https://api.zincsearch.com/api/enterprise_16O476E9609y7Lz/_bulk -i -u ankalago@gmail.com:Qb7xFe2169z8g4y30c5r --data-binary "@enron_mail_20110402.json"
```

## Install dependencies

```
go mod tidy
```

## Create logs folder

```
mkdir logs
```

## Run indexer (with 250 go routines)

```
go run indexer.go
```

## Run server and endpoint

```
go run main.go
```

### Without user

`
http://localhost:3000/feed
`

### With user

`
http://localhost:3000/feed/arora-h
`

## Run profiling

```
cd testing && go test
```

```
go test -cpuprofile=cpu.out
```

```
go tool pprof cpu.out
```

(proff) > `list TestIndex`
![Product_Png](https://raw.githubusercontent.com/ankalago/enron-mail-back/main/images/screenshot2.png)

(proff) > `web or pdf`
![Product_Png](https://raw.githubusercontent.com/ankalago/enron-mail-back/main/images/screenshot1.png)
