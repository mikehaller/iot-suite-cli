call openapi-generator generate -g go --additional-properties="packageName=iotthings" -o out-things -i things_2.yml

return

call openapi-generator generate -g go --additional-properties="packageName=devprov" -o out-device-provisioning -i api_1.yml
call openapi-generator generate -g go --additional-properties="packageName=iotmgr" -o out-iot-manager -i bosch-iot-manager.json
call openapi-generator generate -g go --additional-properties="packageName=hubdevreg" -o out-hub-device-registry -i bosch-iot-hub-device-registry.yml
call openapi-generator generate -g go --additional-properties="packageName=hubhttp" -o out-hub-http-adapter -i bosch-iot-hub-http-adapter.yml
