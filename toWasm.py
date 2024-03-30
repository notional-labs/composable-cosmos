import base64

code_base64 = ""
with open("msg_contract.json", "rb") as file:
    contract_bytes = file.read()
    print(contract_bytes)
    import json

    contract_data = json.loads(contract_bytes)
    code_value = contract_data["body"]["messages"][0]["code"]
    code_base64 = base64.b64decode(code_value)



with open('contract.wasm', 'wb') as wasm_file:
    quarter_length = len(code_base64)* 20 // 40 # This is to get a modified contract size 

    wasm_file.write(code_base64[:quarter_length])

