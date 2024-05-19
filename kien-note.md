# Running the test

## How did i build the `hyperspace` and `cw-grandpa` binaries
Go to the `ComposableFi/composable-ibc`, and building it in the master branch 

https://github.com/ComposableFi/composable-ibc/tree/master/hyperspace

## How to reproduce the case

1. Build the `picad` binary
```bash
make install
```

2. Run the localnet pica with `wasmClient` deployed
```bash
make localnet-pica
```

3. Run the localnet picasso 
```bash
make localnet-picasso`
```

4. Run create clients 
```bash
make relayer-create-clients
```


# Current issue 
## Client State decode into 08-wasm
`6981`
### Logging msg at ValidateBasic on cosmos side on v47
add this log at ValidateBasic() function, at modules/core/02-client/types/msgs/go

```go
fmt.Printf("msg.ClientState : %v\n", msg.ClientState)
clientState, err := UnpackClientState(msg.ClientState)
if err != nil {
	return err
}
```

then, i get this value, basically, it says the clientState constructed from hyperspace is `08-wasm`

```
msg.ClientState : &Any{TypeUrl:/ibc.lightclients.wasm.v1.ClientState,Value:[10 204 1 10 40 47 105 98 99 46 108 105 103 104 116 99 108 105 101 110 116 115 46 103 114 97 110 100 112 97 46 118 49 46 67 108 105 101 110 116 83 116 97 116 101 18 159 1 10 32 146 240 69 84 49 1 104 67 240 92 67 123 199 101 70 152 115 0 205 47 91 76 8 16 108 92 254 197 173 59 35 206 16 67 24 5 40 2 48 167 16 56 20 66 36 10 32 199 203 131 204 79 173 68 214 92 44 140 240 46 150 66 83 144 154 219 206 75 174 118 245 90 177 194 134 163 255 44 194 16 1 66 36 10 32 20 164 40 156 190 217 24 43 102 84 102 53 144 50 192 109 195 183 142 183 238 104 52 237 95 167 180 159 149 51 85 167 16 1 66 36 10 32 184 160 36 189 114 208 123 150 110 8 135 7 155 93 135 60 197 160 19 53 186 39 222 219 43 141 20 36 169 95 136 128 16 1 18 32 157 80 86 242 181 81 33 48 148 160 59 120 141 184 2 116 68 81 117 186 231 140 143 237 227 139 242 90 45 220 132 188 26 5 8 167 16 16 20],XXX_unrecognized:[]}
```


### Logging it at v50

`hyperspace`

```
msg.ClientState : &Any{TypeUrl:/ibc.lightclients.grandpa.v1.ClientState,Value:[10 32 106 46 120 116 154 178 37 115 118 3 65 47 241 205 37 168 10 129 250 125 20 121 236 81 77 119 185 185 155 72 122 11 16 17 40 2 48 167 16 56 1 66 36 10 32 199 203 131 204 79 173 68 214 92 44 140 240 46 150 66 83 144 154 219 206 75 174 118 245 90 177 194 134 163 255 44 194 16 1 66 36 10 32 20 164 40 156 190 217 24 43 102 84 102 53 144 50 192 109 195 183 142 183 238 104 52 237 95 167 180 159 149 51 85 167 16 1 66 36 10 32 184 160 36 189 114 208 123 150 110 8 135 7 155 93 135 60 197 160 19 53 186 39 222 219 43 141 20 36 169 95 136 128 16 1 66 36 10 32 140 4 179 243 122 62 235 12 208 118 154 107 153 87 196 106 60 118 44 155 224 91 89 26 250 8 121 224 61 159 194 212 16 1],XXX_unrecognized:[]}
