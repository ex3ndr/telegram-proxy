# Simplest Docker image for starting proxy for Telegram 

By ex Telegram developer that just want to keep talking to family and friends.

## Step 1: Start Proxy

```
docker run --restart always -d --name telegram-proxy -p 1080:1080 ex3ndr/telegram-proxy
```

or with custom credentials
```
docker run --restart always -d --name telegram-proxy -v SOCKS_USER=telegram -v SOCKS_PASSWORD=telegram -p 1080:1080 ex3ndr/telegram-proxy
```

or if you want to allow only specific ip addresses as destination for proxy
```
docker run --restart always -d --name telegram-proxy -e "IP_ADDRESSES=8.8.8.8,8.8.4.4,8.8.8.0/22" -p 1080:1080 ex3ndr/telegram-proxy
```

## Step 2: Test Proxy
https://t.me/socks?server=127.0.0.1&port=1080&user=user&pass=password

## Step 3 (Optional): Contribute your proxy to our pool.

Email me (korshakov.stepan@gmail.com) with public IP address of your server and i will include your address to our community-supported pool.

# Licensing

Public Domain
