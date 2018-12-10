pragma solidity ^0.4.24;

//wandapp ETH.
import "./wandappBase.sol";

contract wandappETH is wandappBase {
    //payment channel.
    struct Channel {
        uint256  chlID;          //channel id.
        uint256  depositeCnt;    //total deposition count.
        uint256  withdrawCnt;    //total withdraw count.
        uint256  transferCnt;    //total translation count.
        uint64   userNonce;      //user latest nonce which has been dealed with.
        uint64   operatorNonce;  //operator latest nonce which has been dealed with.
        uint256  exitWin;        //start time of the exit window.
        ChlState chlState;       //channel status.
        Proof exitCachedProof;    //cached proof for exiting process. 
    }
    
    /* Constructor
       Parmeters:
            tokenOrigAddr:  address of ERC20 token.
            operAddr:       operator address.
            duration:  duration of the exit window. 
    */
    constructor(address tokenOrigAddr, address operAddr, uint256 duration) wandappBase(tokenOrigAddr, operAddr, duration) public
    {
    }
	
    /* Deposite and create channel.
       Parmeters:
            amount: Token count to be deposited.
    */
    function deposite( uint256 amount ) public notNegative(amount) {
        require(DemoERC20(tokenAddr).transferFrom(msg.sender, this, amount));
        if(channels[msg.sender].chlID == 0) {
            //channel does not exist, so create it.
            channels[msg.sender].chlID = getChlID();
            channels[msg.sender].chlState = ChlState.NORMAL;
        }
        channels[msg.sender].depositeCnt = SafeMath.add(channels[msg.sender].depositeCnt, amount);
        
        emit DepositeEvent(
                channels[msg.sender].chlID,
                msg.sender,
                operatorAddr,
                tokenAddr,
                amount,
                channels[msg.sender].depositeCnt
            );
    }
    
    /* Return the deposition information.
       Parmeters:
            user: user address.
    */
    function getChannelInfo(address user) public view returns(uint256,uint256,uint256,uint256,uint64,uint64,uint256){
        return (channels[user].chlID,
                channels[user].depositeCnt,
                channels[user].withdrawCnt,
                channels[user].transferCnt,
                channels[user].userNonce,
                channels[user].operatorNonce,
                channels[user].exitWin);
    }
    
    /* withdraw deposition by user.
            Parmeters:
                proof: user proof byte flow.
    */
    function withdraw( bytes proof ) public {
        Proof memory p = decodeAndVerifyProof(proof); //verify user signature, operator signature and hashx = hash(X)
        require(channels[p.from_].chlState != ChlState.CREATOREXITING,"Wrong option in the state.");
        
        //common verification for proof.
        check_proof_params(p);
        //additional verfication for withdraw.
        require(p.nonce > channels[p.from_].userNonce, "Illegal nonce.");

        uint256 cnt = SafeMath.sub(p.totalWithdrawCnt, channels[p.from_].withdrawCnt);
        require(DemoERC20(tokenAddr).transfer(p.from_, cnt));
        channels[p.from_].withdrawCnt = p.totalWithdrawCnt;
        channels[p.from_].userNonce = p.nonce;
        
        emit WithdrawEvent(
            p.from_,
            p.to, 
            operatorAddr, 
            p.totalTransferCnt, 
            p.totalWithdrawCnt, 
            p.nonce, 
            p.timestamp_,
            channels[p.from_].depositeCnt,
            channels[p.from_].withdrawCnt,
            channels[p.from_].transferCnt
        );
    }

    /* withdraw deposition by operator.
            Parmeters:
                proof: operator proof byte flow.
    */
    function operatorWithdraw( bytes proof ) public onlyBy(operatorAddr) {
        Proof memory p = decodeAndVerifyProof(proof);
        require(channels[p.from_].chlState != ChlState.OPERATOREXITING,"Wrong option in the state.");
        
        //common verification for proof.
        check_proof_params(p);
        //additional verfication for operatorWithdraw.
        require(p.nonce > channels[p.from_].operatorNonce, "Illegal nonce.");

        uint256 cnt = SafeMath.sub(p.totalTransferCnt, channels[p.from_].transferCnt);
        require(DemoERC20(tokenAddr).transfer(operatorAddr, cnt));
        channels[p.from_].transferCnt = p.totalTransferCnt;
        channels[p.from_].operatorNonce = p.nonce;
        
        emit WithdrawEvent(
            p.from_,
            p.to, 
            p.operator, 
            p.totalTransferCnt, 
            p.totalWithdrawCnt, 
            p.nonce, 
            p.timestamp_,
            channels[p.from_].depositeCnt,
            channels[p.from_].withdrawCnt,
            channels[p.from_].transferCnt
        );
    }

    /* exit channel. Used by both user and operator.
            Parmeters:
                proof: proof.
    */
    function exit( bytes proof ) public {
        Proof memory p = decodeAndVerifyProof(proof);
        
        //common verification for proof.
        check_proof_params(p);
        
        Channel storage chl = channels[p.from_];
        if(chl.chlState == ChlState.NORMAL){
            chl.exitWin = now + exitWinDuration; //timeout when exitWin.
            if(msg.sender == operatorAddr){
                require(p.nonce >= channels[p.from_].operatorNonce, "Illegal nonce.");
                chl.chlState = ChlState.OPERATOREXITING;
            }else{
                require(p.nonce >= channels[p.from_].userNonce, "Illegal nonce.");
                chl.chlState = ChlState.CREATOREXITING;
            }
            channels[p.from_].exitCachedProof = p;  //cache the proof for later process.
        } else {
            if(chl.chlState == ChlState.OPERATOREXITING) {
                require(msg.sender == p.from_, "Wrong caller in this state.");
            }
            if(chl.chlState == ChlState.CREATOREXITING) {
                require(msg.sender == operatorAddr, "Wrong caller in this state.");
            }
            
            if(p.nonce > chl.exitCachedProof.nonce) {
                redeem(p.from_, p);
            } else {
                redeem(p.from_, chl.exitCachedProof);
            }
        }

        emit ExitEvent(
            p.from_,
            p.to,
            operatorAddr,
            channels[p.from_].transferCnt,
            channels[p.from_].withdrawCnt,
            p.nonce,
            p.timestamp_,
            channels[p.from_].chlID,
            channels[p.from_].chlState);
    }
    
    /* close channel by channel exiting initiator.
            Parmeters:
                None.
    */
    function settle() public {
        Channel storage chnl = channels[msg.sender];
        require(chnl.chlID != 0x0,"unknowed channel");
        require(chnl.exitWin <= now,"the time of settlinng channel isn't  coming");
        require((chnl.chlState == ChlState.CREATOREXITING) ||
            ((chnl.chlState == ChlState.OPERATOREXITING) && (msg.sender == operatorAddr)),
            "User has no auth in this state.");
        redeem(msg.sender, chnl.exitCachedProof); //no chanllage so we can redeem by the cached proof.
    }
    
    function check_proof_params ( Proof proof) internal view {
        require(proof.nonce > 0,"Illegal nonce");
        require(proof.deltaCnt >= 0, "Illegal nonce");
        require(proof.operator == operatorAddr, "unknowed operator address");
        require((channels[proof.from_].chlID != 0) && (channels[proof.from_].chlID == proof.chlID),"channel out of date");
        
        require(SafeMath.add(proof.totalWithdrawCnt,proof.totalTransferCnt) <= channels[proof.from_].depositeCnt,"have not enough deposition");
        require(proof.totalWithdrawCnt >= channels[proof.from_].withdrawCnt, "withdraw count is smaller than the last one");
        require(proof.totalTransferCnt >= channels[proof.from_].transferCnt, "transfer count is smaller than the last one");
    }

    function redeem(address from, Proof p) internal{
        DemoERC20 token = DemoERC20(tokenAddr);
        Channel storage chl = channels[from];
        
        uint256 value_to_user = SafeMath.sub(chl.depositeCnt, SafeMath.add(chl.withdrawCnt, p.totalTransferCnt));
        uint256 value_to_operator = SafeMath.sub(p.totalTransferCnt, chl.transferCnt);
        if(value_to_user > 0){
            require(token.transfer(from,value_to_user),"tranfer token to creator failed");
            emit SettleEvent(from,value_to_user,chl.userNonce, chl.operatorNonce);
        }
        if(value_to_operator > 0){
            require(token.transfer(operatorAddr,value_to_operator),"tranfer token to operator failed");
            emit SettleEvent(operatorAddr,value_to_operator,chl.userNonce, chl.operatorNonce);
        }

	    delete channels[from];
    }
    
    Channel channel;
    mapping(address=>Channel) channels;
    
    //Event
    event DepositeEvent(uint256 chlId, address creator, address operator, address tokenAddr, uint256 amount, uint256 totalBalance);
    event ExitEvent(address from, address to, address operator, uint256 transferCnt, uint256 withdrawCnt, uint64 nonce, int64 timeStamp, uint256 chlID,ChlState chlstate);
    event WrongAuth(address sender,address target);
    event SettleEvent(address to,uint256 value, uint256 userNonce,uint256 operNonce);
}




