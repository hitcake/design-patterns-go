package observer

import "fmt"

type Listener interface {
	notify(url string, status string)
}

type NativeListener struct {
	name        string
	serviceUrls []string
}

func NewNativeListener(name string) Listener {
	return &NativeListener{serviceUrls: make([]string, 0), name: name}
}

func (d *NativeListener) notify(url string, status string) {
	if status == "up" {
		d.serviceUrls = append(d.serviceUrls, url)
	} else {
		for i := 0; i < len(d.serviceUrls); i++ {
			if d.serviceUrls[i] == url {
				d.serviceUrls = append(d.serviceUrls[:i], d.serviceUrls[i+1:]...)
			}
		}
	}
	fmt.Printf("url %s %s\n", url, status)
}

type RemoteListener struct {
	name        string
	callBackUrl string
}

func NewRemoteListener(name, callBackUrl string) Listener {
	return &RemoteListener{name: name, callBackUrl: callBackUrl}
}
func (d *RemoteListener) notify(url string, status string) {
	fmt.Printf("call %s: {url:%s,status:%s}\n", d.callBackUrl, url, status)
}
