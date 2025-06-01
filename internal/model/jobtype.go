package model

type JobType string

const (
	JobTypeSendEmail   JobType = "send_email"
	JobTypeSendSMS     JobType = "send_sms"
	JobTypeResizeImage JobType = "resize_image"
)
