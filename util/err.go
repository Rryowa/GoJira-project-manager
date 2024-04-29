package util

import "errors"

var NameRequiredError = errors.New("User Name is required")
var ProjectIDRequiredError = errors.New("Project ID is required")
var UserIDRequiredError = errors.New("User ID is required")
var CreateTaskError = errors.New("Error creating a task")
var InvalidRequestError = errors.New("Invalid request payload")
var IdRequiredError = errors.New("Id is required")
var TaskRequired = errors.New("Task is required")
