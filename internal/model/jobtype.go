package model

type JobType string

const (
	JobTypeSendEmail   JobType = "send_email"
	JobTypeSendSMS     JobType = "send_sms"
	JobTypeResizeImage JobType = "resize_image"
	JobTypeGeneratePDF JobType = "generate_pdf"
	JobTypeWebhook     JobType = "webhook"
	JobTypeDataExport  JobType = "data_export"
	JobTypeDataImport  JobType = "data_import"
)
