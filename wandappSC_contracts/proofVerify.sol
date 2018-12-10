pragma solidity ^0.4.24;

library proofVerify{

    //event E1(string,bytes);
    //event E2(string,bytes32);
    //event E3(string,address);
    function ecverify(bytes32 hashdata, bytes signature) internal pure returns (address signature_address)
    {
        //emit E2("data",hashdata);
        //emit E1("signature",signature);
        require(signature.length == 65);

        bytes32 r;
        bytes32 s;
        uint8 v;

        // The signature format is a compact form of:
        //   {bytes32 r}{bytes32 s}{uint8 v}
        // Compact means, uint8 is not padded to 32 bytes.
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))

            // Here we are loading the last 32 bytes, including 31 bytes of 's'.
            //v := byte(0, mload(add(signature, 96)))
            v := and(mload(add(signature, 65)), 255)
        }

        // Version of signature should be 27 or 28, but 0 and 1 are also possible
        if (v < 27) {
            v += 27;
        }

        require(v == 27 || v == 28);

        /* prefix is needed for geth only
        * https://github.com/ethereum/go-ethereum/issues/3731
        */
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        //bytes32 hashdata = keccak256(prefix, hashdata);
        //emit E2("hashdata",hashdata);
        //bytes32 hashdata = keccak256(data);
        signature_address = ecrecover(keccak256(abi.encodePacked(prefix, hashdata)), v, r, s);
        //emit E3("signature_address",signature_address);
        //emit E1("r",r);
        //emit E1("s",s);
        //emit E4("v",v);
        // ecrecover returns zero on error
        require(signature_address != address(0x0));

        //emit E2("address",signature_address);

        return signature_address;
    }
    
}
