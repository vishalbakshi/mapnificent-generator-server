# mapnificent-generator-server
A Go server that runs the Mapnificent Generator script to prepare GTFS binary files

## Version 2
This version has the following (untested) functionality:

  - [x] Retrieve zip file from transitfeeds API
  - [x] Save the zip file to local memory
  - [x] Generate .bin file from zip file using `mapnificent.go`

If you would like to run this locally, make sure you have go installed (https://golang.org/dl/)

  1. Clone this repo
  2. Add a `bin` folder
  3. Set your $GOPATH to the root directory
  4. build the packages, `cd` into `src/server` and `go install`
  5. When you run the server, it will download the `gtfs.zip` file from the transitfeeds API and convert it to `bayarea.bin`
  
To test your .bin file in Mapnificent:

  1. In the `_cities/bayarea` folder, delete the file `bayarea.bin` and replace it with your new `bayarea.bin`
  2. In the command line enter `bower install`
  3. Enter `jeyll serve -w`
  4. Go to `localhost://4000`, click on `Bay Area` and see if it correctly maps your data
  5. If it doesn't, post an issue here
  
Resources: 
  - https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html
  - https://jonathanmh.com/tracing-preventing-http-redirects-golang/

  


## Version 1
Hello World example from: https://cloud.google.com/appengine/docs/standard/go/building-app/creating-your-application
  - Key takeaways:
    - set the CLOUDSDK_PYTHON environment variable to your python2.7 location
      - i.e. `export CLOUDSDK_PYTHON=/usr/bin/python2.7`
    - Generate a Service Account from your Google Cloud Platform project and place the JSON key in the app's root folder
    - If you've disabled your app, don't forget to enable it!
    
