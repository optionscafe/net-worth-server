# Net Worth Server

A RESTful API for the net worth application. Written in Golang. 

If your anything like me you are constantly adding money to your investments. Sometimes you take money out of your investments to pay for the extras in life. With all the movement of money in and out of your investment accounts it is hard to keep track how your doing overall. Are you an investment sage or merely matching the S&P 500? This net worth server application gives you an easy to use RESTful API for tracking your investments. Keep track of your investments on a "per share basis". Use these tools to figure out if your the next Warren Buffet.

The Net worth server is also helpful for keeping track of your investments across platforms. For example I like to know how I am doing from my brokerage, options, futures trading accounts. Lets throw in investments such as [Lending Club](https://www.lendingclub.com/), [Prosper](https://www.prosper.com), [Peer Street](https://www.peerstreet.com), [Realtyshares](https://www.realtyshares.com/). Heck lets track how our Real Estate investments are doing via [Zillow](https://www.zillow.com/). 

# Ansible Deployment 

There are some Ansible scripts in place to deploy this application to a docker server. You can find these scripts in ```ansible```. There are a few files you have to put into place as I did not want to commit them to git.

* ```ansible/vars/env.yml```
* ```ansible/hosts```

One idea would be to fork this repo remove these files from ```.gitignore``` and then encrypt these files with ansible-vault. 

To help deploys we have ```scripts/deploy.sh```. You will have to install ```scripts/.env``` for the script to work. 