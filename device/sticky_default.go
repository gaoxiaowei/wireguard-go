// +build !linux android

package device

import (
	"github.com/gaoxiaowei/wireguard/conn"
	"github.com/gaoxiaowei/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
