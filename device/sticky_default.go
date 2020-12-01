// +build !linux android

package device

import (
	"github.com/gaoxiaowei/wireguard-go/conn"
	"github.com/gaoxiaowei/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
