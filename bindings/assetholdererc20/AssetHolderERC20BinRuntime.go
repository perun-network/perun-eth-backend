package assetholdererc20 // import "github.com/perun-network/perun-eth-backend/bindings/assetholdererc20"

// AssetHolderERC20BinRuntime is the runtime part of the compiled bytecode used for deploying new contracts.
var AssetHolderERC20BinRuntime = "6080604052600436106100705760003560e01c8063ae9ee18c1161004e578063ae9ee18c146100e7578063d945af1d14610122578063fc0c546a14610162578063fca0f7781461019657600080fd5b80631de26e1614610075578063295482ce1461008a57806353c2ed8e146100aa575b600080fd5b610088610083366004610dfb565b6101b6565b005b34801561009657600080fd5b506100886100a5366004610e62565b610230565b3480156100b657600080fd5b506002546100ca906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100f357600080fd5b50610114610102366004610edc565b60006020819052908152604090205481565b6040519081526020016100de565b34801561012e57600080fd5b5061015261013d366004610edc565b60016020526000908152604090205460ff1681565b60405190151581526020016100de565b34801561016e57600080fd5b506100ca7f000000000000000000000000000000000000000000000000000000000000000081565b3480156101a257600080fd5b506100886101b1366004610ef5565b610554565b6101c0828261076a565b6000828152602081905260409020546101d990826107ce565b6000838152602081905260409020556101f282826107e1565b817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a98260405161022491815260200190565b60405180910390a25050565b6002546001600160a01b0316331461029d5760405162461bcd60e51b815260206004820152602560248201527f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960448201526431b0ba37b960d91b60648201526084015b60405180910390fd5b8281146102fe5760405162461bcd60e51b815260206004820152602960248201527f7061727469636970616e7473206c656e6774682073686f756c6420657175616c6044820152682062616c616e63657360b81b6064820152608401610294565b60008581526001602052604090205460ff161561036b5760405162461bcd60e51b815260206004820152602560248201527f747279696e6720746f2073657420616c726561647920736574746c6564206368604482015264185b9b995b60da1b6064820152608401610294565b600085815260208190526040812080549082905590808567ffffffffffffffff81111561039a5761039a610f94565b6040519080825280602002602001820160405280156103c3578160200160208202803683370190505b50905060005b8681101561049757600061040e8a8a8a858181106103e9576103e9610faa565b90506020028101906103fb9190610fc0565b610409906020810190610ff5565b6108ba565b90508083838151811061042357610423610faa565b60200260200101818152505061045460008083815260200190815260200160002054866107ce90919063ffffffff16565b945061048187878481811061046b5761046b610faa565b90506020020135856107ce90919063ffffffff16565b935050808061048f90611028565b9150506103c9565b508183106105075760005b86811015610505578585828181106104bc576104bc610faa565b905060200201356000808484815181106104d8576104d8610faa565b602002602001015181526020019081526020016000208190555080806104fd90611028565b9150506104a2565b505b6000888152600160208190526040808320805460ff19169092179091555189917fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b891a25050505050505050565b823560009081526001602052604090205460ff166105aa5760405162461bcd60e51b815260206004820152601360248201527218da185b9b995b081b9bdd081cd95d1d1b1959606a1b6044820152606401610294565b61061d836040516020016105be9190611051565b60408051601f198184030181526020601f86018190048102840181019092528483529190859085908190840183828082843760009201919091525061060a925050506020870187610fc0565b610618906020810190610ff5565b6108ff565b6106695760405162461bcd60e51b815260206004820152601d60248201527f7369676e617475726520766572696669636174696f6e206661696c65640000006044820152606401610294565b600061067d84356103fb6020870187610fc0565b600081815260208190526040902054909150606085013511156106d75760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610294565b6000818152602081905260409020546106f490606086013561098d565b60008281526020819052604090205561070e848484610999565b807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81606086018035906107449060408901610ff5565b604080519283526001600160a01b0390911660208301520160405180910390a250505050565b34156107ca5760405162461bcd60e51b815260206004820152602960248201527f6d6573736167652076616c7565206d757374206265203020666f7220746f6b656044820152681b8819195c1bdcda5d60ba1b6064820152608401610294565b5050565b60006107da8284611134565b9392505050565b6040516323b872dd60e01b8152336004820152306024820152604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610854573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610878919061114c565b6107ca5760405162461bcd60e51b81526020600482015260136024820152721d1c985b9cd9995c919c9bdb4819985a5b1959606a1b6044820152606401610294565b600082826040516020016108e19291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120905092915050565b60008061096085805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600061096e8286610a8a565b6001600160a01b0390811690851614925050509392505050565b505050565b60006107da828461116e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb6109d86060860160408701610ff5565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152606086013560248201526044016020604051808303816000875af1158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c919061114c565b6109885760405162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b6044820152606401610294565b6000806000610a998585610aae565b91509150610aa681610b1c565b509392505050565b6000808251604103610ae45760208301516040840151606085015160001a610ad887828585610cd5565b94509450505050610b15565b8251604003610b0d5760208301516040840151610b02868383610dc2565b935093505050610b15565b506000905060025b9250929050565b6000816004811115610b3057610b30611185565b03610b385750565b6001816004811115610b4c57610b4c611185565b03610b995760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610294565b6002816004811115610bad57610bad611185565b03610bfa5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610294565b6003816004811115610c0e57610c0e611185565b03610c665760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610294565b6004816004811115610c7a57610c7a611185565b03610cd25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610294565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d0c5750600090506003610db9565b8460ff16601b14158015610d2457508460ff16601c14155b15610d355750600090506004610db9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d89573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610db257600060019250925050610db9565b9150600090505b94509492505050565b6000806001600160ff1b03831681610ddf60ff86901c601b611134565b9050610ded87828885610cd5565b935093505050935093915050565b60008060408385031215610e0e57600080fd5b50508035926020909101359150565b60008083601f840112610e2f57600080fd5b50813567ffffffffffffffff811115610e4757600080fd5b6020830191508360208260051b8501011115610b1557600080fd5b600080600080600060608688031215610e7a57600080fd5b85359450602086013567ffffffffffffffff80821115610e9957600080fd5b610ea589838a01610e1d565b90965094506040880135915080821115610ebe57600080fd5b50610ecb88828901610e1d565b969995985093965092949392505050565b600060208284031215610eee57600080fd5b5035919050565b600080600060408486031215610f0a57600080fd5b833567ffffffffffffffff80821115610f2257600080fd5b9085019060808288031215610f3657600080fd5b90935060208501359080821115610f4c57600080fd5b818601915086601f830112610f6057600080fd5b813581811115610f6f57600080fd5b876020828501011115610f8157600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112610fd657600080fd5b9190910192915050565b6001600160a01b0381168114610cd257600080fd5b60006020828403121561100757600080fd5b81356107da81610fe0565b634e487b7160e01b600052601160045260246000fd5b60006001820161103a5761103a611012565b5060010190565b803561104c81610fe0565b919050565b602081528135602082015260006020830135603e1984360301811261107557600080fd5b608060408401528301803561108981610fe0565b6001600160a01b031660a0840152602081013536829003601e190181126110af57600080fd5b0160208101903567ffffffffffffffff8111156110cb57600080fd5b8036038213156110da57600080fd5b604060c08501528060e0850152610100818382870137600081838701015261110460408701611041565b6001600160a01b03811660608701529250606095909501356080850152601f01601f191690920190920192915050565b6000821982111561114757611147611012565b500190565b60006020828403121561115e57600080fd5b815180151581146107da57600080fd5b60008282101561118057611180611012565b500390565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220ad0874027a67427bfc782242eb9ed3252b2e96dbb6f2b27da3f355e3da35f48564736f6c634300080f0033"
