pragma solidity ^0.4.24;
//pragma experimental ABIEncoderV2;

//wandapp for WAN.
import "./wandappBase.sol";

contract wandappWAN is wandappBase {
    struct InfoStore {
        uint256 chlID;                          //Channel id.
        uint256 totalNonce;                     //The last nonce which has been used in the way of totalTargetCnt. 
        uint256 totalDealedCnt;
        uint256 maxDeltaNonce;
        mapping(uint256=>bool) deltaNonces;
    }

    struct AccountStatus {
        mapping(address=>InfoStore) InfoStores;
    }

    /* Constructor
        Parmeters:
            tokenOrigAddr: erc20 token addresss.
            operAddr: operator address in the WAN side.
            operInProof: operator address in the ETH side.
            xTimeout: x timeout.
            exitWinDuration: the duration of the exit window.
    */
   constructor(address tokenOrigAddr, address operAddr, address operInProof, uint256 xTimeout, uint256 exitWinDuration) 
        wandappBase(tokenOrigAddr,operAddr, exitWinDuration) public {
            xDuration = xTimeout;           //Suggestion: 3600 seconds.
            operatorInProof = operInProof;  //the operator address in proof.
    }

    /* Operator deposition.
        Parmeters:
            amount: the amount to be deposited.
        This function can only be called by operator.
    */
    function deposite(uint256 amount) public onlyBy(operatorAddr) notNegative(amount) {
        require(DemoERC20(tokenAddr).transferFrom(operatorAddr, this, amount));
        totalDeposition = SafeMath.add(totalDeposition, amount);
        emit DepositeEvent(msg.sender, tokenAddr, amount, totalDeposition);
    }
    
    /* User gets their cross-transferred tokens in WAN side via this functon.
    */
    function deltaWithdraw(bytes proof) public {
        //channel info does not exist, then create it.
        Proof memory p = decodeAndVerifyProof(proof);
        
        if(accountStatus[p.to].InfoStores[p.from_].chlID == 0) {
            accountStatus[p.to].InfoStores[p.from_].chlID = p.chlID;
        }
        
        //common verification for proof.
        check_proof_params(p);
        //additional verification for channel id.
        require((p.chlID != 0) && (p.chlID == accountStatus[p.to].InfoStores[p.from_].chlID), "channel id does not match.");
        
        require(p.nonce > accountStatus[p.to].InfoStores[p.from_].totalNonce, "delta nonce out of range.");
        require(accountStatus[p.to].InfoStores[p.from_].deltaNonces[p.nonce] == false, "delta nonce has been used.");
        require(DemoERC20(tokenAddr).transfer(p.to, p.deltaCnt));
        accountStatus[p.to].InfoStores[p.from_].deltaNonces[p.nonce] = true;
        accountStatus[p.to].InfoStores[p.from_].totalDealedCnt = 
            SafeMath.add(accountStatus[p.to].InfoStores[p.from_].totalDealedCnt, p.deltaCnt);
        totalUsedDeposition = SafeMath.add(totalUsedDeposition, p.deltaCnt);
        
        if( p.nonce > accountStatus[p.to].InfoStores[p.from_].maxDeltaNonce ) {
            accountStatus[p.to].InfoStores[p.from_].maxDeltaNonce = p.nonce;
        }

        emit WithdrawEvent(
            p.from_,
            p.to, 
            operatorInProof, 
            p.totalTransferCnt, 
            p.totalWithdrawCnt, 
            p.nonce,
            p.timestamp_,
            p.deltaCnt,
            p.targetTotalCnt,
            0
        );
    }

    /* User gets their cross-transferred tokens in WAN side via this functon.
    */
    function totalWithdraw(bytes proof) public {
        Proof memory p = decodeAndVerifyProof(proof);
        
        if(accountStatus[p.to].InfoStores[p.from_].chlID == 0) {
            accountStatus[p.to].InfoStores[p.from_].chlID = p.chlID;
        }
        
        //common verification for proof.
        check_proof_params(p);
        //additional verification for channel id.
        require((p.chlID != 0) && (p.chlID == accountStatus[p.to].InfoStores[p.from_].chlID), "channel id does not match.");
            
        require(p.nonce > accountStatus[p.to].InfoStores[p.from_].totalNonce, "nonce too small.");
        require(p.nonce > accountStatus[p.to].InfoStores[p.from_].maxDeltaNonce, "nonce is smaller than max delta nonce.");
        uint256 cnt = SafeMath.sub(p.targetTotalCnt, accountStatus[p.to].InfoStores[p.from_].totalDealedCnt);
        require(DemoERC20(tokenAddr).transfer(p.to, cnt));
        totalUsedDeposition = SafeMath.add(totalUsedDeposition, cnt);
        accountStatus[p.to].InfoStores[p.from_].totalDealedCnt = p.targetTotalCnt;
        accountStatus[p.to].InfoStores[p.from_].totalNonce = p.nonce;
        
        emit WithdrawEvent(
            p.from_,
            p.to, 
            operatorInProof, 
            p.totalTransferCnt, 
            p.totalWithdrawCnt, 
            p.nonce,
            p.timestamp_,
            p.deltaCnt,
            p.targetTotalCnt,
            cnt
        );
    }

    /* exit.
        only can be used by operator.
    */
    function exit() public onlyBy(operatorAddr){
        exitWin = now;
        emit ExitEvent(operatorAddr, exitWin);
    }

    /* settle.
        only can be used by operator.
    */
    function settle() public onlyBy(operatorAddr) {
        uint256 aim = SafeMath.add(exitWin, exitWinDuration);
        if(now <= aim) {
            emit SettleEvent(now, aim, 1);  //It is not time for settle.
            return;
        }
        uint256 cnt = SafeMath.sub(totalDeposition, totalUsedDeposition);
        require(DemoERC20(tokenAddr).transfer(operatorAddr, cnt));
        emit SettleEvent(now, aim, 2);  //Settle done.
    }
    
    function check_proof_params ( Proof proof) internal view {
        require(proof.nonce > 0,"Illegal nonce");
        require(proof.deltaCnt >= 0, "Illegal nonce");
        require(proof.operator == operatorInProof, "unknowed operator address");
        require(uint256(proof.timestamp_) <= now, "timestamp issue.");
        require(SafeMath.add(uint256(proof.timestamp_), xDuration) >= now, "X timeout.");  //check x timeout. (safeMath does not support int256.)
    }

    mapping(address => AccountStatus) accountStatus;
    uint256 totalDeposition;        //the total depisition of operator.
    uint256 totalUsedDeposition;    //the total transferred count.
    uint256 xDuration;
    address operatorInProof;
    uint256 exitWin;

    /*  deposite event.
        operator: operator address.
        tokenAddr: erc20 token address.
        amount: the erc20 token count which has be deposited.
        totalBalance: the total erc20 deposition count.
    */
    event DepositeEvent(address operator, address tokenAddr, uint256 amount, uint256 totalBalance);
    event SettleEvent(uint256 current, uint256 aim, uint8 errcode);
    event ExitEvent(address operator, uint256 time);
}



