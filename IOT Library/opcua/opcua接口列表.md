- 获取全部设备([]\*OpcuaDevice)
  ```go
  type OpcuaDevice struct {
    models.Base
    Name       string
    SN         string
    Endpoint   string
    RootNodeID string
  }
  ```
- 获取全部节点 ID([]OpcuaNode)

  ```go
  type OpcuaNode struct {
      NodeID      string
      DataType    string
      Description string
      Unit        string
      Scale       string
      BrowseName  string
  }
  ```

- 获取节点 ID 的原始值(string)
- 获取节点 ID 的解析值(string)
- 获取节点 ID 值(value \*OpcuaValue)

  ```go
  type OpcuaValue struct {
      Time        time.Time
      NodeID      string
      Raw         string
      Parse       string
      Description string
      }
  ```

- 获取设备 ID 原始值([]\*OpcuaBase)

  ```go

  type OpcuaBase struct {
      models.ValueBase
      NodeID      string
      Raw         string
      Parse       string
      Description string
  }
  ```

- 获取设备 ID 解析后的值

  ```go
  type OpcuaBase struct {
      models.ValueBase
      NodeID      string
      Raw         string
      Parse       string
      Description string
  }
  ```
- 获取设备ID值([]*OpcuaValue)

    ```go
    type OpcuaValue struct {
        Time        time.Time
        NodeID      string
        Raw         string
        Parse       string
        Description string
    }
    ```
- AddNodeIDCronJob(deviceID, nodeID, spec string)
- AddNodeIDReadCronJob(opcuaClinet *gopcua.Client, userDeviceAttributeID, deviceTableName, spec string)
- DeleteNodeIDReadCronJob(opcuaClinet *gopcua.Client, nodeID string) (err error)
- UpdateNodeIDReadCronJob(opcuaClinet *gopcua.Client, nodeID string) (err error)
- AddNodeIDSubscription(opcuaClinet *gopcua.Client, nodeID string) (err error)
- DeleteNodeIDSubscription(opcuaClinet *gopcua.Client, nodeID string) (err error)
- NewOPCUAClient(protocol OPCUADeviceProtocol, timeout time.Duration) (*gopcua.Client, error)
- DeleteOPCUAClient(opcuaClinet *gopcua.Client) (err error)
- BrowseOPCUAClient(opcuaClinet *gopcua.Client, opcuaAddr, nodeID string) (nodeList []*ResultNodeList, err error) 
    ```go
    type ResultNodeList struct {
        Name        string
        Type        string
        Addr        string
        Unit        string
        Scale       string
        Min         string
        Max         string
        Writable    string
        Description string
    }
    ```