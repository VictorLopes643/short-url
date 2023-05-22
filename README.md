
# Short-URL


Short-URL is a link shortening service implemented using Golang, AWS CDK, AWS Lambda, and API Gateway. This project allows you to create shortened URLs that redirect users to the original long URLs, making it easier to share and remember links.




### Features
* Shorten long URLs into concise and easy-to-share links.
* Customizable URLs with user-defined aliases or random strings.
* Redirect users from the shortened URLs to the original long URLs seamlessly.
* Scalable and highly available architecture using AWS services.

### Technologies Used
* **Go**: The backend logic and API endpoints are implemented using the Go programming language, known for its performance and simplicity.
* **AWS CDK**: AWS Cloud Development Kit (CDK) is used for infrastructure as code to define and deploy the necessary AWS resources.
* **AWS Lambda**: AWS Lambda is a serverless computing service that executes code in response to events. In this project, it runs the backend logic for creating shortened URLs and handling redirects.
* **API Gateway**: AWS API Gateway is a fully managed service that makes it easy to create, publish, maintain, monitor, and secure APIs at any scale. It acts as the entry point for accessing the Short-URL service.
readme.so
