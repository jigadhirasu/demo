::javascript
protoc -I=protofile ^
 --js_out=import_style=commonjs:jspb ^
 --grpc-web_out=import_style=typescript,mode=grpcwebtext:jspb ^
 protofile/*.proto protofile/*.service 
