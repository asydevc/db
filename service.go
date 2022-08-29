// author: asydevc <asydev@163.com>
// date: 2021-02-13

package db

import (
	"xorm.io/xorm"
)

// Struct for anonymous in service.
//
//	type ExampleService struct{
//	    xdb.Service
//	}
//
//	func NewExampleService(s ...*xorm.Session) *ExampleService {
//	    o := &ExampleService{}
//	    o.Use(s...)
//	    return o
//	}
type Service struct {
	sess *xorm.Session
}

// Read master connection.
func (o *Service) Master() *xorm.Session {
	if o.sess == nil {
		return Master()
	}
	return o.sess
}

// Return slave connection.
func (o *Service) Slave() *xorm.Session {
	if o.sess == nil {
		return Slave()
	}
	return o.sess
}

// Use specified connection.
func (o *Service) Use(s ...*xorm.Session) {
	if s != nil && len(s) > 0 {
		o.sess = s[0]
	}
}
