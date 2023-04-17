### Running Locally
1. Install docker
2. Then run the script start-localhost

### Sending chat messages
As of writing only chat messages are supported. They are sent in the format of 
`{"service":"chat","action":"send","data":"text":{"text":"\"Hello\""}}`
and received on the client side
`{"service":"chat","action":"received","data":{"text":"\"Hello\"","connectionId":"475955b2-8aba-49c9-b957-58f719aefbc9"}}`