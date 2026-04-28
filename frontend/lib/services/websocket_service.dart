import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import '../models/message.dart';

class WebSocketService {
  late WebSocketChannel _channel;

  void connect(String url) {
    _channel = WebSocketChannel.connect(Uri.parse(url));
  }

  Stream<Message> get messageStream {
    return _channel.stream.map((data) {
      final json = jsonDecode(data);
      return Message(
        id: json['id'],
        senderId: json['sender_id'],
        content: json['content'],
        type: json['type'],
        timestamp: DateTime.parse(json['timestamp']),
      );
    });
  }

  void sendMessage(Message message) {
    final data = jsonEncode({
      'id': message.id,
      'sender_id': message.senderId,
      'content': message.content,
      'type': message.type,
    });
    _channel.sink.add(data);
  }

  void dispose() {
    _channel.sink.close();
  }
}