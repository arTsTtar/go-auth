# User Creation and Management app (BACK-END)

 Auth Features:
 - Create Clients (regular password + 6 one-time passwords + **(optional)** QR code and Secret for 2FA) ```POST /api/register```
 - Login via Email + Regular Password ```POST api/login```
 - Login via Email + One-time Password ```POST api/altLogin```
 - Logout ```POST api/logout```

 User Features:
 - Get Single user details by request cookie jwt token data ```GET /api/user```
 - Change regular password (user has to be logged in) ```POST api/user/changePassword```
 - Reset To random password (can call target user or any user with role ROLE_ADMIN) ```POST /api/user/{user_id}/resetPassword```

 Admin Endpoints:
 - Get all users ```GET /api/admin/users```


 Technologies Golang:
 - Golang v1.16
 - GORM v1.21.16
 - Fiber v2.20.2
 - jwt-go v3.2.0

 Technologies DB:
 - MariaDB 10.6.4
