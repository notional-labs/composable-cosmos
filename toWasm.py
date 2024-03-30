import base64

# code_base64 = ""
# with open("msg_contract.json", "rb") as file:
#     contract_bytes = file.read()
#     print(contract_bytes)
#     import json

#     contract_data = json.loads(contract_bytes)
#     code_value = contract_data["body"]["messages"][0]["code"]
#     code_base64 = base64.b64decode(code_value)



# with open('contract.wasm', 'wb') as wasm_file:
#     wasm_file.write(code_base64)


import gzip

with open('contract.wasm', 'rb') as original_wasm_file:
    with gzip.open('contract.wasm.gz', 'wb') as compressed_wasm_file:
        compressed_wasm_file.writelines(original_wasm_file)
