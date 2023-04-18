# GoLang Server for InstaFeed Component
This Go server servers as a proxy for the [InstaFeed React component](https://github.com/fs1g17/instafeed), as well as providing caching capabilities. 

# Instructions
## Intended Use
This project is for displaying your 9 latest instagram posts in your website. A detailed tutorial can be found on my [medium blog](https://medium.com/dev-genius/caching-instagram-api-response-in-a-go-server-b67beaeaa1fd). 

## Prerequisites
To use this, you need to have a [Meta developer account](https://developers.facebook.com/), and have a Facebook app with  [Instagram Basic Display](https://developers.facebook.com/docs/instagram-basic-display-api/) enabled. A detailed tutorial of how to set that app can be found [here](https://docs.oceanwp.org/article/487-how-to-get-instagram-access-token). 

To use this project, you need: 
  - A registered Facebook app with Instagram Basic Display enabled and have yourself added as an Instagram tester [detailed tutorial](https://developers.facebook.com/docs/instagram-basic-display-api/)
  - Have your `access token` and `user id`
  - Have an existing google cloud project

## How to Run 
Pull the repository and add your `access token` and `user id` into the `.env` file. 

I deployed this project in google cloud. To do the same, you need to follow the [guide](https://firebase.google.com/docs/hosting/cloud-run).  In short, the steps are the following: 
  - Install `gcloud` CLI tools 
  - Initialize your project with `gcloud init` and point to your google cloud project
  - Run `gcloud builds submit --tag gcr.io/PROJECT_NAME/CONTAINER_NAME`, replacing `PROJECT_NAME` with your google cloud project name and `CONTAINER_NAME` with whatever you would like to call this container
  - Run `gcloud run deploy --image gcr.io/PROJECT_NAME/CONTAINER_NAME` 
  
And that's it! You can now add the [InstaFeed React component](https://github.com/fs1g17/instafeed) to your React project and it will display your latest Instagram posts. 

## Demo 
This is how the [InstaFeed React component](https://github.com/fs1g17/instafeed) displays the response from the published project:

![image](https://user-images.githubusercontent.com/47851444/232896708-13e222b2-8b70-4517-a803-25ee8562297c.png)
