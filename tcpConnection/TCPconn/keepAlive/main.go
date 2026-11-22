package keepalive

import "time"
func keepAlive(){
//to enable keepalives on a net.TCPConn object, pass true to its SetKeepAlive method
err := tcpConn.SetKeepAlive(true)
//if you need control over how often the connection sends keepAlive messages using the SetKeepAlivePeriod
err := tcpConn.SetKeepAlivePeriod(time.Minute)
}