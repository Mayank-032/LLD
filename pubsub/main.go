package main

import "sync"

type Broker struct {
	Id          int
	Publishers  map[int]*Publisher
	Topics      map[int]*Topic
	Subscribers map[int]*Subscriber
	mu          sync.RWMutex
}

// --------------------------------------- Publisher ---------------------------------------

type Publisher struct {
	Id   int
	Name string
}

type PublisherService struct {
	broker             *Broker
	topicService       *TopicService
	subcriptionService *SubscriptionService
	deliveryService    IMessageDeliveryInterface
}

func NewPublisherService(
	broker *Broker,
	topicService *TopicService,
	subcriptionService *SubscriptionService,
	deliveryService IMessageDeliveryInterface,
) *PublisherService {
	return &PublisherService{
		broker:             broker,
		topicService:       topicService,
		subcriptionService: subcriptionService,
		deliveryService:    deliveryService,
	}
}

func (ps *PublisherService) Publish(message *Message, topicId int) error {
	return nil
}

// --------------------------------------- Topic ---------------------------------------

type Topic struct {
	Id          int
	Name        string
	CreatedAt   int64
	Subscribers map[int]bool
	mu          sync.RWMutex
}

type TopicService struct {
	broker *Broker
}

func NewTopicService(broker *Broker) *TopicService {
	return &TopicService{broker: broker}
}

type CreateTopicReq struct {
	Name string
}

func (ts *TopicService) Create(req CreateTopicReq) (*Topic, error) {
	return nil, nil
}

func (ts *TopicService) Delete(id int) error {
	return nil
}

func (ts *TopicService) GetAll() ([]*Topic, error) {
	return nil, nil
}

func (ts *TopicService) GetById(id int) (*Topic, error) {
	return nil, nil
}

func (ts *TopicService) GetByName(name string) (*Topic, error) {
	return nil, nil
}

// --------------------------------------- Subscriber ---------------------------------------

type Subscriber struct {
	Id      int
	Name    string
	MsgChan chan *Message
}

type SubscriberService struct {
	broker *Broker
}

func NewSubscriberService(broker *Broker) *SubscriberService {
	return &SubscriberService{broker: broker}
}

type CreateSubscriberReq struct {
	Name string
}

func (ss *SubscriberService) Create(req CreateSubscriberReq) (*Subscriber, error) {
	return nil, nil
}

func (ss *SubscriberService) Delete(id int) error {
	return nil
}

func (ss *SubscriberService) GetAll() ([]*Subscriber, error) {
	return nil, nil
}

func (ss *SubscriberService) GetById(id int) (*Subscriber, error) {
	return nil, nil
}

// --------------------------------------- Subscription Service ---------------------------------------

type SubscriptionService struct {
	broker *Broker
}

func NewSubscriptionService(broker *Broker) *SubscriptionService {
	return &SubscriptionService{
		broker: broker,
	}
}

func (ss *SubscriptionService) Subscribe(subscriberId, topicId int) error {
	return nil
}

func (ss *SubscriptionService) Unsubscribe(subscriberId, topicId int) error {
	return nil
}

func (ss *SubscriptionService) IsSubscribed(subscriberId, topicId int) (bool, error) {
	return false, nil
}

func (ss *SubscriptionService) GetAllSubscribersForTopic(topicId int) ([]*Subscriber, error) {
	return nil, nil
}

func (ss *SubscriptionService) GetAllTopicsForSubscriber(subscriberId int) ([]*Topic, error) {
	return nil, nil
}

// --------------------------------------- Message ---------------------------------------

type Message struct {
	Id          int
	Payload     string // not keeping it interface, thinking to keep it as string (encoded value and it gets decoded on subscriber's end)
	PublishedAt int64
}

type IMessageDeliveryInterface interface {
	DeliverMessage(topic *Topic, message *Message, subscribers map[int]*Subscriber) error
}

type PushMessageDeliveryService struct {
}

func NewPushMessageDeliveryService() IMessageDeliveryInterface {
	return &PushMessageDeliveryService{}
}

func (pms *PushMessageDeliveryService) DeliverMessage(topic *Topic, message *Message, subscribers map[int]*Subscriber) error {
	return nil
}
