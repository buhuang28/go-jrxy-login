package main

type LoginReq struct {
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	SchoolName string `json:"school_name"`
}
