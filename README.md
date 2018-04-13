# Simplest Docker image for starting proxy for Telegram 

By ex Telegram developer that just want to keep talking to family and friends.

## Step 1: Start Proxy

`
docker run --restart always -d --name telegram-proxy -p 1080:1080 ex3ndr/telegram-proxy
`

## Step 2: Test Proxy
https://t.me/socks?server=127.0.0.1&port=1080&user=user&pass=password

# Licensing

Public Domain
