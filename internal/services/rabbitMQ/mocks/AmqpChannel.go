// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	amqp "github.com/NeowayLabs/wabbit/amqp"
	mock "github.com/stretchr/testify/mock"

	wabbit "github.com/NeowayLabs/wabbit"
)

// AmqpChannel is an autogenerated mock type for the AmqpChannel type
type AmqpChannel struct {
	mock.Mock
}

type AmqpChannel_Expecter struct {
	mock *mock.Mock
}

func (_m *AmqpChannel) EXPECT() *AmqpChannel_Expecter {
	return &AmqpChannel_Expecter{mock: &_m.Mock}
}

// Channel provides a mock function with given fields: uri
func (_m *AmqpChannel) Channel(uri string) (wabbit.Channel, error) {
	ret := _m.Called(uri)

	var r0 wabbit.Channel
	if rf, ok := ret.Get(0).(func(string) wabbit.Channel); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wabbit.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uri)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AmqpChannel_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type AmqpChannel_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
//  - uri string
func (_e *AmqpChannel_Expecter) Channel(uri interface{}) *AmqpChannel_Channel_Call {
	return &AmqpChannel_Channel_Call{Call: _e.mock.On("Channel", uri)}
}

func (_c *AmqpChannel_Channel_Call) Run(run func(uri string)) *AmqpChannel_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AmqpChannel_Channel_Call) Return(_a0 wabbit.Channel, _a1 error) *AmqpChannel_Channel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Consume provides a mock function with given fields: queue, consumer, autoAck, exclusive, noLocal, noWait
func (_m *AmqpChannel) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool) (<-chan amqp.Delivery, error) {
	ret := _m.Called(queue, consumer, autoAck, exclusive, noLocal, noWait)

	var r0 <-chan amqp.Delivery
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool) <-chan amqp.Delivery); ok {
		r0 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp.Delivery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool, bool, bool, bool) error); ok {
		r1 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AmqpChannel_Consume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Consume'
type AmqpChannel_Consume_Call struct {
	*mock.Call
}

// Consume is a helper method to define mock.On call
//  - queue string
//  - consumer string
//  - autoAck bool
//  - exclusive bool
//  - noLocal bool
//  - noWait bool
func (_e *AmqpChannel_Expecter) Consume(queue interface{}, consumer interface{}, autoAck interface{}, exclusive interface{}, noLocal interface{}, noWait interface{}) *AmqpChannel_Consume_Call {
	return &AmqpChannel_Consume_Call{Call: _e.mock.On("Consume", queue, consumer, autoAck, exclusive, noLocal, noWait)}
}

func (_c *AmqpChannel_Consume_Call) Run(run func(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool)) *AmqpChannel_Consume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(bool), args[3].(bool), args[4].(bool), args[5].(bool))
	})
	return _c
}

func (_c *AmqpChannel_Consume_Call) Return(_a0 <-chan amqp.Delivery, _a1 error) *AmqpChannel_Consume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Publish provides a mock function with given fields: uri, queueName, exchange, exchangeType, body, reliable
func (_m *AmqpChannel) Publish(uri string, queueName string, exchange string, exchangeType string, body string, reliable bool) error {
	ret := _m.Called(uri, queueName, exchange, exchangeType, body, reliable)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, bool) error); ok {
		r0 = rf(uri, queueName, exchange, exchangeType, body, reliable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AmqpChannel_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type AmqpChannel_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//  - uri string
//  - queueName string
//  - exchange string
//  - exchangeType string
//  - body string
//  - reliable bool
func (_e *AmqpChannel_Expecter) Publish(uri interface{}, queueName interface{}, exchange interface{}, exchangeType interface{}, body interface{}, reliable interface{}) *AmqpChannel_Publish_Call {
	return &AmqpChannel_Publish_Call{Call: _e.mock.On("Publish", uri, queueName, exchange, exchangeType, body, reliable)}
}

func (_c *AmqpChannel_Publish_Call) Run(run func(uri string, queueName string, exchange string, exchangeType string, body string, reliable bool)) *AmqpChannel_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(string), args[5].(bool))
	})
	return _c
}

func (_c *AmqpChannel_Publish_Call) Return(_a0 error) *AmqpChannel_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}