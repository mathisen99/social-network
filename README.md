
# Anglular / Golang Social-network boilerplate
boilerplate code to setup the basics for our project.



## Features

- angular is setup and ready to use
- go webserver (With decent settings)
- angular proxy for development


## Deployment

To deploy this project run

```bash
  git clone https://github.com/mathisen99/social-network.git
  cd social-network
  go run . (this will start the go webserver)

  (New terminal or tab)
  cd social-network/social-network (yes it should be 2 social-network)
  ng serve (this will start the angular frontend and use our golang proxy)
```
After that both should be running and you can reach the project at localhost:8080

if you want to use build version of the project and not work in "developer mode"
make sure you are in the correct folder " social-network/social-network " and then..


```bash
ng build --configuration production (this will build it and put the project into the dist folder)
cp -R dist/social-network ../build-version-0.1

modify main.go to start the realserver instead of the dev server
modify server.go inside backend folder and change the angluar app path to whatever you named your build (build-version-0.1 we use in this example)
```

## Author

- [@Mathisen](https://www.github.com/mathisen99)