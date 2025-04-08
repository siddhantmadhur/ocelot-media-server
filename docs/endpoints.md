# Endpoints

[^1]: This endpoint requires admin privileges
[^2]: This endpoint's authentication requirement can be bypassed in the "wizard stage"
[^3]: This endpoint has not been implemented yet

## Authentication

Something important to note is that a lot of endpoints (such as those managing files) require elevated privileges, this means that the request has to be sent with an access token like this:
```
{
  "Authorization": "Bearer {YOUR_ACCESS_TOKEN}",
  // ...rest of the header
}
```
The only exception is when the server is in the "wizard stage", i.e. it has just been created and no admin users exist yet, and even then only a limited number of endpoints would be available.
  
- **POST `/v1/auth/user/create` [^1][^2][^3]**
  
  Create a new user. If a user is in the "wizard stage", this replaces/updates any pre-existing user account so only one account is present by the end of the "wizard stage"
  ``` 
  { 
    "username": "johndoe128",
    "password": "secret",
    "display_name": "John Doe",
  }
  ```
  Returns:
  ```
  {
    "msg": "User has been created",
    "user_id": "<users unique UUID here>",
  }
  ```
- **GET `/v1/auth/users/all` [^3]**

  Grab a list of all the users display names

  Returns:
  ```
  {
    "count": 3,
    "users": [
        {
          "display_name:" "John Doe",
        },
        // ... 
    ],
  }
  ```

