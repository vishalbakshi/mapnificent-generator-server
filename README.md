# mapnificent-generator-server
A Go server that runs the Mapnificent Generator script to prepare GTFS binary files

## Version 2
This version should have the following functionality:

  - [ ] Retrieve zip file from transitfeeds API
  - [ ] Save the zip file to local memory
  - [ ] Generate .bin file from zip file using `mapnificent.go`
  
  
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
    
