contract PuffinBridge is Ownable, Pausable {

    uint256 public threshold = 0;

    mapping(address => bool) public isVoter;

    address public puffinKYC;
    address public puffinAssets;
    address public puffinWarmWallet;

    mapping(bytes32 => bool) public bridgeInComplete;
    mapping(bytes32 => bool) public bridgeOutComplete;
    mapping(bytes32 => uint256) public requestCount;

    mapping(bytes32 => BridgeRequest) private requestInfo;

    event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id);
    event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id);
    event BridgOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id);
    event BridgOutCanceled(address indexed user, bytes32 id);

    struct BridgeRequest {
        bytes32 id;
        address user;
        address asset;
        uint256 amount;
        uint256 expiry;
    }

    constructor() {
        isVoter[msg.sender] = true;
        puffinWarmWallet = msg.sender;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function bridgeIn(uint256 amount, address asset) public whenNotPaused returns (bytes32) {
        require(PuffinApprovals(puffinKYC).isApproved(msg.sender), "PuffinBridge: User is not KYC approved");
        require(PuffinApprovals(puffinAssets).isApproved(asset), "PuffinBridge: Asset is not approved");
        require(amount > 0);

        bytes32 id = keccak256(abi.encodePacked(amount, msg.sender, block.timestamp, asset));
        IERC20(asset).transferFrom(msg.sender, address(this), amount);
        emit BridgeIn(msg.sender, asset, amount, id);
        return id;
    }

    function markInComplete(bytes32 requestId) public onlyOwner {
        require(!bridgeInComplete[requestId], "PuffinBridge: Request already complete");
        bridgeInComplete[requestId] = true;
    }

    function proposeOut(address asset, address user, uint256 amount, bytes32 requestId) public whenNotPaused {
        require(PuffinApprovals(puffinAssets).isApproved(asset), "PuffinBridge: Asset is not approved");
        require(!bridgeOutComplete[requestId], "PuffinBridge: Request already complete");
        require(isVoter[msg.sender], "PuffinBridge: Not a voter");
        require(user != address(0));

        if (requestCount[requestId] == 0) {
            requestInfo[requestId] = BridgeRequest(requestId, user, asset, amount, (block.timestamp + (1 days)));
            requestCount[requestId] ++;
        } else {
            require(requestInfo[requestId].user == user && requestInfo[requestId].asset == asset && requestInfo[requestId].amount == amount, "PuffinBridge: Invalid input");
            requestCount[requestId] ++;
        }

        if (requestCount[requestId] >= threshold) {
            bridgeOutComplete[requestId] = true;
            if (IERC20(asset).balanceOf(address(this)) > amount) {
                emit BridgOutWarm(user, asset, amount, requestId);
                return;
            }
            emit BridgeOut(user, asset, amount, requestId);

            IERC20(asset).transfer(user, amount);
        }
    }

    function cancelOut(bytes32 requestId) public onlyOwner {
        require(!bridgeOutComplete[requestId], "PuffinBridge: Request already complete");
        bridgeOutComplete[requestId] = true;
        emit BridgOutCanceled(msg.sender, requestId);
    }

    function addVoter(address user) external onlyOwner {
        isVoter[user] = true;
    }

    function removeVoter(address user) external onlyOwner {
        isVoter[user] = false;
    }

    function getRequestInfo(bytes32 requestId) public view returns (BridgeRequest memory) {
        return requestInfo[requestId];
    }

    function setKYC(address _contract) external onlyOwner {
        puffinKYC = _contract;
    }

    function setAssets(address _contract) external onlyOwner {
        puffinAssets = _contract;
    }

    function setWarm(address _contract) external onlyOwner {
        puffinWarmWallet = _contract;
    }
}