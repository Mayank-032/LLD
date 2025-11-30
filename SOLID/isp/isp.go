package isp

// ISP : Interface Segregation Principle

// ISP - Violated
type IMessageDelivery interface {
	DeliverEmail(message string, recipient string) error
	DeliverSMS(message string, phoneNumber string) error
	DeliverPushNotification(message string, deviceID string) error
}

type MessageDeliveryViolation struct{}
func (mdv *MessageDeliveryViolation) DeliverEmail(message string, recipient string) error {
	// deliver email
	return nil
}

func (mdv *MessageDeliveryViolation) DeliverSMS(message string, phoneNumber string) error {
	// deliver SMS
	return nil
}

func (mdv *MessageDeliveryViolation) DeliverPushNotification(message string, deviceID string) error {
	// deliver push notification
	return nil
}

// ISP - Followed
type IEmailDelivery interface {
	DeliverEmail(message string, recipient string) error
}

type ISMSDelivery interface {
	DeliverSMS(message string, phoneNumber string) error
}

type IPushNotificationDelivery interface {
	DeliverPushNotification(message string, deviceID string) error
}

type EmailDelivery struct{}
func (ed *EmailDelivery) DeliverEmail(message string, recipient string) error {
	// deliver email
	return nil
}

type SMSDelivery struct{}
func (sd *SMSDelivery) DeliverSMS(message string, phoneNumber string) error {
	// deliver SMS
	return nil
}

type PushNotificationDelivery struct{}
func (pnd *PushNotificationDelivery) DeliverPushNotification(message string, deviceID string) error {
	// deliver push notification
	return nil
}

type IMessageDeliveryInterface interface {
	IEmailDelivery
	ISMSDelivery
	IPushNotificationDelivery
}