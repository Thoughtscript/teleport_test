package models

import "sync"

// Tracks user and password
// Most basic in-memory implementation - no encryption or TFA
var authenticationTable = make(map[string]string)

// Define mutex
var aLock = &sync.Mutex{}

// ----------------------------
// Helpers
// ----------------------------

func AuthenticationTable() map[string]string {
	aLock.Lock()
	defer aLock.Unlock()
	return authenticationTable
}

func DeleteFromAuthenticationTable(user string) {
	aLock.Lock()
	defer aLock.Unlock()
	delete(authenticationTable, user)
}

func AddToAuthenticationTable(user string, password string) {
	aLock.Lock()
	defer aLock.Unlock()
	authenticationTable[user] = password
}

func ReadFromAuthenticationTable(user string) string {
	aLock.Lock()
	defer aLock.Unlock()
	return authenticationTable[user]
}