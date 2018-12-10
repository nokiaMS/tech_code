pragma solidity ^0.4.24;
//pragma experimental ABIEncoderV2;

//wandapp base.
import "./SafeMath.sol";
import "./proofVerify.sol";

//interface for communicating with erc20 token.
contract  DemoERC20 {
    function  approve(address  _spender,  uint  _value)  public  returns  (bool  success);
    function  transfer(address  _to,  uint  _value)  public   returns  (bool  success);
    function  transferFrom(address  _from,  address  _to,  uint  _value)  public returns  (bool  success);
}

contract wandappBase {
    struct Proof {
        address from_;
        address to;
        address operator;
        uint256 totalTransferCnt;
        uint256 totalWithdrawCnt;
        uint64 nonce;
        int64 timestamp_;
        //bytes32 hashX;
        uint256 chlID;
        uint256 deltaCnt;
        uint256 targetTotalCnt;
        // bytes signUser;
        // bytes signOper;
        // bytes32 x;
    }
    //Channel status.
    enum ChlState {
        NORMAL,             //In the NORMAL state when channel is created. 
        CREATOREXITING,     //creator exiting.
        OPERATOREXITING    //operator exiting.
    }

    
    //variable
    address tokenAddr;                      //token address.
    address operatorAddr;                   //operator address.
    uint256 exitWinDuration;                //duration of the exit window.
    uint256 chlNum;
    
    // Constructor
    constructor(address tokenOrigAddr, address operAddr, uint256 duration) public
    {
        tokenAddr = tokenOrigAddr;
        operatorAddr = operAddr;
        exitWinDuration = duration;
    }
    
    modifier onlyBy(address _account) {
        require( msg.sender == _account, "Sender not authorized." );
        _;
    }
    
    modifier notNegative(uint256 count) {
        require( count >= 0, "Number is negative." );
        _;
    }
    

    
    //return channel id.
    function getChlID() public returns(uint256) {
        chlNum = SafeMath.add(chlNum,1);
        return chlNum;
    }
    
    function get_from(bytes data) internal pure returns(address _from){
        assembly{
            _from := and(mload(add(data, 20)),0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
        }
    }
    
    function get_to(bytes data) internal pure returns(address to){
        assembly{
            to := and(mload(add(data, 40)),0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
        }
    }
    function get_operator(bytes data) internal pure returns(address operator){
        assembly{
            operator := and(mload(add(data, 60)),0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
        }
    }
    function get_totalTransferCnt(bytes data) internal pure returns(uint256 totalTransferCnt){
        assembly{
            totalTransferCnt := mload(add(data, 92))
        }
    }
    function get_totalWithdrawCnt(bytes data) internal pure returns(uint256 totalWithdrawCnt){
        assembly{
            totalWithdrawCnt := mload(add(data, 124))
        }
    }
    function get_nonce(bytes data) internal pure returns(uint64 nonce){
        assembly{
            nonce := and(mload(add(data, 132)),0xFFFFFFFFFFFFFFFF)
        }
    }
    function get_timestamp(bytes data) internal pure returns(int64 timestamp_){
        assembly{
            timestamp_ := and(mload(add(data, 140)),0xFFFFFFFFFFFFFFFF)
        }
    }
    function get_hashX(bytes data) internal pure returns(bytes32 hashX){
        assembly{
            hashX := mload(add(data, 172))
        }
    }
    function get_chlID(bytes data) internal pure returns(uint256 chlID){
        assembly{
            chlID := mload(add(data, 204))
        }
    }
    function get_deltaCnt(bytes data) internal pure returns(uint256 deltaCnt){
        assembly{
            deltaCnt := mload(add(data, 236))
        }
    }
    function get_targetTotalCnt(bytes data) internal pure returns(uint256 targetTotalCnt){
        assembly{
            targetTotalCnt := mload(add(data, 268))
        }
    }
    function get_signUser(bytes data) internal pure returns(bytes signUser){
        signUser = new bytes(65);
        assembly{
            mstore(add(signUser,32),mload(add(data, 300)))
            mstore(add(signUser,64),mload(add(data, 332)))
            mstore(add(signUser,65),or(mload(add(signUser, 65)),and(mload(add(data, 333)), 255)))
        }
    }
    function get_signOper(bytes data) internal pure returns(bytes signOper){
        signOper = new bytes(65);
        assembly{
            mstore(add(signOper,32),mload(add(data, 365)))
            mstore(add(signOper,64),mload(add(data, 397)))
            mstore(add(signOper,65),or(mload(add(signOper, 65)),and(mload(add(data, 398)), 255)))
        }
    }
    function get_x(bytes data) internal pure returns(bytes32 X){
        assembly{
            X := mload(add(data, 430))
        }
    }
    
    // function verifySignature(bytes data,bytes sig) internal pure returns (bool ret){
    //     ret = false;
    // }
    
    function decodeAndVerifyProof(bytes data)internal pure returns (Proof){
        require(data.length == 430);
        bytes memory user_sign_data = new bytes(236);
        assembly{
            mstore(add(user_sign_data,32),mload(add(data, 32)))
            mstore(add(user_sign_data,64),mload(add(data, 64)))
            mstore(add(user_sign_data,96),mload(add(data, 96)))
            mstore(add(user_sign_data,128),mload(add(data, 128)))
            mstore(add(user_sign_data,160),mload(add(data, 160)))
            mstore(add(user_sign_data,192),mload(add(data, 192)))
            mstore(add(user_sign_data,224),mload(add(data, 224)))
            mstore(add(user_sign_data,256),mload(add(data, 256)))
        }
        require(proofVerify.ecverify(keccak256(user_sign_data), get_signUser(data)) == get_from(data),"the signaure of user is wrong");
        
        bytes memory operator_sign_data = new bytes(268);
        assembly{
            mstore(add(operator_sign_data,32),mload(add(data, 32)))
            mstore(add(operator_sign_data,64),mload(add(data, 64)))
            mstore(add(operator_sign_data,96),mload(add(data, 96)))
            mstore(add(operator_sign_data,128),mload(add(data, 128)))
            mstore(add(operator_sign_data,160),mload(add(data, 160)))
            mstore(add(operator_sign_data,192),mload(add(data, 192)))
            mstore(add(operator_sign_data,224),mload(add(data, 224)))
            mstore(add(operator_sign_data,256),mload(add(data, 256)))
            mstore(add(operator_sign_data,288),mload(add(data, 288)))
        }
        require(proofVerify.ecverify(keccak256(operator_sign_data), get_signOper(data)) == get_operator(data),"the signaure of operator is wrong");
        
        require(get_hashX(data) == keccak256(abi.encodePacked(get_x(data))),"bad secret for this proof");
        
        Proof memory new_proof;
        new_proof.from_ = get_from(data);
        new_proof.to = get_to(data);
        new_proof.operator = get_operator(data);
        new_proof.totalTransferCnt = get_totalTransferCnt(data);
        new_proof.totalWithdrawCnt = get_totalWithdrawCnt(data);
        new_proof.nonce = get_nonce(data);
        new_proof.timestamp_ = get_timestamp(data);
        //new_proof.hashX = get_hashX(data);
        new_proof.chlID = get_chlID(data);
        new_proof.deltaCnt = get_deltaCnt(data);
        new_proof.targetTotalCnt = get_targetTotalCnt(data);
        //new_proof.signUser = get_signUser(data);
        //new_proof.signOper = get_signOper(data);
        //new_proof.x = get_x(data);
        return new_proof;
    }
    
    //event
    event HashEvent(bytes32 index);
    event WithdrawEvent(address from, address to, address operator, uint256 transferCnt, uint256 withdrawCnt, uint64 nonce, int64 timeStamp_, uint256 deposition, uint256 totalTransfer, uint256 totalWithdraw);
    event BadProofEvent(address from, address to, address operator, uint256 transferCnt, uint256 withdrawCnt, uint64 nonce, int64 timeStamp_, uint256 chlID,uint8 errcode);
}


