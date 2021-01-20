package node

var GenesisJson = `{
  "genesis_time": "2020-07-09T00:00:00.051674734Z",
  "chain_id": "Binance-Chain-Ganges",
  "consensus_params": {
    "block_size": {
      "max_bytes": "1048576",
      "max_gas": "-1"
    },
    "evidence": {
      "max_age": "100000"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    }
  },
  "app_hash": "",
  "app_state": {
    "tokens": [
      {
        "name": "Binance Chain Native Token",
        "symbol": "BNB",
        "total_supply": "20000000000000000",
        "owner": "tbnb1l9ffdr8e2pk7h4agvhwcslh2urwpuhqm2u82hy",
        "mintable": false
      }
    ],
    "accounts": [
      {
        "name": "Fuji",
        "address": "tbnb1l9ffdr8e2pk7h4agvhwcslh2urwpuhqm2u82hy",
        "consensus_addr": ""
      },
      {
        "name": "Fuji",
        "address": "tbnb1l9ffdr8e2pk7h4agvhwcslh2urwpuhqm2u82hy",
        "consensus_addr": "9CCDDD479C0AD8DCD01D754DD95FE15384E8BBDC"
      },
      {
        "name": "Kita",
        "address": "tbnb104s2prgzua8e3te6jumqsp0qn2dvqhlyemrf0g",
        "consensus_addr": ""
      },
      {
        "name": "Kita",
        "address": "tbnb104s2prgzua8e3te6jumqsp0qn2dvqhlyemrf0g",
        "consensus_addr": "F42F1D05AC568D12E26B9655395E7FBBD46BC5BB"
      },
      {
        "name": "Everest",
        "address": "tbnb19kecg6c93u7wh2cts4vsql52l33fa6a37y2tte",
        "consensus_addr": ""
      },
      {
        "name": "Everest",
        "address": "tbnb19kecg6c93u7wh2cts4vsql52l33fa6a37y2tte",
        "consensus_addr": "D4CECEF238E778C7063552C3A7AB95DA35C3FB47"
      }
    ],
    "dex": {},
    "param": {
      "fees": [
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "submit_proposal",
            "fee": "1000000000",
            "fee_for": 1
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "deposit",
            "fee": "125000",
            "fee_for": 1
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "vote",
            "fee": "0",
            "fee_for": 3
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "create_validator",
            "fee": "1000000000",
            "fee_for": 1
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "remove_validator",
            "fee": "100000000",
            "fee_for": 1
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "dexList",
            "fee": "200000000000",
            "fee_for": 2
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "orderNew",
            "fee": "0",
            "fee_for": 3
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "orderCancel",
            "fee": "0",
            "fee_for": 3
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "issueMsg",
            "fee": "100000000000",
            "fee_for": 2
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "mintMsg",
            "fee": "20000000000",
            "fee_for": 2
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "tokensBurn",
            "fee": "100000000",
            "fee_for": 1
          }
        },
        {
          "type": "params/FixedFeeParams",
          "value": {
            "msg_type": "tokensFreeze",
            "fee": "1000000",
            "fee_for": 1
          }
        },
        {
          "type": "params/TransferFeeParams",
          "value": {
            "fixed_fee_params": {
              "msg_type": "send",
              "fee": "62500",
              "fee_for": 1
            },
            "multi_transfer_fee": "50000",
            "lower_limit_as_multi": "2"
          }
        },
        {
          "type": "params/DexFeeParam",
          "value": {
            "dex_fee_fields": [
              {
                "fee_name": "ExpireFee",
                "fee_value": "50000"
              },
              {
                "fee_name": "ExpireFeeNative",
                "fee_value": "10000"
              },
              {
                "fee_name": "CancelFee",
                "fee_value": "50000"
              },
              {
                "fee_name": "CancelFeeNative",
                "fee_value": "10000"
              },
              {
                "fee_name": "FeeRate",
                "fee_value": "1000"
              },
              {
                "fee_name": "FeeRateNative",
                "fee_value": "400"
              },
              {
                "fee_name": "IOCExpireFee",
                "fee_value": "25000"
              },
              {
                "fee_name": "IOCExpireFeeNative",
                "fee_value": "5000"
              }
            ]
          }
        }
      ]
    },
    "stake": {
      "pool": {
        "loose_tokens": "20000000000000000",
        "bonded_tokens": "0"
      },
      "params": {
        "unbonding_time": "604800000000000",
        "max_validators": 21,
        "bond_denom": "BNB"
      },
      "validators": null,
      "bonds": null
    },
    "gov": {
      "starting_proposalID": "1",
      "deposit_params": {
        "min_deposit": [
          {
            "denom": "BNB",
            "amount": "100000000000"
          }
        ],
        "max_deposit_period": "172800000000000"
      },
      "tally_params": {
        "quorum": "50000000",
        "threshold": "50000000",
        "veto": "33400000"
      }
    },
    "gentxs": [
      {
        "type": "auth/StdTx",
        "value": {
          "msg": [
            {
              "type": "cosmos-sdk/MsgCreateValidatorProposal",
              "value": {
                "MsgCreateValidator": {
                  "Description": {
                    "moniker": "Fuji",
                    "identity": "",
                    "website": "",
                    "details": ""
                  },
                  "Commission": {
                    "rate": "0",
                    "max_rate": "0",
                    "max_change_rate": "0"
                  },
                  "delegator_address": "tbnb1l9ffdr8e2pk7h4agvhwcslh2urwpuhqm2u82hy",
                  "validator_address": "bva1l9ffdr8e2pk7h4agvhwcslh2urwpuhqmy407f3",
                  "pubkey": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "TRvmTw6aRmwuZqU0M5KBkng+Kfj6Ib6yEzSZte93D2A="
                  },
                  "delegation": {
                    "denom": "BNB",
                    "amount": "1000000000000"
                  }
                },
                "proposal_id": "0"
              }
            }
          ],
          "signatures": [
            {
              "pub_key": {
                "type": "tendermint/PubKeySecp256k1",
                "value": "A//eGMq4IXOWB/WC35sexMnrspP+2aC8QHWGOF3isHYs"
              },
              "signature": "eyDu8Pr93FRxz6+7nNNDNbynU9uLYxnTWMcexhMSEC9VCvpk1LZ16uvqApQw07LyZQuliB52QCb3VJfhUqRMGA==",
              "account_number": "0",
              "sequence": "0"
            }
          ],
          "memo": "",
          "source": "0",
          "data": null
        }
      },
      {
        "type": "auth/StdTx",
        "value": {
          "msg": [
            {
              "type": "cosmos-sdk/MsgCreateValidatorProposal",
              "value": {
                "MsgCreateValidator": {
                  "Description": {
                    "moniker": "Kita",
                    "identity": "",
                    "website": "",
                    "details": ""
                  },
                  "Commission": {
                    "rate": "0",
                    "max_rate": "0",
                    "max_change_rate": "0"
                  },
                  "delegator_address": "tbnb104s2prgzua8e3te6jumqsp0qn2dvqhlyemrf0g",
                  "validator_address": "bva104s2prgzua8e3te6jumqsp0qn2dvqhlyhjta3a",
                  "pubkey": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "AXdpIP8LDzjXjPlcAzwhrfcEV4URTjkqdUQXllLgphI="
                  },
                  "delegation": {
                    "denom": "BNB",
                    "amount": "1000000000000"
                  }
                },
                "proposal_id": "0"
              }
            }
          ],
          "signatures": [
            {
              "pub_key": {
                "type": "tendermint/PubKeySecp256k1",
                "value": "AxsnY0fJovfRkhmroAglhnKdPy57GOSM3lw2yeOt90on"
              },
              "signature": "lD/MdF306w0yoTsQG4eD62tfgzbSyd4NjvGnkRtq/hhzkdPo0WkK0xPohVtDlceLmJlPQduE0+roRczlif/3gA==",
              "account_number": "0",
              "sequence": "0"
            }
          ],
          "memo": "",
          "source": "0",
          "data": null
        }
      },
      {
        "type": "auth/StdTx",
        "value": {
          "msg": [
            {
              "type": "cosmos-sdk/MsgCreateValidatorProposal",
              "value": {
                "MsgCreateValidator": {
                  "Description": {
                    "moniker": "Everest",
                    "identity": "",
                    "website": "",
                    "details": ""
                  },
                  "Commission": {
                    "rate": "0",
                    "max_rate": "0",
                    "max_change_rate": "0"
                  },
                  "delegator_address": "tbnb19kecg6c93u7wh2cts4vsql52l33fa6a37y2tte",
                  "validator_address": "bva19kecg6c93u7wh2cts4vsql52l33fa6a3sdzl4v",
                  "pubkey": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": "mTCKo2XEBVS8iZgq9QXYXalSUURdXdSpuzfdJYT9ktM="
                  },
                  "delegation": {
                    "denom": "BNB",
                    "amount": "1000000000000"
                  }
                },
                "proposal_id": "0"
              }
            }
          ],
          "signatures": [
            {
              "pub_key": {
                "type": "tendermint/PubKeySecp256k1",
                "value": "AgVUOiG88AZKnRiHGjZnx2wKp0FrU9sseuDxFsui+Fmj"
              },
              "signature": "eXeQxIPQpHpwA1i1vxOMdj6d5Wov0xyAGjpgZs/eKx8PMRoay0n9sQQB4yZCSO/A6Ihn/pCuxtwCU2gK7zAkuA==",
              "account_number": "0",
              "sequence": "0"
            }
          ],
          "memo": "",
          "source": "0",
          "data": null
        }
      }
    ]
  }
}`
