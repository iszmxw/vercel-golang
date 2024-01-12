package wsServer

import (
	"errors"
	"goapi/pkg/logger"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 声明并发安全的ws

type WsConn struct {
	*websocket.Conn              // 基础连接
	Mux             sync.RWMutex // 锁
	IsClose         bool         // 是否关闭
}

// 定义 WebSocket 服务器结构体

type WebSocketServer struct {
	upgrader    websocket.Upgrader
	connections []*WsConn
	lock        sync.RWMutex
}

// 定义将连接添加到连接列表中的方法

func (s *WebSocketServer) AddConnection(conn *WsConn) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.connections = append(s.connections, conn)
}

// 定义将连接从连接列表中删除的方法

func (s *WebSocketServer) RemoveConnection(conn *WsConn) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for i, c := range s.connections {
		if c == conn {
			s.connections = append(s.connections[:i], s.connections[i+1:]...)
			s.WsClose(conn)
			break
		}
	}
}

// 定义广播消息给所有连接的客户端的方法

func (s *WebSocketServer) BroadcastMessage(messageType int, message []byte) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, conn := range s.connections {
		err := conn.WriteMessage(messageType, message)
		if err != nil {
			logger.Info("广播消息失败")
			logger.Error(err)
			continue
		}
	}
}

func Init() *WebSocketServer {
	// 创建 WebSocket 服务器结构体
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			EnableCompression: false, // 苹果某些旧版本的ios不兼容该压缩功能，所以苹果机型直接禁用掉就好了，或者找出相关兼容的机型进行白名单
			CheckOrigin: func(r *http.Request) bool { // CheckOrigin防止跨站点的请求伪造
				return true
			},
			//ReadBufferSize:  1024,
			//WriteBufferSize: 1024,
		},
		connections: []*WsConn{},
		lock:        sync.RWMutex{},
	}
}

func (s *WebSocketServer) Upgrade(c *gin.Context) (*WsConn, error) {
	if c.IsWebsocket() {
		// 升级get请求为webSocket协议
		conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return &WsConn{
			conn,
			sync.RWMutex{},
			false,
		}, nil
	} else {
		return nil, errors.New("不是socket請求")
	}
}

// 关闭socket资源

func (s *WebSocketServer) WsClose(ws *WsConn) {
	if ws.IsClose {
		logger.Info("当前socket已关闭")
	} else {
		ws.Mux.Lock()
		closeErr := ws.Conn.Close()
		if closeErr != nil {
			logger.Error(closeErr)
		} else {
			ws.IsClose = true
		}
		ws.Mux.Unlock()
		logger.Info("关闭成功")
	}
}
