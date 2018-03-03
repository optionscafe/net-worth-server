# Net Worth Server

A RESTful API for the net worth application. Written in Golang. 

If you're anything like me you are constantly adding money to your investments. Sometimes you take money out of your investments to pay for the extras in life. With all the movement of money in and out of your investment accounts it is hard to keep track how you're doing overall. Are you an investment sage or merely matching the S&P 500? This net worth server application gives you an easy to use RESTful API for tracking your investments. Keep track of your investments on a "per share basis". Use these tools to figure out if you're the next Warren Buffett.

The Net worth server is also helpful for keeping track of your investments across platforms. For example I like to know how I am doing from my brokerage, options, futures trading accounts. Let's throw in investments such as [Lending Club](https://www.lendingclub.com/), [Prosper](https://www.prosper.com), [Peer Street](https://www.peerstreet.com), [Realtyshares](https://www.realtyshares.com/). Heck let's track how our Real Estate investments are doing via [Zillow](https://www.zillow.com/). 

# Client Software For Engaging With Net Worth Server

Net worth server is not really useful without a client. Here is a list of clients to checkout. 

* [Net Worth CLI](https://github.com/optionscafe/net-worth-cli)

* [Net Worth Ionic](https://github.com/optionscafe/net-worth-ionic)

# Unit Testing 

With unit testing we setup a testing database (based off the value of the ```DB_DATABASE_TESTING``` OS variable). We run our tests at the highest levels. For example for an HTTP call (which is most of the app) we test the HTTP response controller. This way we are testing the models and helper packages that goes into generating the HTTP response, by doing it this way we get pretty great coverage. The goal is to cover the entire app just at the highest levels.  

# Ansible Deployment 

There are some Ansible scripts in place to deploy this application to a docker server. You can find these scripts in ```ansible```. There are a few files you have to put into place as I did not want to commit them to git.

* ```ansible/vars/env.yml```
* ```ansible/hosts```

One idea would be to fork this repo remove these files from ```.gitignore``` and then encrypt these files with ansible-vault. 

To help deploys we have ```scripts/deploy.sh```. You will have to install ```scripts/.env``` for the script to work. 

# Applications

Since this is a RESTful API only we have clients that connect. We define clients as "Applications". There is an applications db table that stores the name of the different applications. At this time we do not have any full blown oAuth this Applications system is the bare bones of a password grant type with oAuth 2. 

To get started enter the name of your application into the applications db table (or use the CMD tool below). Come up with a random 15 char (or so) client id and enter that into the db as well. 

When logging in (`oauth/token` route) pass along a `grant_type` and the `client_id` you set in the db. Here is an example POST

```
POST https://your-nw-server-domain.com/token
  grant_type=password&
  email=joe@example.com&
  password=foobar&
  client_id=CLIENT_ID
```   

We have a CMD tool for creating new applications `go run main.go -cmd=create-application -name="Ionic App"`

NOTE: By no means is client_id a form of security. This is an id users are allowed to see. It just verifies which application this request is designed for. 