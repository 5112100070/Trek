package app

import (
	"log"

	"github.com/5112100070/Trek/src/entity"
	"github.com/5112100070/Trek/src/global"

	"github.com/gin-gonic/gin"
)

func SaveContact(c *gin.Context) {
	fullname := c.PostForm("fullname")
	email := c.PostForm("email")
	company := c.PostForm("company")
	phone := c.PostForm("phone")
	project := c.PostForm("project")

	if len(fullname) <= 2 || len(fullname) >= 40 {
		global.BadRequestResponse(c, "Invalid name")
		return
	}

	if len(company) <= 2 || len(company) > 40 {
		global.BadRequestResponse(c, "Invalid company name")
		return
	}

	if len(phone) <= 2 || len(phone) >= 20 {
		global.BadRequestResponse(c, "Invalid phone number")
		return
	}

	if !global.IsValidEmail(email) || len(fullname) <= 2 || len(fullname) > 35 {
		global.BadRequestResponse(c, "Invalid email")
		return
	}

	if len(project) <= 2 || len(fullname) > 40 {
		global.BadRequestResponse(c, "Invalid name")
		return
	}

	if len(project) <= 10 || len(project) > 150 {
		global.BadRequestResponse(c, "Invalid project description")
		return
	}

	// define service pub
	pubService := global.GetServicePublic()

	subs := entity.UserSubscriber{
		Fullname:           fullname,
		PhoneNumber:        phone,
		Email:              email,
		Company:            company,
		ProjectDescription: project,
	}
	err := pubService.SaveSubscriber(subs)
	if err != nil {
		log.Println("cannot save subscriber. Err", err)
		global.InternalServerErrorResponse(c, err.Error())
		return
	}

	global.OKResponse(c, nil)
}
