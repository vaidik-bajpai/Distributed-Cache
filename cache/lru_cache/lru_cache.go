package cache

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
)

var (
	retainSnapshotCount = 2
	raftTimeout         = 10 * time.Second
)

var (
	ErrSetCalledOnCandidate = errors.New("error this node is not the leader node")
	ErrEncodingPutCmd       = errors.New("error could not encode the put command")
)

type LRUCache struct {
	sync.RWMutex
	capacity uint64
	store    map[string]*list.Element
	list     *list.List //stores only values
	raft     *raft.Raft
	inmem    bool
	RaftDir  string
	RaftBind string
}

func NewLRUCache(capacity uint64) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		store:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

type Pair struct {
	key, value []byte
}

type command struct {
	Op    string `json:"op"`
	Key   []byte `json:"key"`
	Value []byte `json:"val"`
}

func (c *LRUCache) Open(enableSingle bool, localID string) error {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)

	addr, err := net.ResolveTCPAddr("tcp", c.RaftBind)
	if err != nil {
		return err
	}

	transport, err := raft.NewTCPTransport(c.RaftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	snapshots, err := raft.NewFileSnapshotStore(c.RaftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}

	var logStore raft.LogStore
	var stableStore raft.StableStore

	if c.inmem {
		logStore = raft.NewInmemStore()
		stableStore = raft.NewInmemStore()
	} else {
		boltDB, err := raftboltdb.New(raftboltdb.Options{
			Path: filepath.Join(c.RaftDir, fmt.Sprintf("rafts-%d.db", rand.Intn(20))),
		})
		if err != nil {

			return fmt.Errorf("err bolt store: %s", err)
		}

		logStore = boltDB
		stableStore = boltDB
	}

	ra, err := raft.NewRaft(config, c, logStore, stableStore, snapshots, transport)
	if err != nil {
		return fmt.Errorf("new raft error: %s", err)
	}
	c.raft = ra

	if enableSingle {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	}

	return nil
}

func (c *LRUCache) Join(nodeID, addr string) error {
	log.Printf("received join request for remote node %s at %s", nodeID, addr)

	configFuture := c.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		log.Printf("failed to get raft configuration: %v", err)
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, then nothing -- not even
			// a join operation -- is needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(nodeID) {
				log.Printf("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return nil
			}

			future := c.raft.RemoveServer(srv.ID, 0, 0)
			if err := future.Error(); err != nil {
				return fmt.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			}
		}
	}

	f := c.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if f.Error() != nil {
		return f.Error()
	}
	log.Printf("node %s at %s joined successfully", nodeID, addr)
	return nil
}

func (c *LRUCache) Get(key []byte) ([]byte, error) {
	c.PrintStore(key)

	c.RLock()
	defer c.RUnlock()
	log.Printf("key %s inside cache get\n", string(key))
	if elem, found := c.store[string(key)]; found {
		log.Println("cache hit")
		c.list.MoveToFront(elem)
		log.Println(string(elem.Value.(*Pair).value))
		return elem.Value.(*Pair).value, nil
	}
	return nil, nil
}

func (c *LRUCache) Set(key, value []byte) error {
	if c.raft.State() != raft.Leader {
		return ErrSetCalledOnCandidate
	}

	command := &command{Op: "PUT", Key: key, Value: value}
	data, err := json.Marshal(command)
	if err != nil {
		return ErrEncodingPutCmd
	}

	f := c.raft.Apply(data, raftTimeout)
	return f.Error()
	/* log.Println("Set called")
	// Lock to ensure the cache is thread-safe
	c.Lock()
	defer c.Unlock()

	// If the element is already present in the cache, move it to the front
	if elem, found := c.store[string(key)]; found {
		elem.Value.(*Pair).value = value
		c.list.MoveToFront(elem)
		return
	}

	// If the element is not present, check if the cache has reached its capacity
	if len(c.store) >= int(c.capacity) {
		lruElem := c.list.Back()

		if lruElem != nil {
			// Remove the least recently used element
			delete(c.store, string(lruElem.Value.(*Pair).key))
			c.list.Remove(lruElem)
		}
	}

	// Add the new element to the front of the list and store it
	pair := &Pair{key: key, value: value}
	elem := c.list.PushFront(pair)
	c.store[string(key)] = elem

	log.Println("store after insertion:", c.store) */
}

func (c *LRUCache) Apply(l *raft.Log) interface{} {
	log.Printf("Applying log entry: index=%d, term=%d, command=%s\n", l.Index, l.Term, string(l.Data))
	var cmd command
	if err := json.Unmarshal(l.Data, &cmd); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}

	switch cmd.Op {
	case "PUT":
		return c.applySet(&cmd)
	default:
		panic(fmt.Sprintf("unrecognisable command op: %s", cmd.Op))
	}
}

func (c *LRUCache) applySet(cmd *command) interface{} {
	// Lock to ensure the cache is thread-safe
	c.Lock()
	defer c.Unlock()

	// If the element is already present in the cache, move it to the front
	if elem, found := c.store[string(cmd.Key)]; found {
		elem.Value.(*Pair).value = cmd.Value
		c.list.MoveToFront(elem)
		return nil
	}

	// If the element is not present, check if the cache has reached its capacity
	if len(c.store) >= int(c.capacity) {
		lruElem := c.list.Back()

		if lruElem != nil {
			// Remove the least recently used element
			delete(c.store, string(lruElem.Value.(*Pair).key))
			c.list.Remove(lruElem)
		}
	}

	// Add the new element to the front of the list and store it
	pair := &Pair{key: cmd.Key, value: cmd.Value}
	elem := c.list.PushFront(pair)
	c.store[string(cmd.Key)] = elem

	return nil
}

// Restore implements raft.FSM.
func (c *LRUCache) Restore(snapshot io.ReadCloser) error {
	// Read all data from the snapshot.
	allData, err := io.ReadAll(snapshot)
	if err != nil {
		return err
	}

	// Split the data into parts (list data).
	dataParts := strings.Split(string(allData), "\n")
	if len(dataParts) != 1 {
		return fmt.Errorf("invalid snapshot data: expected one part (list), but got %d", len(dataParts))
	}

	// Deserialize the list data (list of Pair).
	var listData []Pair
	if err := json.Unmarshal([]byte(dataParts[0]), &listData); err != nil {
		return fmt.Errorf("failed to unmarshal list data: %v", err)
	}

	// Clear the existing store and list.
	c.store = make(map[string]*list.Element)
	c.list = list.New()

	// Rebuild the list with the restored elements (Pair type).
	for _, pair := range listData {
		element := c.list.PushBack(pair)
		// Restore the relationship between the map and the list.
		c.store[string(pair.key)] = element
	}

	return nil
}

// Snapshot implements raft.FSM.
// Snapshot implements raft.FSM.
func (c *LRUCache) Snapshot() (raft.FSMSnapshot, error) {
	c.Lock() // Lock to ensure thread safety when accessing cache state.
	defer c.Unlock()

	// Clone the list to ensure we're not modifying the original list while taking the snapshot.
	var listSnapshot []Pair
	for ele := c.list.Front(); ele != nil; ele = ele.Next() {
		pair := ele.Value.(Pair)
		listSnapshot = append(listSnapshot, pair)
	}

	// Return a snapshot with only the list data (not the full store).
	return &fsmSnapshot{list: listSnapshot}, nil
}

type fsmSnapshot struct {
	list []Pair
}

// Persist implements raft.FSMSnapshot.

func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Serialize the list data (list of Pair).
		listDataSerialized, err := json.Marshal(f.list)
		if err != nil {
			return err
		}

		// Write the serialized list data to the sink.
		if _, err := sink.Write(listDataSerialized); err != nil {
			return err
		}

		// Close the sink after writing all data.
		return sink.Close()
	}()

	if err != nil {
		// If an error occurs, cancel the sink to ensure the snapshot is not persisted.
		sink.Cancel()
	}

	return err
}

// Release implements raft.FSMSnapshot.
func (f *fsmSnapshot) Release() {}

func (c *LRUCache) PrintStore(key []byte) {
	for k, v := range c.store {
		log.Printf("a%qaa%qa", k, string(key))
		if k == string(key) {
			log.Println("key found")
		}
		fmt.Printf("key: %s\tvalue: %s\n", k, string(v.Value.(*Pair).value))
	}
}
