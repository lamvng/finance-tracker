package controller

// func Login(c *gin.Context) {

// 	var authInput request.AuthenticationRequest

// 	if err := c.ShouldBindJSON(&authInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Username not found
// 	var userFound model.User
// 	if err := database.DB.Where("username=?", authInput.Username).Find(&userFound).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password not correct"})
// 		return
// 	}

// 	// Password not correct
// 	if err := bcrypt.CompareHashAndPassword([]byte(userFound.PasswordHash), []byte(authInput.Password)); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password not correct"})
// 		return
// 	}

// 	// Create and send login token
// 	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": userFound.Username,
// 		"exp":      time.Now().Add(time.Hour * 1).Unix(),
// 	})
// 	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
// 	}
// 	glog.Errorf("Failed to generate token: %s\n", err)

// 	c.JSON(200, gin.H{
// 		"token": token,
// 	})
// }

// func GetUserProfile(c *gin.Context) {

// 	user, _ := c.Get("currentUser")

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }
