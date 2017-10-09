# Understanding OAuth2

## To really integrate OAuth2 into our web application it’s good to understand how it works.

###That’s the flow of OAuth2:
1. The user opens the website and clicks the login button.
2. The user gets redirected to the google login handler page. This page generates a random state string by which it will identify the user, and constructs a google login link using it. The user then gets redirected to that page.
3. The user logs in and gets back a code and the random string we gave him. He gets redirected back to our page, using a POST request to give us the code and state string.
4. We verify if it’s the same state string. If it is then we use the code to ask google for a short-lived access token. We can save the code for future use to get another token later.
5. We use the token to initiate actions regarding the user account.