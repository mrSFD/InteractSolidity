// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;
contract Todo{
    address public owner;
    Task[] tasks;
    struct Task{
        string content;
        uint MsgID;
    }
    constructor(){
        owner = msg.sender;
    }
    modifier isOwner(){
        require(owner == msg.sender);
        _;
    }
    function add(string memory _content, uint _MsgID) public isOwner {
        tasks.push(Task(_content, _MsgID));
    }
    function get(uint _MsgID) public view returns (Task memory){
        return tasks[_MsgID];
    }
    function list() public view returns (Task[] memory){
        return tasks;
    }
}