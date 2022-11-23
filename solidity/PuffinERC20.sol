
contract PuffinERC20Deployer is Ownable {

    address public puffinApprovedAssets;
    address public bridgeContract;

    bool public isGlobalPaused;
    mapping(address => bool) public isUserPaused;
    mapping(address => bool) public isTokenPaused;

    mapping(address => bool) public isMinter;
    mapping(address => uint256) public mainnetTokenChainID;
    mapping(address => address) public mainnetToSubnetTokenAddress;
    mapping(address => mapping(uint256 => address)) public subnetToMainnetTokenAddress;
    mapping(uint256 => address) public chainIdToMainnetTokenAddress;
    mapping(uint256 => address) public chainIdToSubnetTokenAddress;

    event NewToken(address indexed mainnetToken, address indexed subnetToken, uint256 indexed chainId);
    event NewMainnetToken(address indexed mainnetToken, address indexed subnetToken, uint256 indexed chainId);
    event Mint(address indexed token, address indexed user, uint256 indexed amount);
    event Burn(address indexed token, address indexed user, uint256 indexed amount);

    function setNewMainnetToken(address mainnetToken, address subnetToken, uint256 chainId, string memory name, string memory symbol) external onlyOwner {
        if (mainnetToken == subnetToken) {
            require(PuffinApprovals(puffinApprovedAssets).isApproved(mainnetToken), "PuffinERC20Deployer: Token isnt allowed");
            require(mainnetToSubnetTokenAddress[mainnetToken] == address(0), "PuffinERC20Deployer: Subnet token already exists");
            PuffinERC20 newToken = new PuffinERC20(name, symbol);
            require(subnetToMainnetTokenAddress[address(newToken)][chainId] == address(0), "PuffinERC20Deployer: ChainID already exists");
            subnetToMainnetTokenAddress[address(newToken)][chainId] = mainnetToken;
            mainnetToSubnetTokenAddress[mainnetToken] = address(newToken);
            mainnetTokenChainID[mainnetToken] = chainId;
            emit NewToken(mainnetToken, subnetToken, chainId);
        } else {
            require(PuffinERC20(subnetToken).isPuffinERC20());
            require(PuffinApprovals(puffinApprovedAssets).isApproved(mainnetToken), "PuffinERC20Deployer: Token isnt allowed");
            require(subnetToMainnetTokenAddress[subnetToken][chainId] == address(0), "PuffinERC20Deployer: ChainID already exists");
            subnetToMainnetTokenAddress[subnetToken][chainId] = mainnetToken;
            mainnetToSubnetTokenAddress[mainnetToken] = subnetToken;
            mainnetTokenChainID[mainnetToken] = chainId;
            emit NewMainnetToken(mainnetToken, subnetToken, chainId);
        }
    }

    function mint(address token, address recipient, uint256 amount, uint256 _type) external {
        require(isMinter[_msgSender()], "PuffinERC20Deployer: Sender is not a minter");
        PuffinERC20 _contract;
        if (_type == 0) {
            _contract = PuffinERC20(token);
        } else {
            _contract = PuffinERC20(mainnetToSubnetTokenAddress[token]);
        }
        _contract.mint(amount);
        _contract.transfer(recipient, amount);
        emit Mint(address(_contract), recipient, amount);
    }

    function burn(address token, address recipient, uint256 amount, uint256 _type) external {
        require(recipient == _msgSender());
        require(isMinter[_msgSender()]);
        PuffinERC20 _contract;
        if (_type == 0) {
            _contract = PuffinERC20(token);
        } else {
            _contract = PuffinERC20(mainnetToSubnetTokenAddress[token]);
        }
        _contract.transferFrom(recipient, address(this), amount);
        _contract.burn(amount);
        emit Burn(address(_contract), recipient, amount);
    }

    function setPuffinApprovedAssets(address _contract) external onlyOwner {
        puffinApprovedAssets = _contract;
    }

    function setBridgeContract(address _contract) external onlyOwner {
        bridgeContract = _contract;
    }

    function setMinter(address _minter, bool status) external onlyOwner {
        isMinter[_minter] = status;
    }
}

interface IPuffinERC20Deployer {
    function isGlobalPaused() external view returns (bool);
    function isUserPaused(address user) external view returns (bool);
    function isTokenPaused(address user) external view returns (bool);
}

contract PuffinERC20 is ERC20, Ownable {

    bool public isPuffinERC20 = true;

    constructor(string memory name, string memory symbol) ERC20(name, symbol) {

    }

    function mint(uint256 amount) external onlyOwner {
        _mint(owner(), amount);
    }

    function burn(uint256 amount) external onlyOwner {
        _burn(owner(), amount);
    }

    function transfer(address recipient, uint256 amount) override public returns (bool) {
        require(IAllowList(0x0200000000000000000000000000000000000002).readAllowList(recipient) > 0, "PuffinERC20: User unauthorized");
        require(!IPuffinERC20Deployer(owner()).isGlobalPaused(), "PuffinERC20: Global Pause");
        require(!IPuffinERC20Deployer(owner()).isUserPaused(_msgSender()), "PuffinERC20: User Pause");
        require(!IPuffinERC20Deployer(owner()).isTokenPaused(address(this)), "PuffinERC20: Token Pause");
        _transfer(_msgSender(), recipient, amount);
        return true;
    }
}
