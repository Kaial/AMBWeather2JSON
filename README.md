AMBWeather2JSON

Copyright 2022 Kyle Ceschi

Distributed under terms of the GPLv3

Simple http server that listens to an ambweather station and turns the data into json then POSTS it to a different server of your choosing. I wanted to learn a bit of Go and I needed a very simple utility to convert this weather stations data into json and send it off to another place for down stream processing. I've never written a go program before, so if the code looks weird that's why i'm still learning. I only tested with my specific version of the weather station (AMBWeatherV4.3.3) so your mileage may vary.



Usage 

Configure your Ambient weather station to send data to your server
 https://ambientweather.com/faqs/question/view/id/1857/
Load the code onto your server, edit the config file as needed and run it
`go run .` 
