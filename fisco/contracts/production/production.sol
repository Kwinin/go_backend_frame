pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;
contract production {
  event ItemSet(string pid, int _opType, int64 opId, string value);

  string public version;
  string public pid;
  string public pname;
  string public principal;
  struct Item {
    string opName;
    int8 opType;
    int64 opId;
    string itemValue;
    bool isVaild;
    string principal;
  }
  struct OpArr{
    int8[] opType;
    int64[] opId;
  }

  mapping (string => mapping (int8 => mapping (int64 => Item))) public itemMap;
  mapping (string => OpArr) opArrMap;

  constructor(string memory _pid, string memory _v, string memory _lot, string memory _pname, string memory _principal) public {
    version = _v;
    pid = _pid;
    pname = _pname;
    principal =_principal;
  }

  function setItem(string memory _pid, int8 _opType, int64 _opId, string memory _opName, string memory _value, string memory principal) external {
    require(!itemMap[_pid][_opType][_opId].isVaild, "cann't set itemMap again");
    opArrMap[_pid].opType.push(_opType);
    opArrMap[_pid].opId.push(_opId);
    itemMap[_pid][_opType][_opId].isVaild = true;
    itemMap[_pid][_opType][_opId].opName = _opName;
    itemMap[_pid][_opType][_opId].opType = _opType;
    itemMap[_pid][_opType][_opId].opId = _opId;
    itemMap[_pid][_opType][_opId].itemValue = _value;
    itemMap[_pid][_opType][_opId].principal = principal;
    emit ItemSet(_pid, _opType, _opId, _value);
  }

  function getOpTypes(string memory _pid) public view returns (int8[] memory opTypes){
    int8[] memory opTypes = opArrMap[_pid].opType;
    return opTypes;
  }

  function getOpIds(string memory _pid) public view returns (int64[] memory opIds){
    int64[] memory opIds = opArrMap[_pid].opId;
    return opIds;
  }

  function getItems(string memory _pid) public view returns (Item[] memory items){
    int64[] memory opIds = opArrMap[_pid].opId;
    int8[] memory opTypes = opArrMap[_pid].opType;
    Item[] memory items = new Item[](opIds.length);
    for (uint i = 0; i < opIds.length; i++) {
      int64 item_opId = opIds[i];
      int8 item_opType = opTypes[i];
      Item  memory item =  itemMap[_pid][item_opType][item_opId];

      items[i] = item;

    }
    return items;
  }

  function getItem(string memory _pid, int8 _opType, int64 _opId ) public view returns (string memory opName,string memory itemValue, bool isVaild, string memory principal){
    Item  memory item =  itemMap[_pid][_opType][_opId];
    return (item.opName, item.itemValue, item.isVaild, item.principal);
  }
}
