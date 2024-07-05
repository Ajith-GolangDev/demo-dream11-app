package controllers

import (
	"dream_11/database"
	"dream_11/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)

	var userDetails models.User
	if err := database.DB.Debug().Where("email=?", user.Email).Find(&userDetails); err != nil {
		if err.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error ": err.Error})
			return
		}
	}

	fmt.Println("user :", userDetails)

	database.DB.Create(&models.Wallet{
		UserID:  userDetails.ID,
		Balance: 0,
	})

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDetails models.User
	if err := database.DB.Find(&userDetails).Where("email=?", user.Email); err != nil {
		if err.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error ": err.Error})
			return
		}
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

//load money to the user wallet
func LoadMoney(c *gin.Context) {
	var wallet models.Wallet
	userID := c.Param("user_id")
	amountStr := c.Param("amount")

	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	amount, _ := strconv.ParseFloat(amountStr, 64)

	wallet.Balance += amount

	database.DB.Save(&wallet)
	c.JSON(http.StatusOK, gin.H{"message": "Money loaded successfully"})
}

//join a contest by contest-id
func JoinContest(c *gin.Context) {
	var wallet models.Wallet
	var contest models.Contest
	userID := c.Param("user_id")
	contestID := c.Param("contest_id")

	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}

	if err := database.DB.Where("id = ?", contestID).First(&contest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contest not found"})
		return
	}

	if wallet.Balance < contest.EntryFee {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	wallet.Balance -= contest.EntryFee
	database.DB.Save(&wallet)

	c.JSON(http.StatusOK, gin.H{"message": "Joined contest successfully"})
}

// Create a new player
func CreateTeam(c *gin.Context) {
	var userTeam models.UserTeam

	if err := c.ShouldBindJSON(&userTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerIDs := strings.Split(userTeam.PlayerIDs, ",")
	var totalCreditScore float64
	teamCount := make(map[string]int)

	var DuplicateCheck bool
	var playerId string

	for index, id := range playerIDs {
		if index == 0 {
			playerId = id
		} else {
			if playerId == id {
				DuplicateCheck = true
			}
		}
		var player models.Player
		if err := database.DB.Where("id = ?", id).First(&player).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
			return
		}

		totalCreditScore += player.CreditScore
		teamCount[player.Team]++
	}

	if DuplicateCheck {
		c.JSON(http.StatusBadRequest, gin.H{"error": `Can't select same player`})
		return
	}

	if totalCreditScore > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Total credit score exceeds 100"})
		return
	}

	for _, count := range teamCount {
		fmt.Println("count :", count)
		if count > 7 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot select more than 7 players from a single team"})
			return
		}
	}

	database.DB.Create(&userTeam)
	c.JSON(http.StatusOK, gin.H{"message": "Team created successfully"})
}

// Create a new player
func CreatePlayer(c *gin.Context) {
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
}

// Create a new contest
func CreateContest(c *gin.Context) {
	var contest models.Contest
	if err := c.ShouldBindJSON(&contest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&contest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contest)
}

//View team by team-id
func ViewTeam(c *gin.Context) {
	var userTeam models.UserTeam
	teamID := c.Param("team_id")

	if err := database.DB.Find(&userTeam).Where("id=?", teamID); err != nil {
		if err.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error ": err.Error})
			return
		}
	}

	playerIDs := strings.Split(userTeam.PlayerIDs, ",")
	var players []models.Player
	if err := database.DB.Where("id IN ?", playerIDs).Find(&players).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	c.JSON(http.StatusOK, players)
}
