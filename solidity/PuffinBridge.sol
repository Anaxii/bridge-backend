// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

abstract contract PuffinBridge is Ownable, Pausable {

    uint256 public threshold = 0;
    uint256 public chainId;
    mapping(address => bool) public isVoter;

    address public puffinKYC;
    address public puffinAssets;
    address public puffinWarmWallet;
    address public subnetTokenDeployer;

    mapping(bytes32 => bool) public bridgeInComplete;
    mapping(bytes32 => bool) public bridgeOutComplete;
    mapping(bytes32 => bool) public bridgeIds;
    mapping(bytes32 => bool) public requiresWarmWallet;

    mapping(bytes32 => uint256) public requestCount;
    mapping(address => mapping(bytes32 => bool)) public hasVoted;

    mapping(bytes32 => BridgeRequest) public requestInfo;

    event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId);
    event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId);
    event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId);
    event BridgeOutCanceled(address indexed user, bytes32 indexed id);

    struct BridgeRequest {
        bytes32 id;
        address user;
        address asset;
        uint256 amount;
        uint256 expiry;
    }

    constructor(uint256 _chainId) {
        isVoter[msg.sender] = true;
        puffinWarmWallet = msg.sender;
        chainId = _chainId;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function markInComplete(bytes32 requestId) public {
        require(isVoter[msg.sender]);
        require(!bridgeInComplete[requestId], "PuffinBridge: Request already complete");
        bridgeInComplete[requestId] = true;
    }

    function cancelOut(bytes32 requestId) public onlyOwner {
        require(!bridgeOutComplete[requestId], "PuffinBridge: Request already complete");
        bridgeOutComplete[requestId] = true;
        emit BridgeOutCanceled(msg.sender, requestId);
    }

    function setThreshold(uint256 _threshold) external onlyOwner {
        threshold = _threshold;
    }

    function addVoter(address user) external onlyOwner {
        isVoter[user] = true;
    }

    function removeVoter(address user) external onlyOwner {
        isVoter[user] = false;
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

    function setSubnetTokenDeployer(address _contract) external onlyOwner {
        subnetTokenDeployer = _contract;
    }

    function transferWarm(address asset, uint256 amount) external onlyOwner {
        IERC20(asset).transfer(puffinWarmWallet, amount);
    }

    function getRequestInfo(
        bytes32 requestId
    ) public view returns (BridgeRequest memory) {
        return requestInfo[requestId];
    }
}

contract PuffinMainnetbridge is PuffinBridge, ReentrancyGuard  {

    constructor(uint256 _chainId) PuffinBridge(_chainId) {

    }

    function proposeOut(
        address asset,
        address user,
        uint256 amount,
        bytes32 requestId,
        uint256 fromChainId
    ) external whenNotPaused nonReentrant {
        require(PuffinApprovals(puffinAssets).isApproved(asset), "PuffinBridge: Asset is not approved");
        require(!bridgeOutComplete[requestId], "PuffinBridge: Request already complete");
        require(isVoter[msg.sender], "PuffinBridge: Not a voter");
        require(!hasVoted[msg.sender][requestId], "PuffinBridge: User has voted");

        hasVoted[msg.sender][requestId] = true;

        if (threshold == 0)
            require(msg.sender == owner());

        if (requestCount[requestId] == 0) {
            requestInfo[requestId] = BridgeRequest(requestId, user, asset, amount, (block.timestamp + (1 days)));
            requestCount[requestId] ++;
        } else {
            require(requestInfo[requestId].user == user && requestInfo[requestId].asset == asset && requestInfo[requestId].amount == amount, "PuffinBridge: Invalid input");
            requestCount[requestId] ++;
        }

        if (requestCount[requestId] >= threshold) {
            bridgeOutComplete[requestId] = true;
            if (IERC20(asset).balanceOf(address(this)) < amount) {
                emit BridgeOutWarm(user, asset, amount, requestId, fromChainId);
                requiresWarmWallet[requestId] = true;
                return;
            }
            emit BridgeOut(user, asset, amount, requestId, fromChainId);

            IERC20(asset).transfer(user, amount);
        }
    }

    function bridgeIn(
        uint256 amount,
        address asset
    ) public whenNotPaused {
        require(PuffinApprovals(puffinKYC).isApproved(msg.sender), "PuffinBridge: User is not KYC approved");
        require(PuffinApprovals(puffinAssets).isApproved(asset), "PuffinBridge: Asset is not approved");
        require(amount > 0);

        bytes32 id = keccak256(abi.encodePacked(amount, msg.sender, block.timestamp, asset));
        require(!bridgeIds[id], "PuffinBridge: Invalid ID, try again");
        bridgeIds[id] = true;
        IERC20(asset).transferFrom(msg.sender, address(this), amount);
        emit BridgeIn(msg.sender, asset, amount, id, chainId);
    }
}

contract PuffinSubnetbridge is PuffinBridge, ReentrancyGuard  {

    constructor(uint256 _chainId) PuffinBridge(_chainId) {

    }

    function proposeOut(
        address asset,
        address user,
        uint256 amount,
        bytes32 requestId,
        uint256 fromChainId
    ) public whenNotPaused nonReentrant {
        require(!bridgeOutComplete[requestId], "PuffinBridge: Request already complete");
        require(isVoter[msg.sender], "PuffinBridge: Not a voter");
        require(!hasVoted[msg.sender][requestId], "PuffinBridge: User has voted");

        hasVoted[msg.sender][requestId] = true;

        if (threshold == 0)
            require(msg.sender == owner(), "PuffinBridge: Threshold 0 not owner");

        if (requestCount[requestId] == 0) {
            requestInfo[requestId] = BridgeRequest(requestId, user, asset, amount, (block.timestamp + (1 days)));
            requestCount[requestId] ++;
        } else {
            require(requestInfo[requestId].user == user && requestInfo[requestId].asset == asset && requestInfo[requestId].amount == amount, "PuffinBridge: Invalid input");
            requestCount[requestId] ++;
        }

        if (requestCount[requestId] >= threshold) {
            bridgeOutComplete[requestId] = true;
            PuffinERC20Deployer(subnetTokenDeployer).mint(asset, user, amount, 1);
            emit BridgeOut(user, asset, amount, requestId, fromChainId);
        }
    }

    function bridgeIn(
        uint256 amount,
        address asset,
        uint256 chainId
    ) public whenNotPaused {
        require(IAllowList(0x0200000000000000000000000000000000000002).readAllowList(msg.sender) > 0, "PuffinBridge: User is not KYC approved");
        require(amount > 0, "PuffinBridge: Amount == 0");

        bytes32 id = keccak256(abi.encodePacked(amount, msg.sender, block.timestamp, asset));
        require(!bridgeIds[id], "PuffinBridge: Invalid ID, try again");
        bridgeIds[id] = true;
        IERC20(asset).transferFrom(msg.sender, address(this), amount);
        IERC20(asset).approve(subnetTokenDeployer, amount);
        PuffinERC20Deployer(subnetTokenDeployer).burn(asset, address(this), amount, 0);
        emit BridgeIn(msg.sender, PuffinERC20Deployer(subnetTokenDeployer).subnetToMainnetTokenAddress(asset, chainId), amount, id, chainId);
    }
}