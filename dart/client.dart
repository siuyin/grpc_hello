// client.dart implements hello grpc client.

import 'dart:async';
import 'package:grpc/grpc.dart';

import 'package:hello/hello.pb.dart';
import 'package:hello/hello.pbgrpc.dart';

Future<void> main(List<String> args) async {
  final channel = ClientChannel('localhost',
      //port: 50051,
      port: 8080,
      options:
          const ChannelOptions(credentials: ChannelCredentials.insecure()));
  final stub = GreeterClient(channel);

  final name = args.isNotEmpty ? args[0] : 'world';

  try {
    //await for (var m in stub.sayHello(HelloRequest()..name = name)) {
    var req = HelloRequest();
    req.name = name;
    await for (var m in stub.sayHello(req)) {
      print('recv: ${m.message}');
    }
  } catch (e) {
    print('Caught error: $e');
  }
  await channel.shutdown();
}
