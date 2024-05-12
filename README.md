AMBWeather2JSON

Copyright 2024 Kyle Ceschi

Distributed under terms of the GPLv3

Simple http server that listens to an ambweather station and turns the data into json then POSTS it to a different server of your choosing. I wanted to learn a bit of Go and I needed a very simple utility to convert this weather stations data into json and send it off to another place for down stream processing. I've never written a go program before, so if the code looks weird that's why i'm still learning. I only tested with my specific version of the weather station (AMBWeatherV4.3.3) so your mileage may vary.



Usage 

Configure your Ambient weather station to send data to your server
 https://ambientweather.com/faqs/question/view/id/1857/
Load the code onto your server, edit the config file as needed and run it
`go run .`

Project Goals:

the end goal of this project will have multiple inter connected components. 

The first component is the current POC implementation, which includes a go HTTP server that is 
able to recieve and process the data that the base station currently exports via HTTP calls, the 
data is encoded in the URI of the call. This system is functional however a bit jank, parsing lots of data out of the URI is annoying and messy. 

The next component will be a step to resolve issues in the POC, this component is going to access the RAW encoded radio data that comes from the weather station itself. The first step will be to establish what frequencies the system uses, what the protocol is, and then implement a system likely with GNU radio that is able to demodulate the data in a format that can be consumed by a go service, which will then translate that information into sometime more transferable like json, and then forward that data to consumers via the MQTT messaging protocol.

The final component will combine the above two components and get the data into home assistant for viewing. There is an existing project that solves this right now if you are looking for something that works now. There is a plugin available for home assistant that is able to use MQTT with home assistant to get the weather data displayed from the sensors. My project aims to expand that scope, and eliminate the base station entirely, by using the raw radio data from the weather vanes instrumentation. It will then use a similar pattern of translating it to a more standard format such as json, which can then be stored for historical data, and translated to messages for consumption by other services. For reference this is an existing solution https://github.com/neilenns/ambientweather2mqtt . 
