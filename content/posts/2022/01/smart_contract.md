---
author: "jdlau"
date: 2022-01-26
linktitle: 智能合约
menu:
next:
prev:
title: 智能合约
weight: 10
categories: ['blockchain', 'smart contract']
tags: ['eth', 'evm']
---

## 智能合约

[智能合约wiki](https://zh.wikipedia.org/wiki/%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6)

> （英语：Smart contract）是一种特殊协议，在区块链内制定合约时使用，当中内含了代码函数 (Function)，亦能与其他合约进行交互、做决策、存储资料及发送以太币等功能。智能合约主要提供**验证及执行合约内所订立的条件**。智能合约允许在没有第三方的情况下进行可信交易。这些**交易可追踪且不可逆转**。

### 安全问题

> 智能合约是“执行合约条款的计算机交易协议”。区块链上的**所有用户都可以看到**基于区块链的智能合约。但是，这会导致包括安全漏洞在内的所有漏洞都可见，并且可能无法迅速修复。
>
> 这样的攻击难以迅速解决，例如：
>
>> 2016年6月The DAOEther的漏洞造成损失5000万美元，而开发者试图达成共识的解决方案。DAO的程序在黑客删除资金之前有一段时间的延迟。以太坊软件的一个硬分叉在时限到期之前完成了攻击者的资金回收工作。
>
> 以太坊智能合约中的问题包括合约编程Solidity、编译器错误、以太坊虚拟机错误、对区块链网络的攻击、程序错误的不变性以及其他尚无文档记录的攻击。
>
>> 2018年4月22日， BeautyChain智能合约出现重大漏洞，黑客通过此漏洞无限生成代币，导致 BitEclipse (BEC)的价值接近归零。同月25日，SmartMesh出现疑似重大安全漏洞，宣布暂停所有SMT交易和转账直至另行通知，导致损失约1.4亿美金。28日，EOS被指可能存在BEC代币合约类似的整数溢出漏洞，但没消息详细说明。5月24日， BAI交易存在大量异常问题， 损失金额未知。8月22日， GODGAME 合约被黑客入侵，GOD智能合约上的以太坊总数归零。

### 合约开发、测试和部署

[eth智能合约文档](https://ethereum.org/en/developers/docs/smart-contracts/)

vending machine(自动售货机): `money + snack selection = snack dispensed`, 给钱并选择小吃，小吃就会出来 -- 是给刚好的钱，还是过量的钱，过量了在发放小吃的同时退钱呢？

合约长这样：

```sol
// 表明使用的sol版本 
pragma solidity 0.8.7;

// Solidity 合约类似于面向对象语言中的类。合约中有用于数据持久化的状态变量，和可以修改状态变量的函数。 调用另一个合约实例的函数时，会执行一个 EVM 函数调用，这个操作会切换执行时的上下文，这样，前一个合约的状态变量就不能访问了。
contract VendingMachine {

    // Declare state variables of the contract
    address public owner; // owner变量
    mapping (address => uint) public cupcakeBalances; // cupcakeBalances变量

    // 创建合约时，合约的构造函数会执行一次。构造函数是可选的。只允许有一个构造函数，这意味着不支持重载。
    // When 'VendingMachine' contract is deployed:
    // 1. set the deploying address as the owner of the contract
    // 2. set the deployed smart contract's cupcake balance to 100
    constructor() {
        owner = msg.sender; // 设置部署本合约的地址为合约所有者

        // address(this)是将this转型为地址吗？
        // this不是代表合约对象吗，还能转为address？
        cupcakeBalances[address(this)] = 100; // 设置蛋糕余量
    }

    // Allow the owner to increase the smart contract's cupcake balance
    function refill(uint amount) public { // public表示方法可导出
        require(msg.sender == owner, "Only the owner can refill."); // 如果前面的条件不成立，则报错，后面为内容；此处要求消息的发送者必须是本合约所有者
        cupcakeBalances[address(this)] += amount; // 补充蛋糕
    }

    // Allow anyone to purchase cupcakes
    function purchase(uint amount) public payable { // payable表示方法含有支付逻辑
        require(msg.value >= amount * 1 ether, "You must pay at least 1 ETH per cupcake"); // 此处要求每个蛋糕最少支付一个eth
        require(cupcakeBalances[address(this)] >= amount, "Not enough cupcakes in stock to complete this purchase"); // 此处要求蛋糕余量不小于需要的数量
        cupcakeBalances[address(this)] -= amount; // 本合约所有者减少蛋糕
        cupcakeBalances[msg.sender] += amount; // 消息发送者添加蛋糕
    }
}
```

[address(this)](https://ethereum.stackexchange.com/questions/40018/what-is-addressthis-in-solidity)

> **this** refers to **the instance of the contract where the call is made** (you can have multiple instances of the same contract). -- 调用发生时合约的实例
> 
> **address(this)** refers to **the address of the instance of the contract where the call is being made**. -- 调用发生的地方的合约的实例的地址
>
> **msg.sender** refers to **the address where the contract is being called from**.
>
> Therefore, address(this) and msg.sender are two unique addresses, the first referring to **the address of the contract instance** and the second referring to **the address where the contract call originated from**.
>
>> this: 实例; 
>>
>> address(this): 实例的地址; 
>>
>> msg.sender: 消息发送者，最初的合约调用者。

看上面的合约，`msg`是哪里来的，它里面有些什么东西呢？消息发送者买了蛋糕之后，还能再卖出去吗？

`msg`是一个[全局变量](https://medium.com/upstate-interactive/what-you-need-to-know-about-msg-global-variables-in-solidity-566f1e83cc69)

> For instance, `msg.sender` is always the address where the current (external) function call came from. 
>
>> For instance, if a function call came from a user or smart contract with the address 0xdfad6918408496be36eaf400ed886d93d8a6c180 then msg.sender equals 0xdfad6918408496be36eaf400ed886d93d8a6c180.
>
> `msg.data` — The complete calldata which is a non-modifiable, non-persistent area where function arguments are stored and behave mostly like memory
>
> `msg.gas` — Returns the available gas remaining for a current transaction (you can learn more about gas in Ethereum here). The function `gasleft` was previously known as `msg.gas`, which was deprecated in version 0.4.21 and removed in version 0.5.0.
>
> `msg.sig` — The first four bytes of the calldata for a function that specifies the function to be called (i.e., it’s function identifier), 函数标识
>
> `msg.value` — The amount of wei sent with a message to a contract (wei is a denomination of ETH), 1 ETH = 10^18 wei

[更多全局变量和函数](hhttps://docs.soliditylang.org/en/develop/units-and-global-variables.html?highlight=special%20variables%20and%20functions)

> The special variables and functions are always available globally and are mainly used to provide information about the blockchain (i.e., transactions, contract info, address info, error handling, math and crypto functions).

- Block and Transaction Properties

- ABI Encoding and Decoding Functions

- Members of bytes

- Error Handling

- Mathematical and Cryptographic Functions

- Members of Address Types

- Contract Related

- Type Information

[一些工具库](https://ethereum.org/en/developers/docs/smart-contracts/testing/)

> [truffle](https://github.com/trufflesuite/truffle) - A tool for developing smart contracts. Crafted with the finest cacaos. 
>
> [Waffle](https://github.com/EthWorks/Waffle) - A framework for advanced smart contract development and testing (based on ethers.js).
>
> [Solidity-Coverage](https://github.com/sc-forks/solidity-coverage) - Alternative solidity code coverage tool.
>
> [hevm](https://github.com/dapphub/dapptools/tree/master/src/hevm) - Implementation of the EVM made specifically for unit testing and debugging smart contracts.
>
> [Whiteblock Genesis](https://github.com/whiteblock/genesis) - An end-to-end development sandbox and testing platform for blockchain.
>
> [OpenZeppelin Test Environment](https://github.com/OpenZeppelin/openzeppelin-test-environment)(已归档) - Blazing fast smart contract testing. One-line setup for an awesome testing experience.
>
> [OpenZeppelin Test Helpers](https://github.com/OpenZeppelin/openzeppelin-test-helpers) - Assertion library for Ethereum smart contract testing. Make sure your contracts behave as expected!

### 实战

安装[`ETH钱包MetaMask(firefox插件)`](https://metamask.io/download/)

新建账户，会提示输入密码，然后生成成功后，会有一堆助记词，需要记录好：

```
jeans term salt true cereal hobby cheese awesome link nice never choose
```

钱包默认连接的是主网，因为主网要花钱买币，所以我们选择测试网。在主网那里点击设置打开`Show test networks`配置。

获得地址：`0x758c40f09207f9e6F72A8C24029Be865D28eF219`

[eth rinkeby测试网水龙头](https://faucet.rinkeby.io/)

为上述地址去水龙头领取币。

需要先把上述地址发到推特、脸书、谷家任一地方，然后再使用对应帖子的地址(如：)来获取。

如果报错`nsufficient funds for gas * price + value`，可尝试使用[另外的网站](https://faucets.chain.link/rinkeby)直接使用地址获取。

拿到之后，在`metamask`查看余额。

[eth rinkeby测试网浏览器](https://rinkeby.etherscan.io/)

在浏览器上可以查找到上述的[转账记录](https://rinkeby.etherscan.io/address/0x758c40f09207f9e6F72A8C24029Be865D28eF219)

[参照这个游戏示例](https://github.com/upstateinteractive/blockchain-puzzle)

[编译、部署合约](http://remix.ethereum.org/)

TODO:

## evm

[代码](github.com/ethereum/go-ethereum/core/vm)

- stack based

- register based

## wasm

[切换到ewasm的背景和好处](https://medium.com/chainsafe-systems/ethereum-2-0-a-complete-guide-ewasm-394cac756baf)

[wasm工具集](https://github.com/webassembly/wabt)

## solidity

[contract](https://learnblockchain.cn/docs/solidity/contracts.html)

## NFT

[NFT wiki](https://zh.wikipedia.org/wiki/%E9%9D%9E%E5%90%8C%E8%B3%AA%E5%8C%96%E4%BB%A3%E5%B9%A3)

> 非同质化代币（英语：Non-Fungible Token，简称：NFT），是一种被称为区块链数字账本上的数据单位，**每个代币可以代表一个独特的数字资料**，作为虚拟商品所有权的电子认证或证书。由于其不能互换，非同质化代币可以代表数字文件，如画作、声音、视频、游戏中的项目或其他形式的创意作品。**虽然文件（作品）本身是可以无限复制的，但代表它们的代币在其底层区块链上被追踪，并为买家提供所有权证明**。诸如以太币、比特币等加密货币都有自己的代币标准以定义对NFT的使用。
>
> 非同质化代币一种存储在区块链（数位账本）上的数据单位，它可以**代表艺术品等独一无二的数字物品**。其是一种加密代币，但与比特币等加密货币不同，其**不可互换**。一个非同质化代币是透过上传一个文件，如艺术品，到非同质化代币拍卖市场。**这将创建一个记录在数字账本上的文件副本作为非同质化代币，它可以用加密货币购买和转售**。虽然艺术家可以出售代表作品的非同质化代币，但**艺术家仍然可以保留作品的著作权，并创造更多的同一作品的非同质化代币**。非同质化代币的**买家并不能获得对作品的独家访问权**，买家也不能获得对原始数字文件的独占。将某一作品作为非同质化代币上传的人不必证明他们是原创艺术家，在许多争议案例中，在未经创作者许可的情况下，艺术品被盗用于非同质化代币。

[eth nft](https://ethereum.org/zh/nft/)

- 一种将任何独特物品表现为以太坊资产的方式。
- NFT 为内容创建人赋予了比以往更多的权力。
- 由以太坊区块链上的智能合约提供支持。

> NFT 是我们用以代表独特物品所有权的代币。 NFT 让我们把诸如艺术品、收藏品、甚至房地产等物品代币化。 **他们一次只有一个正式主人，并且受到以太坊区块链的保护** - 没有人可以修改所有权记录或者根据现有的 NFT 复制粘贴一份新的。
>
>> NFT 代表非同质化代币。非同质化是一个经济术语，您可以用它来描述家具、歌曲文件或您的电脑等物品。**这些东西不能与其他物品互换，因为它们具有独特属性**。
>>
>> 另一方面，**同质化物品可以互换，这取决于它们的价值而非独特属性**。 例如，ETH 或美元具有同质化属性，因为 1 ETH/1 USD 可以兑换成另外的 1 ETH/1 USD。

当他们(数字产品作者)出售内容时，资金直接转给他们。如果新所有者随后出售 NFT，原创建人(数字产品作者)甚至可以自动收到版税。这在每次出售时都有保证，因为创建人的地址是代币元数据的一部分 - 元数据无法修改。
