package worldcomm

import (
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512 // NOTE let's adjust this later
	commRadius     = 10
	minParcelX     = -3000
	minParcelZ     = -3000
	maxParcelX     = 3000
	maxParcelZ     = 3000
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int32, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func now() (time.Time, float64) {
	now := time.Now()
	return now, float64(now.UnixNano() / int64(time.Millisecond))
}

type parcel struct {
	x int32
	z int32
}

type commarea struct {
	vMin parcel
	vMax parcel
}

func (ca *commarea) contains(p parcel) bool {
	vMin := ca.vMin
	vMax := ca.vMax
	return p.x >= vMin.x && p.z >= vMin.z && p.x <= vMax.x && p.z <= vMax.z
}

func makeCommArea(center parcel) commarea {
	return commarea{
		vMin: parcel{max(center.x-commRadius, minParcelX), max(center.z-commRadius, minParcelZ)},
		vMax: parcel{min(center.x+commRadius, maxParcelX), min(center.z+commRadius, maxParcelZ)},
	}
}

type clientPosition struct {
	quaternion [7]float32
	parcel     parcel
	lastUpdate time.Time
}

type IWebsocket interface {
	SetWriteDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetReadLimit(int64)
	SetPongHandler(h func(appData string) error)

	ReadMessage() (messageType int, p []byte, err error)

	WriteMessage(messageType int, data []byte) error
	WritePreparedMessage(pm *websocket.PreparedMessage) error

	Close() error
}

type client struct {
	conn       IWebsocket
	position   *clientPosition
	flowStatus FlowStatus
	peerId     string
	send       chan websocket.PreparedMessage

	transient struct {
		genericMessage *GenericMessage
	}
}

func makeClient(conn IWebsocket) *client {
	return &client{
		conn:       conn,
		position:   nil,
		flowStatus: FlowStatus_UNKNOWN_STATUS,
		send:       make(chan websocket.PreparedMessage, 256),
	}
}

func (c *client) close() {
	close(c.send)
}

type enqueuedMessage struct {
	client  *client
	msgType MessageType
	ts      time.Time
	bytes   []byte
}

type WorldCommunicationState struct {
	clients    map[*client]bool
	queue      chan *enqueuedMessage
	register   chan *client
	unregister chan *client

	transient struct {
		positionMessage *PositionMessage
	}
}

func MakeState() *WorldCommunicationState {
	state := &WorldCommunicationState{
		clients:    make(map[*client]bool),
		queue:      make(chan *enqueuedMessage),
		register:   make(chan *client),
		unregister: make(chan *client),
	}

	return state
}

func Close(state *WorldCommunicationState) {
	close(state.queue)
	close(state.register)
	close(state.unregister)
}

func Connect(state *WorldCommunicationState, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("socket connect error", err)
		return
	}

	log.Println("socket connect")
	c := makeClient(conn)
	state.register <- c
	go read(state, c)
	go write(state, c)
}

func read(state *WorldCommunicationState, c *client) {
	defer func() {
		state.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, bytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexcepted close error: %v", err)
			}
			log.Printf("read error: %v", err)
			break
		}

		if c.transient.genericMessage == nil {
			c.transient.genericMessage = &GenericMessage{}
		}
		generic := new(GenericMessage)
		if err := proto.Unmarshal(bytes, generic); err != nil {
			log.Println("Failed to load:", err)
			continue
		}

		ts := time.Unix(0, int64(generic.GetTime())*int64(time.Millisecond))
		// if ts.After(time.Now()) {
		// 	// TODO
		// 	continue
		// }

		message := &enqueuedMessage{client: c, msgType: generic.GetType(), ts: ts, bytes: bytes}
		state.queue <- message
	}
}

func write(state *WorldCommunicationState, c *client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WritePreparedMessage(&message); err != nil {
				log.Println(err)
				return
			}

			for i := 0; i < len(c.send); i++ {
				message, ok = <-c.send
				if !ok {
					c.conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				if err := c.conn.WritePreparedMessage(&message); err != nil {
					log.Println(err)
					return
				}
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func Process(state *WorldCommunicationState) {
	for {
		select {
		case c := <-state.register:
			register(state, c)
		case c := <-state.unregister:
			unregister(state, c)
		case enqueuedMessage := <-state.queue:
			processMessage(state, enqueuedMessage)
		}
	}
}

func register(state *WorldCommunicationState, c *client) {
	state.clients[c] = true
}

func unregister(state *WorldCommunicationState, c *client) {
	delete(state.clients, c)
	c.close()

	if c.peerId != "" {
		_, ms := now()
		msg := &ClientDisconnectedFromServerMessage{
			Type:   MessageType_CLIENT_DISCONNECTED_FROM_SERVER,
			Time:   ms,
			PeerId: c.peerId,
		}
		bytes, err := proto.Marshal(msg)
		if err != nil {
			log.Println("error sending DISCONNECTED_FROM_SERVER msg", err)
			return
		}

		broadcast(state, c, bytes)
	}
}

func processMessage(state *WorldCommunicationState, enqueuedMessage *enqueuedMessage) {
	c := enqueuedMessage.client
	msgTs := enqueuedMessage.ts
	msgType := enqueuedMessage.msgType
	bytes := enqueuedMessage.bytes

	switch msgType {
	case MessageType_FLOW_STATUS:
		message := &FlowStatusMessage{}
		if err := proto.Unmarshal(bytes, message); err != nil {
			log.Println("Failed to decode flow status message")
			return
		}

		flowStatus := message.GetFlowStatus()
		if flowStatus != FlowStatus_UNKNOWN_STATUS {
			c.flowStatus = flowStatus
		}
	case MessageType_CHAT:
		message := &ChatMessage{}
		if err := proto.Unmarshal(bytes, message); err != nil {
			log.Println("Failed to decode chat message")
			return
		}
		broadcast(state, c, bytes)
	case MessageType_POSITION:
		if state.transient.positionMessage == nil {
			state.transient.positionMessage = &PositionMessage{}
		}
		message := state.transient.positionMessage
		if err := proto.Unmarshal(bytes, message); err != nil {
			log.Println("Failed to decode position message")
			return
		}

		if c.position == nil || c.position.lastUpdate.Before(msgTs) {
			if c.position == nil {
				c.position = &clientPosition{}
			}
			c.position.quaternion = [7]float32{
				message.GetPositionX(),
				message.GetPositionY(),
				message.GetPositionZ(),
				message.GetRotationX(),
				message.GetRotationY(),
				message.GetRotationZ(),
				message.GetRotationX(),
			}
			c.position.parcel = parcel{
				x: int32(message.GetPositionX()),
				z: int32(message.GetPositionZ()),
			}
			c.position.lastUpdate = msgTs
		}
		broadcast(state, c, bytes)
	case MessageType_PROFILE:
		message := &ProfileMessage{}
		if err := proto.Unmarshal(bytes, message); err != nil {
			log.Println("Failed to decode profile message")
			return
		}

		if c.peerId == "" {
			c.peerId = message.GetPeerId()
		}

		broadcast(state, c, bytes)
	}
}

func broadcast(state *WorldCommunicationState, from *client, bytes []byte) {
	if from.position == nil {
		return
	}

	commArea := makeCommArea(from.position.parcel)

	msg, err := websocket.NewPreparedMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Println("prepare message error on broadcast", err)
		return
	}

	for c := range state.clients {
		if c == from {
			continue
		}

		if c.position == nil {
			continue
		}

		if !commArea.contains(c.position.parcel) {
			continue
		}

		c.send <- *msg
	}
}
