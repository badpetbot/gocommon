package net

import (

  // Import builtin packages.
  "sync"
  "time"

  // Import 3rd party packages.
  "github.com/globalsign/mgo"
  "github.com/rs/zerolog/log"
)

// MgoConfig defines the configuration of a MongoDB client.
type MgoConfig struct {
  ClientName string   `json:"client_name"`
  URLs       []string `json:"urls"`
  Replset    string   `json:"replset"`
  Username   string   `json:"username"`
  Password   string   `json:"password"`
  AuthDB     string   `json:"auth_db"`
  Timeout    int      `json:"timeout"`
}

var mgoSessions map[string]*mgo.Session = map[string]*mgo.Session{}
var mgoSessionsMu sync.Mutex

// MgoConnect creates a MongoDB client and connects it to the server specified in the given config,
// and then fires all registered connect callbacks.
func MgoConnect(config MgoConfig) (*mgo.Session, error) {

  // Declare the options and connect.
  options := &mgo.DialInfo{
    ReplicaSetName: config.Replset,
    Addrs:          config.URLs,
    ReadPreference: &mgo.ReadPreference{
      Mode: mgo.SecondaryPreferred,
    },
    Timeout:  time.Duration(config.Timeout) * time.Second,
    Username: config.Username,
    Password: config.Password,
    Source:   config.AuthDB,
  }
  session, err := mgo.DialWithInfo(options)

  // If an error occurred, log and return a nil MgoDriver.
  if err != nil {
    return nil, err
  }

  log.Info().Msgf("Connected to database. Client name: %q", config.ClientName)

  // Check for errors, but impose no other constraints.
  session.SetSafe(&mgo.Safe{
    W:        1,
    WTimeout: 10000,
  })

  // Put the driver in the public map.
  mgoSessionsMu.Lock()
  mgoSessions[config.ClientName] = session
  mgoSessionsMu.Unlock()

  // Return the driver object.
  return session, nil
}

// MgoGetSession gets the Mgo Session keyed by the specified name. Panics if the session hasn't
// been created as of this call.
func MgoGetSession(name string) *mgo.Session {

  mgoSessionsMu.Lock()
  session, ok := mgoSessions[name]
  mgoSessionsMu.Unlock()
  if !ok {
    panic("unitialized mgo session: " + name)
  }

  return session
}

// MgoCol is a shorthand function for `MgoGetSession(client).DB(database).C(collection)`.
func MgoCol(client, database, collection string) *mgo.Collection {

  return MgoGetSession(client).DB(database).C(collection)
}

// MgoMakeEmbedded embeds all keys in the given map within the path given. For example:
/*
  embedded := MgoMakeEmbedded(map[string]interface{}{"test": 1}, "embedded.struct")
  // embedded: map[string]interface{}{
  //   "embedded.struct.test": 1
  // }
*/
func MgoMakeEmbedded(in map[string]interface{}, path string) map[string]interface{} {

  out := make(map[string]interface{}, len(in))
  for k, v := range in {
    out[path+"."+k] = v
  }
  return out
}
