/*
* Copyright 2020-present Arpabet, Inc. All rights reserved.
 */


package app

type ApplicationContext struct {
}

var applicationContext ApplicationContext

func CreateApplicationContext() {

}

func GetApplicationContext() *ApplicationContext {
	return &applicationContext
}

